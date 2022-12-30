import React, { useState } from "react";
import {
  PlanExecutionRequest,
  PlanExecutionRequests,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { gql } from "../__generated__";
import Table from "react-bootstrap/Table";
import { renderTimeField } from "../utils/table_rendering";
import { Container } from "react-bootstrap";
import Accordion from "react-bootstrap/Accordion";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { dark } from "react-syntax-highlighter/dist/esm/styles/prism";
import Ansi from "ansi-to-react";

const PLAN_EXECUTION_REQUEST_QUERY = gql(`
  query planExecutionRequest($planExecutionRequestId: ID!) {
    planExecutionRequest(planExecutionRequestId: $planExecutionRequestId) {
      planExecutionRequestId
      status
      requestTime
      terraformConfigurationBase64
      initOutput {
        Stdout
        Stderr
      }
      planOutput {
        Stdout
        Stderr
      }
    }
  }
`);

type Response = {
  planExecutionRequest: PlanExecutionRequest;
};

export const PlanExecutionRequestPage = () => {
  const params = useParams();

  const planExecutionRequestId = params.planExecutionRequestId
    ? params.planExecutionRequestId
    : "";

  const { loading, error, data } = useQuery<Response>(
    PLAN_EXECUTION_REQUEST_QUERY,
    {
      variables: {
        planExecutionRequestId: planExecutionRequestId,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  let terraformConfiguration = data?.planExecutionRequest
    .terraformConfigurationBase64
    ? atob(data?.planExecutionRequest.terraformConfigurationBase64)
    : "";
  let initStdout = data?.planExecutionRequest.initOutput?.Stdout
    ? atob(data?.planExecutionRequest.initOutput?.Stdout)
    : "";
  let initStderr = data?.planExecutionRequest.initOutput?.Stderr
    ? atob(data?.planExecutionRequest.initOutput?.Stderr)
    : "";

  let planStdout = data?.planExecutionRequest.planOutput?.Stdout
    ? atob(data?.planExecutionRequest.planOutput?.Stdout)
    : "";
  let planStderr = data?.planExecutionRequest.planOutput?.Stderr
    ? atob(data?.planExecutionRequest.planOutput?.Stderr)
    : "";

  return (
    <Container>
      <h1>
        Plan Execution Request{" "}
        <b>{data?.planExecutionRequest.planExecutionRequestId}</b>
      </h1>
      <p>
        Status: <b>{data?.planExecutionRequest.status}</b>
      </p>
      <Accordion>
        <Accordion.Item eventKey="Config">
          <Accordion.Header>Terraform Configuration</Accordion.Header>
          <Accordion.Body>
            <Container
              className="bg-dark"
              style={{
                overflow: "auto",
                maxHeight: "60vh",
                whiteSpace: "pre-wrap",
                textAlign: "left",
              }}
            >
              <SyntaxHighlighter language="hcl" style={dark}>
                {terraformConfiguration}
              </SyntaxHighlighter>
            </Container>
          </Accordion.Body>
        </Accordion.Item>
        <Accordion.Item eventKey="Init">
          <Accordion.Header>Terraform Init Output</Accordion.Header>
          <Accordion.Body>
            <h3>Std Out</h3>
            <Container
              className="bg-dark"
              style={{
                overflow: "auto",
                maxHeight: "60vh",
                whiteSpace: "pre-wrap",
                textAlign: "left",
              }}
            >
              <Ansi className="ansi-black-bg">{initStdout}</Ansi>
            </Container>
            <h3>Std Err</h3>
            <Container
              className="bg-dark"
              style={{
                overflow: "auto",
                maxHeight: "60vh",
                whiteSpace: "pre-wrap",
                textAlign: "left",
              }}
            >
              <Ansi className="ansi-black-bg">{initStderr}</Ansi>
            </Container>
          </Accordion.Body>
        </Accordion.Item>
        <Accordion.Item eventKey="Plan">
          <Accordion.Header>Terraform Plan Output</Accordion.Header>
          <Accordion.Body>
            <h3>Std Out</h3>
            <Container
              className="bg-dark"
              style={{
                overflow: "auto",
                maxHeight: "60vh",
                whiteSpace: "pre-wrap",
                textAlign: "left",
              }}
            >
              <Ansi className="ansi-black-bg">{planStdout}</Ansi>
            </Container>
            <h3>Std Err</h3>
            <Container
              className="bg-dark"
              style={{
                overflow: "auto",
                maxHeight: "60vh",
                whiteSpace: "pre-wrap",
                textAlign: "left",
              }}
            >
              <Ansi className="ansi-black-bg">{planStderr}</Ansi>
            </Container>
          </Accordion.Body>
        </Accordion.Item>
      </Accordion>
    </Container>
  );
};

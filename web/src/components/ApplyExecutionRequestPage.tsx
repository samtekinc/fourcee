import React, { useState } from "react";
import {
  ApplyExecutionRequest,
  ApplyExecutionRequests,
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

const APPLY_EXECUTION_REQUEST_QUERY = gql(`
  query applyExecutionRequest($applyExecutionRequestId: ID!) {
    applyExecutionRequest(applyExecutionRequestId: $applyExecutionRequestId) {
      applyExecutionRequestId
      status
      requestTime
      terraformConfigurationBase64
      initOutput {
        Stdout
        Stderr
      }
      applyOutput {
        Stdout
        Stderr
      }
    }
  }
`);

type Response = {
  applyExecutionRequest: ApplyExecutionRequest;
};

export const ApplyExecutionRequestPage = () => {
  const params = useParams();

  const applyExecutionRequestId = params.applyExecutionRequestId
    ? params.applyExecutionRequestId
    : "";

  const { loading, error, data } = useQuery<Response>(
    APPLY_EXECUTION_REQUEST_QUERY,
    {
      variables: {
        applyExecutionRequestId: applyExecutionRequestId,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  let terraformConfiguration = data?.applyExecutionRequest
    .terraformConfigurationBase64
    ? atob(data?.applyExecutionRequest.terraformConfigurationBase64)
    : "";
  let initStdout = data?.applyExecutionRequest.initOutput?.Stdout
    ? atob(data?.applyExecutionRequest.initOutput?.Stdout)
    : "";
  let initStderr = data?.applyExecutionRequest.initOutput?.Stderr
    ? atob(data?.applyExecutionRequest.initOutput?.Stderr)
    : "";

  let applyStdout = data?.applyExecutionRequest.applyOutput?.Stdout
    ? atob(data?.applyExecutionRequest.applyOutput?.Stdout)
    : "";
  let applyStderr = data?.applyExecutionRequest.applyOutput?.Stderr
    ? atob(data?.applyExecutionRequest.applyOutput?.Stderr)
    : "";

  return (
    <Container>
      <h1>
        Apply Execution Request{" "}
        <b>{data?.applyExecutionRequest.applyExecutionRequestId}</b>
      </h1>
      <p>
        Status: <b>{data?.applyExecutionRequest.status}</b>
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
        <Accordion.Item eventKey="Apply">
          <Accordion.Header>Terraform Apply Output</Accordion.Header>
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
              <Ansi className="ansi-black-bg">{applyStdout}</Ansi>
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
              <Ansi className="ansi-black-bg">{applyStderr}</Ansi>
            </Container>
          </Accordion.Body>
        </Accordion.Item>
      </Accordion>
    </Container>
  );
};

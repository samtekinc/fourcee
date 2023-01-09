import React, { useState } from "react";
import {
  PlanExecutionRequest,
  PlanExecutionRequests,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import { Container } from "react-bootstrap";
import Accordion from "react-bootstrap/Accordion";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { vscDarkPlus } from "react-syntax-highlighter/dist/esm/styles/prism";
import Ansi from "ansi-to-react";

const PLAN_EXECUTION_REQUEST_QUERY = gql`
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
`;

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
    : "...";
  let initStdout = data?.planExecutionRequest.initOutput?.Stdout
    ? atob(data?.planExecutionRequest.initOutput?.Stdout)
    : "...";
  let initStderr = data?.planExecutionRequest.initOutput?.Stderr
    ? atob(data?.planExecutionRequest.initOutput?.Stderr)
    : "...";

  let planStdout = data?.planExecutionRequest.planOutput?.Stdout
    ? atob(data?.planExecutionRequest.planOutput?.Stdout)
    : "...";
  let planStderr = data?.planExecutionRequest.planOutput?.Stderr
    ? atob(data?.planExecutionRequest.planOutput?.Stderr)
    : "...";

  return (
    <Container>
      <h1>
        Plan Execution Request{" "}
        <b>{data?.planExecutionRequest.planExecutionRequestId}</b>
      </h1>
      <p>
        Status: <b>{renderStatus(data?.planExecutionRequest.status)}</b>
      </p>
      <h2>Terraform Configuration</h2>
      <Container
        className="bg-dark"
        style={{
          overflow: "auto",
          maxHeight: "60vh",
          whiteSpace: "pre-wrap",
          textAlign: "left",
        }}
      >
        <SyntaxHighlighter language="hcl" style={vscDarkPlus}>
          {terraformConfiguration}
        </SyntaxHighlighter>
      </Container>
      <h2>Terraform Init Output</h2>
      Std Out
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
      Std Err
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
      <h2>Terraform Plan Output</h2>
      Std Out
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
      Std Err
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
    </Container>
  );
};

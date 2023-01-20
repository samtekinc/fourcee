import { PlanExecutionRequest, RequestStatus } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import { renderStatus } from "../utils/table_rendering";
import { Col, Container, Row } from "react-bootstrap";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { vscDarkPlus } from "react-syntax-highlighter/dist/esm/styles/prism";
import Ansi from "ansi-to-react";
import { b64DecodeUnicode } from "../utils/decoding";

const PLAN_EXECUTION_REQUEST_QUERY = gql`
  query planExecutionRequest($planExecutionRequestId: ID!) {
    planExecutionRequest(planExecutionRequestId: $planExecutionRequestId) {
      planExecutionRequestId
      status
      requestTime
      terraformConfigurationBase64
      moduleAssignment {
        name
        moduleAssignmentId
        modulePropagation {
          modulePropagationId
          name
        }
        orgAccount {
          orgAccountId
          name
          cloudPlatform
        }
        moduleGroup {
          moduleGroupId
          name
        }
        moduleVersion {
          moduleVersionId
          name
        }
      }
      initOutput
      planOutput
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

  const { loading, error, data, startPolling } = useQuery<Response>(
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

  if (
    data?.planExecutionRequest?.status === RequestStatus.Running ||
    data?.planExecutionRequest?.status === RequestStatus.Pending
  ) {
    startPolling(1000);
  } else {
    startPolling(30000);
  }

  let terraformConfiguration = data?.planExecutionRequest
    .terraformConfigurationBase64
    ? b64DecodeUnicode(data?.planExecutionRequest.terraformConfigurationBase64)
    : "...";

  let initOutput = data?.planExecutionRequest.initOutput ?? "...";
  let planOutput = data?.planExecutionRequest.planOutput ?? "...";

  return (
    <Container
      fluid
      style={{
        paddingTop: "2rem",
        paddingBottom: "2rem",
        paddingLeft: "5rem",
        paddingRight: "5rem",
      }}
    >
      <h1>
        <b>{data?.planExecutionRequest.planExecutionRequestId}</b>
      </h1>
      <p>
        <b>Org Account: </b>
        <NavLink
          to={`/org-accounts/${data?.planExecutionRequest.moduleAssignment.orgAccount.orgAccountId}`}
        >
          {data?.planExecutionRequest.moduleAssignment.orgAccount.name}
        </NavLink>
        <br />
        <b>Module: </b>
        <NavLink
          to={`/module-groups/${data?.planExecutionRequest.moduleAssignment.moduleGroup.moduleGroupId}`}
        >
          {data?.planExecutionRequest.moduleAssignment.moduleGroup.name}
        </NavLink>
        {" / "}
        <NavLink
          to={`/module-groups/${data?.planExecutionRequest.moduleAssignment.moduleGroup.moduleGroupId}/versions/${data?.planExecutionRequest.moduleAssignment.moduleVersion.moduleVersionId}`}
        >
          {data?.planExecutionRequest.moduleAssignment.moduleVersion.name}
        </NavLink>
        <br />
        <b>Plan Status: </b>
        {renderStatus(data?.planExecutionRequest.status)}
        <br />
      </p>
      <h2>Terraform Configuration</h2>
      <Container
        fluid
        className="bg-dark"
        style={{
          overflow: "auto",
          maxHeight: "60vh",
          whiteSpace: "pre",
          textAlign: "left",
        }}
      >
        <SyntaxHighlighter language="hcl" style={vscDarkPlus}>
          {terraformConfiguration}
        </SyntaxHighlighter>
      </Container>
      <br />
      <Row>
        <Col md={6}>
          <h2>Terraform Init Output</h2>
          <Container
            className="bg-dark"
            fluid
            style={{
              overflow: "auto",
              maxHeight: "80vh",
              whiteSpace: "pre",
              textAlign: "left",
            }}
          >
            <Ansi className="ansi-black-bg">{initOutput}</Ansi>
          </Container>
        </Col>
        <Col md={6}>
          <h2>Terraform Plan Output</h2>
          <Container
            className="bg-dark"
            style={{
              overflow: "auto",
              maxHeight: "80vh",
              whiteSpace: "pre",
              textAlign: "left",
            }}
          >
            <Ansi className="ansi-black-bg">{planOutput}</Ansi>
          </Container>
        </Col>
      </Row>
    </Container>
  );
};

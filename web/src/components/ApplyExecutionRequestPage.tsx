import { ApplyExecutionRequest, RequestStatus } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import { renderStatus } from "../utils/table_rendering";
import { Col, Container, Row } from "react-bootstrap";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { vscDarkPlus } from "react-syntax-highlighter/dist/esm/styles/prism";
import Ansi from "ansi-to-react";

const APPLY_EXECUTION_REQUEST_QUERY = gql`
  query applyExecutionRequest($applyExecutionRequestID: ID!) {
    applyExecutionRequest(applyExecutionRequestID: $applyExecutionRequestID) {
      id
      status
      startedAt
      completedAt
      terraformConfiguration
      moduleAssignment {
        name
        modulePropagation {
          name
        }
        orgAccount {
          id
          name
          cloudPlatform
        }
        moduleGroup {
          id
          name
        }
        moduleVersion {
          id
          name
        }
      }
      initOutput
      applyOutput
    }
  }
`;

type Response = {
  applyExecutionRequest: ApplyExecutionRequest;
};

export const ApplyExecutionRequestPage = () => {
  const params = useParams();

  const applyExecutionRequestID = params.applyExecutionRequestID
    ? params.applyExecutionRequestID
    : "";

  const { loading, error, data, startPolling } = useQuery<Response>(
    APPLY_EXECUTION_REQUEST_QUERY,
    {
      variables: {
        applyExecutionRequestID: applyExecutionRequestID,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  if (
    data?.applyExecutionRequest?.status === RequestStatus.Running ||
    data?.applyExecutionRequest?.status === RequestStatus.Pending
  ) {
    startPolling(1000);
  } else {
    startPolling(30000);
  }

  let terraformConfiguration = data?.applyExecutionRequest
    .terraformConfiguration
    ? data?.applyExecutionRequest.terraformConfiguration
    : "...";

  let initOutput = data?.applyExecutionRequest.initOutput ?? "...";
  let applyOutput = data?.applyExecutionRequest.applyOutput ?? "...";
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
        <b>{data?.applyExecutionRequest.id}</b>
      </h1>
      <p>
        <b>Org Account: </b>
        <NavLink
          to={`/org-accounts/${data?.applyExecutionRequest.moduleAssignment.orgAccount.id}`}
        >
          {data?.applyExecutionRequest.moduleAssignment.orgAccount.name}
        </NavLink>
        <br />
        <b>Module: </b>
        <NavLink
          to={`/module-groups/${data?.applyExecutionRequest.moduleAssignment.moduleGroup.id}`}
        >
          {data?.applyExecutionRequest.moduleAssignment.moduleGroup.name}
        </NavLink>
        {" / "}
        <NavLink
          to={`/module-groups/${data?.applyExecutionRequest.moduleAssignment.moduleGroup.id}/versions/${data?.applyExecutionRequest.moduleAssignment.moduleVersion.id}`}
        >
          {data?.applyExecutionRequest.moduleAssignment.moduleVersion.name}
        </NavLink>
        <br />
        <b>Apply Status: </b>
        {renderStatus(data?.applyExecutionRequest.status)}
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
          <h2>Init Output</h2>
          <Container
            className="bg-dark"
            style={{
              overflow: "auto",
              maxHeight: "80vh",
              whiteSpace: "pre",
              textAlign: "left",
              color: "white",
            }}
          >
            <Ansi className="ansi-black-bg">{initOutput}</Ansi>
          </Container>
        </Col>
        <Col md={6}>
          <h2>Apply Output</h2>
          <Container
            className="bg-dark"
            style={{
              overflow: "auto",
              maxHeight: "60vh",
              whiteSpace: "pre",
              textAlign: "left",
            }}
          >
            <Ansi className="ansi-black-bg">{applyOutput}</Ansi>
          </Container>
        </Col>
      </Row>
    </Container>
  );
};

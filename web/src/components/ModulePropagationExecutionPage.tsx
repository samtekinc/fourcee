import {
  ModulePropagationExecutionRequest,
  RequestStatus,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus } from "../utils/table_rendering";
import { Col, Container, Row } from "react-bootstrap";
import { renderApplyDestroy, renderCloudPlatform } from "../utils/rendering";

const MODULE_PROPAGATION_EXECUTION_QUERY = gql`
  query modulePropagationExecutionRequest(
    $modulePropagationExecutionRequestID: ID!
  ) {
    modulePropagationExecutionRequest(
      modulePropagationExecutionRequestID: $modulePropagationExecutionRequestID
    ) {
      id
      modulePropagation {
        name
      }
      startedAt
      status
      terraformExecutionRequests {
        id
        status
        startedAt
        destroy
        moduleAssignment {
          id
          orgAccount {
            id
            cloudPlatform
            name
          }
        }
        planExecutionRequest {
          id
          status
          startedAt
        }
        applyExecutionRequest {
          id
          status
          startedAt
        }
      }
    }
  }
`;

type Response = {
  modulePropagationExecutionRequest: ModulePropagationExecutionRequest;
};

export const ModulePropagationExecutionRequestPage = () => {
  const params = useParams();

  const modulePropagationExecutionRequestID =
    params.modulePropagationExecutionRequestID
      ? params.modulePropagationExecutionRequestID
      : "";
  const modulePropagationID = params.modulePropagationID
    ? params.modulePropagationID
    : "";

  const { loading, error, data, startPolling } = useQuery<Response>(
    MODULE_PROPAGATION_EXECUTION_QUERY,
    {
      variables: {
        modulePropagationExecutionRequestID:
          modulePropagationExecutionRequestID,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  if (
    data?.modulePropagationExecutionRequest?.status === RequestStatus.Running ||
    data?.modulePropagationExecutionRequest?.status === RequestStatus.Pending
  ) {
    startPolling(1000);
  } else {
    startPolling(30000);
  }

  return (
    <Container fluid>
      <h1>Terraform Execution #{data?.modulePropagationExecutionRequest.id}</h1>
      Status: {renderStatus(data?.modulePropagationExecutionRequest.status)}
      <br />
      <Row>
        <Col md={"auto"}>
          <h3>Terraform Workflows</h3>
          <Table hover>
            <thead>
              <tr>
                <th>Status</th>
                <th>Org Account</th>
                <th>Action</th>
                <th>Plan Request</th>
                <th>Apply Request</th>
              </tr>
            </thead>
            <tbody>
              {data?.modulePropagationExecutionRequest.terraformExecutionRequests.map(
                (terraformExecutionRequest) => {
                  return (
                    <tr>
                      <td>{renderStatus(terraformExecutionRequest?.status)}</td>
                      <td>
                        {renderCloudPlatform(
                          terraformExecutionRequest?.moduleAssignment.orgAccount
                            .cloudPlatform
                        )}{" "}
                        <NavLink
                          to={`/org-accounts/${terraformExecutionRequest?.moduleAssignment.orgAccount.id}`}
                        >
                          {
                            terraformExecutionRequest?.moduleAssignment
                              .orgAccount.name
                          }
                        </NavLink>
                      </td>
                      <td>
                        {renderApplyDestroy(
                          terraformExecutionRequest?.destroy ?? false
                        )}
                      </td>
                      <td>
                        <NavLink
                          to={`/plan-execution-requests/${terraformExecutionRequest?.planExecutionRequest?.id}`}
                        >
                          {renderStatus(
                            terraformExecutionRequest?.planExecutionRequest
                              ?.status
                          )}
                        </NavLink>
                      </td>
                      <td>
                        <NavLink
                          to={`/apply-execution-requests/${terraformExecutionRequest?.applyExecutionRequest?.id}`}
                        >
                          {renderStatus(
                            terraformExecutionRequest?.applyExecutionRequest
                              ?.status
                          )}
                        </NavLink>
                      </td>
                    </tr>
                  );
                }
              )}
            </tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};

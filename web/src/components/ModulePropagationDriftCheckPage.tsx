import {
  ModulePropagationDriftCheckRequest,
  RequestStatus,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus } from "../utils/table_rendering";
import { Col, Container, Row } from "react-bootstrap";
import {
  renderApplyDestroy,
  renderCloudPlatform,
  renderSyncStatus,
} from "../utils/rendering";

const MODULE_PROPAGATION_DRIFT_CHECK_QUERY = gql`
  query modulePropagationDriftCheckRequest(
    $modulePropagationDriftCheckRequestID: ID!
  ) {
    modulePropagationDriftCheckRequest(
      modulePropagationDriftCheckRequestID: $modulePropagationDriftCheckRequestID
    ) {
      id
      startedAt
      status
      terraformDriftCheckRequests {
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
        syncStatus
      }
    }
  }
`;

type Response = {
  modulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
};

export const ModulePropagationDriftCheckRequestPage = () => {
  const params = useParams();

  const modulePropagationDriftCheckRequestID =
    params.modulePropagationDriftCheckRequestID
      ? params.modulePropagationDriftCheckRequestID
      : "";
  const modulePropagationID = params.modulePropagationID
    ? params.modulePropagationID
    : "";

  const { loading, error, data, startPolling } = useQuery<Response>(
    MODULE_PROPAGATION_DRIFT_CHECK_QUERY,
    {
      variables: {
        modulePropagationDriftCheckRequestID:
          modulePropagationDriftCheckRequestID,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  if (
    data?.modulePropagationDriftCheckRequest?.status ===
      RequestStatus.Running ||
    data?.modulePropagationDriftCheckRequest?.status === RequestStatus.Pending
  ) {
    startPolling(1000);
  } else {
    startPolling(30000);
  }

  return (
    <Container fluid>
      <h1>{data?.modulePropagationDriftCheckRequest.id}</h1>
      Status: {renderStatus(data?.modulePropagationDriftCheckRequest.status)}
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
                <th>Sync Status</th>
              </tr>
            </thead>
            <tbody>
              {data?.modulePropagationDriftCheckRequest.terraformDriftCheckRequests.map(
                (terraformDriftCheckRequest) => {
                  return (
                    <tr>
                      <td>
                        {renderStatus(terraformDriftCheckRequest?.status)}
                      </td>
                      <td>
                        {renderCloudPlatform(
                          terraformDriftCheckRequest?.moduleAssignment
                            .orgAccount.cloudPlatform
                        )}{" "}
                        <NavLink
                          to={`/org-accounts/${terraformDriftCheckRequest?.moduleAssignment.orgAccount.id}`}
                        >
                          {
                            terraformDriftCheckRequest?.moduleAssignment
                              .orgAccount.name
                          }
                        </NavLink>
                      </td>
                      <td>
                        {renderApplyDestroy(
                          terraformDriftCheckRequest?.destroy ?? false
                        )}
                      </td>
                      <td>
                        <NavLink
                          to={`/plan-execution-requests/${terraformDriftCheckRequest?.planExecutionRequest?.id}`}
                        >
                          {renderStatus(
                            terraformDriftCheckRequest?.planExecutionRequest
                              ?.status
                          )}
                        </NavLink>
                      </td>
                      <td>
                        {renderSyncStatus(
                          terraformDriftCheckRequest?.syncStatus
                        )}
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

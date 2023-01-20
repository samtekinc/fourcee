import React, { useState } from "react";
import {
  ModulePropagationExecutionRequest,
  ModulePropagationExecutionRequests,
  RequestStatus,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import { Breadcrumb, Col, Container, Row } from "react-bootstrap";
import { renderApplyDestroy, renderCloudPlatform } from "../utils/rendering";

const MODULE_PROPAGATION_EXECUTION_QUERY = gql`
  query modulePropagationExecutionRequest(
    $modulePropagationId: ID!
    $modulePropagationExecutionRequestId: ID!
  ) {
    modulePropagationExecutionRequest(
      modulePropagationId: $modulePropagationId
      modulePropagationExecutionRequestId: $modulePropagationExecutionRequestId
    ) {
      modulePropagationId
      modulePropagation {
        name
      }
      modulePropagationExecutionRequestId
      requestTime
      status
      terraformExecutionRequests {
        items {
          terraformExecutionRequestId
          status
          requestTime
          destroy
          moduleAssignment {
            moduleAssignmentId
            orgAccount {
              cloudPlatform
              orgAccountId
              name
            }
          }
          planExecutionRequest {
            planExecutionRequestId
            status
            requestTime
          }
          applyExecutionRequest {
            applyExecutionRequestId
            status
            requestTime
          }
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

  const modulePropagationExecutionRequestId =
    params.modulePropagationExecutionRequestId
      ? params.modulePropagationExecutionRequestId
      : "";
  const modulePropagationId = params.modulePropagationId
    ? params.modulePropagationId
    : "";

  const { loading, error, data, startPolling } = useQuery<Response>(
    MODULE_PROPAGATION_EXECUTION_QUERY,
    {
      variables: {
        modulePropagationExecutionRequestId:
          modulePropagationExecutionRequestId,
        modulePropagationId: modulePropagationId,
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
      <h1>
        {
          data?.modulePropagationExecutionRequest
            .modulePropagationExecutionRequestId
        }
      </h1>
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
              {data?.modulePropagationExecutionRequest.terraformExecutionRequests.items.map(
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
                          to={`/org-accounts/${terraformExecutionRequest?.moduleAssignment.orgAccount.orgAccountId}`}
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
                          to={`/plan-execution-requests/${terraformExecutionRequest?.planExecutionRequest?.planExecutionRequestId}`}
                        >
                          {renderStatus(
                            terraformExecutionRequest?.planExecutionRequest
                              ?.status
                          )}
                        </NavLink>
                      </td>
                      <td>
                        <NavLink
                          to={`/apply-execution-requests/${terraformExecutionRequest?.applyExecutionRequest?.applyExecutionRequestId}`}
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

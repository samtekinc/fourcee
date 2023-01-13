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
import { Container } from "react-bootstrap";

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
    <Container>
      <h1>
        Module Propagation Execution Request{" "}
        {
          data?.modulePropagationExecutionRequest
            .modulePropagationExecutionRequestId
        }
      </h1>
      <p>
        Status: {renderStatus(data?.modulePropagationExecutionRequest.status)}
      </p>
      <h2>Terraform Workflows</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>TF Request</th>
            <th>Apply / Destroy</th>
            <th>Org Account</th>
            <th>Status</th>
            <th>Plan Request</th>
            <th>Apply Request</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagationExecutionRequest.terraformExecutionRequests.items.map(
            (terraformExecutionRequest) => {
              return (
                <tr>
                  <td>
                    {terraformExecutionRequest?.terraformExecutionRequestId}
                  </td>
                  <td>
                    {terraformExecutionRequest?.destroy ? "Destroy" : "Apply"}
                  </td>
                  <td>
                    <NavLink
                      to={`/org-accounts/${terraformExecutionRequest?.moduleAssignment.orgAccount.orgAccountId}`}
                    >
                      {
                        terraformExecutionRequest?.moduleAssignment.orgAccount
                          .name
                      }{" "}
                      (
                      {
                        terraformExecutionRequest?.moduleAssignment.orgAccount
                          .orgAccountId
                      }
                      )
                    </NavLink>
                  </td>
                  <td>{renderStatus(terraformExecutionRequest?.status)}</td>
                  <td>
                    <NavLink
                      to={`/plan-execution-requests/${terraformExecutionRequest?.planExecutionRequest?.planExecutionRequestId}`}
                    >
                      {
                        terraformExecutionRequest?.planExecutionRequest
                          ?.planExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformExecutionRequest?.planExecutionRequest?.status
                    )}
                    )
                  </td>
                  <td>
                    <NavLink
                      to={`/apply-execution-requests/${terraformExecutionRequest?.applyExecutionRequest?.applyExecutionRequestId}`}
                    >
                      {
                        terraformExecutionRequest?.applyExecutionRequest
                          ?.applyExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformExecutionRequest?.applyExecutionRequest?.status
                    )}
                    )
                  </td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};

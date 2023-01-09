import React, { useState } from "react";
import {
  ModulePropagationExecutionRequest,
  ModulePropagationExecutionRequests,
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
      terraformExecutionWorkflowRequests {
        items {
          terraformExecutionWorkflowRequestId
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

  const { loading, error, data } = useQuery<Response>(
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
          {data?.modulePropagationExecutionRequest.terraformExecutionWorkflowRequests.items.map(
            (terraformExecutionWorkflowRequest) => {
              return (
                <tr>
                  <td>
                    {
                      terraformExecutionWorkflowRequest?.terraformExecutionWorkflowRequestId
                    }
                  </td>
                  <td>
                    {terraformExecutionWorkflowRequest?.destroy
                      ? "Destroy"
                      : "Apply"}
                  </td>
                  <td>
                    <NavLink
                      to={`/org-accounts/${terraformExecutionWorkflowRequest?.moduleAssignment.orgAccount.orgAccountId}`}
                    >
                      {
                        terraformExecutionWorkflowRequest?.moduleAssignment
                          .orgAccount.name
                      }{" "}
                      (
                      {
                        terraformExecutionWorkflowRequest?.moduleAssignment
                          .orgAccount.orgAccountId
                      }
                      )
                    </NavLink>
                  </td>
                  <td>
                    {renderStatus(terraformExecutionWorkflowRequest?.status)}
                  </td>
                  <td>
                    <NavLink
                      to={`/plan-execution-requests/${terraformExecutionWorkflowRequest?.planExecutionRequest?.planExecutionRequestId}`}
                    >
                      {
                        terraformExecutionWorkflowRequest?.planExecutionRequest
                          ?.planExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformExecutionWorkflowRequest?.planExecutionRequest
                        ?.status
                    )}
                    )
                  </td>
                  <td>
                    <NavLink
                      to={`/apply-execution-requests/${terraformExecutionWorkflowRequest?.applyExecutionRequest?.applyExecutionRequestId}`}
                    >
                      {
                        terraformExecutionWorkflowRequest?.applyExecutionRequest
                          ?.applyExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformExecutionWorkflowRequest?.applyExecutionRequest
                        ?.status
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

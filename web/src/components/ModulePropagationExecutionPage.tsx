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
      terraformWorkflowRequests {
        items {
          terraformWorkflowRequestId
          status
          requestTime
          destroy
          orgAccount {
            orgAccountId
            name
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
      planExecutionRequests {
        items {
          planExecutionRequestId
          status
          requestTime
          orgAccount {
            orgAccountId
            name
          }
        }
      }
      applyExecutionRequests {
        items {
          applyExecutionRequestId
          status
          requestTime
          orgAccount {
            orgAccountId
            name
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
          {data?.modulePropagationExecutionRequest.terraformWorkflowRequests.items.map(
            (terraformWorkflowRequest) => {
              return (
                <tr>
                  <td>
                    {terraformWorkflowRequest?.terraformWorkflowRequestId}
                  </td>
                  <td>
                    {terraformWorkflowRequest?.destroy ? "Destroy" : "Apply"}
                  </td>
                  <td>
                    <NavLink
                      to={`/org-accounts/${terraformWorkflowRequest?.orgAccount.orgAccountId}`}
                    >
                      {terraformWorkflowRequest?.orgAccount.name} (
                      {terraformWorkflowRequest?.orgAccount.orgAccountId})
                    </NavLink>
                  </td>
                  <td>{renderStatus(terraformWorkflowRequest?.status)}</td>
                  <td>
                    <NavLink
                      to={`/plan-execution-requests/${terraformWorkflowRequest?.planExecutionRequest?.planExecutionRequestId}`}
                    >
                      {
                        terraformWorkflowRequest?.planExecutionRequest
                          ?.planExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformWorkflowRequest?.planExecutionRequest?.status
                    )}
                    )
                  </td>
                  <td>
                    <NavLink
                      to={`/apply-execution-requests/${terraformWorkflowRequest?.applyExecutionRequest?.applyExecutionRequestId}`}
                    >
                      {
                        terraformWorkflowRequest?.applyExecutionRequest
                          ?.applyExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformWorkflowRequest?.applyExecutionRequest?.status
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

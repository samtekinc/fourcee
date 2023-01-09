import React, { useState } from "react";
import {
  ModulePropagationDriftCheckRequest,
  ModulePropagationDriftCheckRequests,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import { Container } from "react-bootstrap";

const MODULE_PROPAGATION_DRIFT_CHECK_QUERY = gql`
  query modulePropagationDriftCheckRequest(
    $modulePropagationId: ID!
    $modulePropagationDriftCheckRequestId: ID!
  ) {
    modulePropagationDriftCheckRequest(
      modulePropagationId: $modulePropagationId
      modulePropagationDriftCheckRequestId: $modulePropagationDriftCheckRequestId
    ) {
      modulePropagationId
      modulePropagationDriftCheckRequestId
      requestTime
      status
      terraformDriftCheckWorkflowRequests {
        items {
          terraformDriftCheckWorkflowRequestId
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
          syncStatus
        }
      }
    }
  }
`;

type Response = {
  modulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
};

export const ModulePropagationDriftCheckRequestPage = () => {
  const params = useParams();

  const modulePropagationDriftCheckRequestId =
    params.modulePropagationDriftCheckRequestId
      ? params.modulePropagationDriftCheckRequestId
      : "";
  const modulePropagationId = params.modulePropagationId
    ? params.modulePropagationId
    : "";

  const { loading, error, data } = useQuery<Response>(
    MODULE_PROPAGATION_DRIFT_CHECK_QUERY,
    {
      variables: {
        modulePropagationDriftCheckRequestId:
          modulePropagationDriftCheckRequestId,
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
        Module Propagation DriftCheck Request{" "}
        {
          data?.modulePropagationDriftCheckRequest
            .modulePropagationDriftCheckRequestId
        }
      </h1>
      <p>
        Status: {renderStatus(data?.modulePropagationDriftCheckRequest.status)}
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
            <th>Sync Status</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagationDriftCheckRequest.terraformDriftCheckWorkflowRequests.items.map(
            (terraformDriftCheckWorkflowRequest) => {
              return (
                <tr>
                  <td>
                    {
                      terraformDriftCheckWorkflowRequest?.terraformDriftCheckWorkflowRequestId
                    }
                  </td>
                  <td>
                    {terraformDriftCheckWorkflowRequest?.destroy
                      ? "Destroy"
                      : "Apply"}
                  </td>
                  <td>
                    <NavLink
                      to={`/org-accounts/${terraformDriftCheckWorkflowRequest?.moduleAssignment.orgAccount.orgAccountId}`}
                    >
                      {
                        terraformDriftCheckWorkflowRequest?.moduleAssignment
                          .orgAccount.name
                      }{" "}
                      (
                      {
                        terraformDriftCheckWorkflowRequest?.moduleAssignment
                          .orgAccount.orgAccountId
                      }
                      )
                    </NavLink>
                  </td>
                  <td>
                    {renderStatus(terraformDriftCheckWorkflowRequest?.status)}
                  </td>
                  <td>
                    <NavLink
                      to={`/plan-execution-requests/${terraformDriftCheckWorkflowRequest?.planExecutionRequest?.planExecutionRequestId}`}
                    >
                      {
                        terraformDriftCheckWorkflowRequest?.planExecutionRequest
                          ?.planExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformDriftCheckWorkflowRequest?.planExecutionRequest
                        ?.status
                    )}
                    )
                  </td>
                  <td>{terraformDriftCheckWorkflowRequest?.syncStatus}</td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};

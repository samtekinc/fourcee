import React, { useState } from "react";
import {
  ModulePropagationDriftCheckRequest,
  ModulePropagationDriftCheckRequests,
  RequestStatus,
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
      terraformDriftCheckRequests {
        items {
          terraformDriftCheckRequestId
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

  const { loading, error, data, startPolling } = useQuery<Response>(
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
          {data?.modulePropagationDriftCheckRequest.terraformDriftCheckRequests.items.map(
            (terraformDriftCheckRequest) => {
              return (
                <tr>
                  <td>
                    {terraformDriftCheckRequest?.terraformDriftCheckRequestId}
                  </td>
                  <td>
                    {terraformDriftCheckRequest?.destroy ? "Destroy" : "Apply"}
                  </td>
                  <td>
                    <NavLink
                      to={`/org-accounts/${terraformDriftCheckRequest?.moduleAssignment.orgAccount.orgAccountId}`}
                    >
                      {
                        terraformDriftCheckRequest?.moduleAssignment.orgAccount
                          .name
                      }{" "}
                      (
                      {
                        terraformDriftCheckRequest?.moduleAssignment.orgAccount
                          .orgAccountId
                      }
                      )
                    </NavLink>
                  </td>
                  <td>{renderStatus(terraformDriftCheckRequest?.status)}</td>
                  <td>
                    <NavLink
                      to={`/plan-execution-requests/${terraformDriftCheckRequest?.planExecutionRequest?.planExecutionRequestId}`}
                    >
                      {
                        terraformDriftCheckRequest?.planExecutionRequest
                          ?.planExecutionRequestId
                      }
                    </NavLink>{" "}
                    (
                    {renderStatus(
                      terraformDriftCheckRequest?.planExecutionRequest?.status
                    )}
                    )
                  </td>
                  <td>{terraformDriftCheckRequest?.syncStatus}</td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};

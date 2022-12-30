import React, { useState } from "react";
import {
  ModulePropagationExecutionRequest,
  ModulePropagationExecutionRequests,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { gql } from "../__generated__/gql";
import Table from "react-bootstrap/Table";
import { renderTimeField } from "../utils/table_rendering";
import { Container } from "react-bootstrap";

const MODULE_PROPAGATION_EXECUTION_QUERY = gql(`
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
      planExecutionRequests {
        items {
          planExecutionRequestId
          status
          requestTime
        }
      }
      applyExecutionRequests {
        items {
          applyExecutionRequestId
          status
          requestTime
        }
      }
    }
  }
`);

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
        <p>Status: {data?.modulePropagationExecutionRequest.status}</p>
      </h1>
      <h2>Plan Requests</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Plan Request ID</th>
            <th>Status</th>
            <th>Request Time</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagationExecutionRequest.planExecutionRequests.items.map(
            (planExecutionRequest) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/plan-execution-requests/${planExecutionRequest?.planExecutionRequestId}`}
                    >
                      {planExecutionRequest?.planExecutionRequestId}
                    </NavLink>
                  </td>
                  <td>{planExecutionRequest?.status}</td>
                  {renderTimeField(planExecutionRequest?.requestTime)}
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
      <h2>Apply Requests</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Apply Request ID</th>
            <th>Status</th>
            <th>Request Time</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagationExecutionRequest.applyExecutionRequests.items.map(
            (applyExecutionRequest) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/apply-execution-requests/${applyExecutionRequest?.applyExecutionRequestId}`}
                    >
                      {applyExecutionRequest?.applyExecutionRequestId}
                    </NavLink>
                  </td>
                  <td>{applyExecutionRequest?.status}</td>
                  {renderTimeField(applyExecutionRequest?.requestTime)}
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};

import React, { useState } from "react";
import {
  ModuleAccountAssociation,
  ModuleAccountAssociations,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, useMutation, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderTimeField } from "../utils/table_rendering";
import { Container } from "react-bootstrap";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const MODULE_ACCOUNT_ASSOCIATION_QUERY = gql`
  query moduleAccountAssociation(
    $modulePropagationId: ID!
    $orgAccountId: ID!
  ) {
    moduleAccountAssociation(
      modulePropagationId: $modulePropagationId
      orgAccountId: $orgAccountId
    ) {
      modulePropagationId
      modulePropagation {
        name
        moduleGroup {
          moduleGroupId
          name
        }
        moduleVersion {
          moduleVersionId
          name
        }
      }
      orgAccount {
        orgAccountId
        name
      }
      status
      planExecutionRequests {
        items {
          planExecutionRequestId
          status
          requestTime
          modulePropagationId
          modulePropagationExecutionRequestId
        }
      }
      applyExecutionRequests {
        items {
          applyExecutionRequestId
          status
          requestTime
          modulePropagationId
          modulePropagationExecutionRequestId
        }
      }
    }
  }
`;

type Response = {
  moduleAccountAssociation: ModuleAccountAssociation;
};

export const ModuleAccountAssociationPage = () => {
  const params = useParams();

  const modulePropagationId = params.modulePropagationId
    ? params.modulePropagationId
    : "";

  const orgAccountId = params.orgAccountId ? params.orgAccountId : "";

  const { loading, error, data } = useQuery<Response>(
    MODULE_ACCOUNT_ASSOCIATION_QUERY,
    {
      variables: {
        modulePropagationId: modulePropagationId,
        orgAccountId: orgAccountId,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>Module Account Association</h1>
      <p>
        <b>Account:</b>{" "}
        <NavLink
          to={`/org-accounts/${data?.moduleAccountAssociation.orgAccount.orgAccountId}`}
        >
          {data?.moduleAccountAssociation.orgAccount.name}
        </NavLink>
        <br />
        <b>Module Propagation</b>{" "}
        <NavLink
          to={`/module-propagations/${data?.moduleAccountAssociation.modulePropagationId}`}
        >
          {data?.moduleAccountAssociation.modulePropagation.name}
        </NavLink>
        <br />
        <b>Module Group:</b>{" "}
        <NavLink
          to={`/module-groups/${data?.moduleAccountAssociation.modulePropagation.moduleGroup.moduleGroupId}`}
        >
          {data?.moduleAccountAssociation.modulePropagation.moduleGroup.name}
        </NavLink>
        <br />
        <b>Module Version:</b>{" "}
        <NavLink
          to={`/module-groups/${data?.moduleAccountAssociation.modulePropagation.moduleGroup.moduleGroupId}/versions/${data?.moduleAccountAssociation.modulePropagation.moduleVersion.moduleVersionId}`}
        >
          {data?.moduleAccountAssociation.modulePropagation.moduleVersion.name}
        </NavLink>
      </p>
      <h2>Plan Requests</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Plan Request ID</th>
            <th>Module Propagation Execution Request ID</th>
            <th>Status</th>
            <th>Request Time</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleAccountAssociation.planExecutionRequests.items.map(
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
                  <td>
                    <NavLink
                      to={`/module-propagations/${planExecutionRequest?.modulePropagationId}/executions/${planExecutionRequest?.modulePropagationExecutionRequestId}`}
                    >
                      {
                        planExecutionRequest?.modulePropagationExecutionRequestId
                      }
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
          {data?.moduleAccountAssociation.applyExecutionRequests.items.map(
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

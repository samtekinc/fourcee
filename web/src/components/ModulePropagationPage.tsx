import React, { useState } from "react";
import {
  ModulePropagation,
  ModulePropagations,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, useMutation, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderTimeField } from "../utils/table_rendering";
import { Container } from "react-bootstrap";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const MODULE_PROPAGATION_QUERY = gql(`
  query modulePropagation($modulePropagationId: ID!) {
    modulePropagation(modulePropagationId: $modulePropagationId) {
      modulePropagationId
      moduleGroup {
        moduleGroupId
        name
      }
      moduleVersion {
        moduleVersionId
        name
      }
      orgUnitId
      orgUnit {
        orgUnitId
        orgDimensionId
        name
        downstreamOrgUnits {
          items {
            orgUnitId
            orgDimensionId
            name
          }
        }
      }
      modulePropagationId
      name
      description
      executionRequests {
        items {
          modulePropagationId
          modulePropagationExecutionRequestId
          requestTime
          status
        }
      }
      moduleAccountAssociations {
        items {
          modulePropagationId
          orgAccountId
          status
        }
      }
    }
  }
`);

type Response = {
  modulePropagation: ModulePropagation;
};

export const ModulePropagationPage = () => {
  const params = useParams();

  const modulePropagationId = params.modulePropagationId
    ? params.modulePropagationId
    : "";

  const { loading, error, data } = useQuery<Response>(
    MODULE_PROPAGATION_QUERY,
    {
      variables: {
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
        Module Propagation{" "}
        <b>
          <u>{data?.modulePropagation.name}</u>
        </b>{" "}
        ({data?.modulePropagation.modulePropagationId})
      </h1>
      <p>
        <b>Module Group:</b>{" "}
        <NavLink
          to={`/module-groups/${data?.modulePropagation.moduleGroup.moduleGroupId}`}
        >
          {data?.modulePropagation.moduleGroup.name}
        </NavLink>
        <br />
        <b>Module Version:</b>{" "}
        <NavLink
          to={`/module-groups/${data?.modulePropagation.moduleGroup.moduleGroupId}/versions/${data?.modulePropagation.moduleVersion.moduleVersionId}`}
        >
          {data?.modulePropagation.moduleVersion.name}
        </NavLink>
      </p>
      <h2>Execution Requests</h2>
      <ExecuteModulePropagationButton
        modulePropagationId={modulePropagationId}
      />
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Request ID</th>
            <th>Status</th>
            <th>Request Time</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagation.executionRequests.items.map(
            (executionRequest) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/module-propagations/${executionRequest?.modulePropagationId}/executions/${executionRequest?.modulePropagationExecutionRequestId}`}
                    >
                      {executionRequest?.modulePropagationExecutionRequestId}
                    </NavLink>
                  </td>
                  <td>{executionRequest?.status}</td>
                  {renderTimeField(executionRequest?.requestTime)}
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
      <h2>Associated Org Units</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Dimension Id</th>
            <th>Name</th>
            <th>Association Type</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagation.orgUnit.downstreamOrgUnits.items
            .concat([data?.modulePropagation.orgUnit])
            .map((orgUnit) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/org-dimensions/${orgUnit?.orgDimensionId}/org-units/${orgUnit?.orgUnitId}`}
                    >
                      {orgUnit?.orgUnitId}
                    </NavLink>
                  </td>
                  <td>{orgUnit?.name}</td>
                  <td>
                    {data?.modulePropagation.orgUnit.orgUnitId ==
                    orgUnit?.orgUnitId
                      ? "Direct"
                      : "Propagated"}
                  </td>
                </tr>
              );
            })}
        </tbody>
      </Table>
      <h2>Account Associations</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Account Id</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagation.moduleAccountAssociations.items.map(
            (moduleAccountAssociation) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/org-accounts/${moduleAccountAssociation?.orgAccountId}`}
                    >
                      {moduleAccountAssociation?.orgAccountId}
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-propagations/${moduleAccountAssociation?.modulePropagationId}/account-associations/${moduleAccountAssociation?.orgAccountId}`}
                    >
                      {moduleAccountAssociation?.status}
                    </NavLink>
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

const EXECUTE_MODULE_PROPAGATION_MUTATION = gql(`
  mutation createModulePropagationExecutionRequest($modulePropagationId: ID!) {
    createModulePropagationExecutionRequest(
      modulePropagationExecutionRequest: {
        modulePropagationId: $modulePropagationId
      }
    ) {
      modulePropagationExecutionRequestId
      status
    }
  }
`);

interface ExecuteModulePropagationButtonProps {
  modulePropagationId: string;
}

type ExecuteModulePropagationResponse = {
  createModulePropagationExecutionRequest: {
    modulePropagationExecutionRequestId: string;
    status: string;
  };
};

const ExecuteModulePropagationButton: React.VFC<
  ExecuteModulePropagationButtonProps
> = (props: ExecuteModulePropagationButtonProps) => {
  const [mutation, { loading }] = useMutation<ExecuteModulePropagationResponse>(
    EXECUTE_MODULE_PROPAGATION_MUTATION,
    {
      variables: {
        modulePropagationId: props.modulePropagationId,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error executing module propagation`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Initiated ${data.createModulePropagationExecutionRequest.modulePropagationExecutionRequestId}`,
          "",
          5000
        );
      },
    }
  );

  return (
    <Button
      disabled={loading}
      onClick={() => {
        mutation();
      }}
    >
      {loading ? "Submitting..." : "Execute Module Propagation"}
    </Button>
  );
};

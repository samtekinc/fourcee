import React, { useState } from "react";
import { ModuleAssignment, ModuleAssignments } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, useMutation, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import { Container } from "react-bootstrap";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { vscDarkPlus } from "react-syntax-highlighter/dist/esm/styles/prism";
import { TriggerTerraformExecutionButton } from "./TriggerTerraformExecutionButton";
import { TriggerTerraformDriftCheckButton } from "./TriggerTerraformDriftCheckButton";
import {
  renderCloudPlatform,
  renderModuleAssignmentStatus,
} from "../utils/rendering";

const MODULE_ACCOUNT_ASSOCIATION_QUERY = gql`
  query moduleAssignment($moduleAssignmentId: ID!) {
    moduleAssignment(moduleAssignmentId: $moduleAssignmentId) {
      name
      modulePropagation {
        modulePropagationId
        name
      }
      orgAccount {
        orgAccountId
        name
        cloudPlatform
      }
      moduleGroup {
        moduleGroupId
        name
      }
      moduleVersion {
        moduleVersionId
        name
      }
      status
      terraformConfiguration
      terraformExecutionWorkflowRequests(limit: 5) {
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
      terraformDriftCheckWorkflowRequests(limit: 5) {
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
  moduleAssignment: ModuleAssignment;
};

export const ModuleAssignmentPage = () => {
  const params = useParams();

  const moduleAssignmentId = params.moduleAssignmentId
    ? params.moduleAssignmentId
    : "";

  const { loading, error, data } = useQuery<Response>(
    MODULE_ACCOUNT_ASSOCIATION_QUERY,
    {
      variables: {
        moduleAssignmentId: moduleAssignmentId,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  let terraformConfiguration = data?.moduleAssignment.terraformConfiguration
    ? atob(data?.moduleAssignment.terraformConfiguration)
    : "...";

  let isPropagated = data?.moduleAssignment.modulePropagation ? true : false;

  return (
    <Container>
      <h1>
        Module Assignment{" "}
        <b>
          {isPropagated
            ? data?.moduleAssignment?.modulePropagation?.name
            : data?.moduleAssignment?.name}
        </b>
      </h1>
      <p>
        <b>Account:</b>{" "}
        <NavLink
          to={`/org-accounts/${data?.moduleAssignment.orgAccount.orgAccountId}`}
        >
          {renderCloudPlatform(data?.moduleAssignment.orgAccount.cloudPlatform)}{" "}
          {data?.moduleAssignment.orgAccount.name}
        </NavLink>
        <br />
        {isPropagated && (
          <>
            <b>Module Propagation</b>{" "}
            <NavLink
              to={`/module-propagations/${data?.moduleAssignment.modulePropagation?.modulePropagationId}`}
            >
              {data?.moduleAssignment.modulePropagation?.name}
            </NavLink>
            <br />
          </>
        )}
        <b>Module Group:</b>{" "}
        <NavLink
          to={`/module-groups/${data?.moduleAssignment.moduleGroup.moduleGroupId}`}
        >
          {data?.moduleAssignment.moduleGroup.name}
        </NavLink>
        <br />
        <b>Module Version:</b>{" "}
        <NavLink
          to={`/module-groups/${data?.moduleAssignment.moduleGroup.moduleGroupId}/versions/${data?.moduleAssignment.moduleVersion.moduleVersionId}`}
        >
          {data?.moduleAssignment.moduleVersion.name}
        </NavLink>
        <br />
        <b>Status:</b>{" "}
        {renderModuleAssignmentStatus(data?.moduleAssignment.status)}
      </p>
      <h2>Terraform Configuration</h2>
      <Container
        className="bg-vscDarkPlus"
        style={{
          overflow: "auto",
          maxHeight: "60vh",
          whiteSpace: "pre-wrap",
          textAlign: "left",
        }}
      >
        <SyntaxHighlighter language="hcl" style={vscDarkPlus}>
          {terraformConfiguration}
        </SyntaxHighlighter>
      </Container>
      <h2>Terraform Execution Workflows</h2>
      <TriggerTerraformExecutionButton
        moduleAssignmentId={moduleAssignmentId}
        destroy={false}
      />
      {"\t"}
      <TriggerTerraformExecutionButton
        moduleAssignmentId={moduleAssignmentId}
        destroy={true}
      />
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>TF Request</th>
            <th>Apply / Destroy</th>
            <th>Org Account</th>
            <th>Status</th>
            <th>Plan Request</th>
            <th>Apply Request</th>
            <th>Request Time</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleAssignment.terraformExecutionWorkflowRequests.items.map(
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
                  {renderTimeField(
                    terraformExecutionWorkflowRequest?.requestTime
                  )}
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
      <h2>Terraform Drift Check Workflows</h2>
      <TriggerTerraformDriftCheckButton
        moduleAssignmentId={moduleAssignmentId}
      />
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>TF Request</th>
            <th>Apply / Destroy</th>
            <th>Org Account</th>
            <th>Status</th>
            <th>Plan Request</th>
            <th>Sync Status</th>
            <th>Request Time</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleAssignment.terraformDriftCheckWorkflowRequests.items.map(
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
                  {renderTimeField(
                    terraformDriftCheckWorkflowRequest?.requestTime
                  )}
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};

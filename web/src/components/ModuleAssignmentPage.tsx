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
      terraformExecutionRequests(limit: 5) {
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
      terraformDriftCheckRequests(limit: 5) {
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
      pollInterval: 3000,
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
          {data?.moduleAssignment.terraformExecutionRequests.items.map(
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
                  {renderTimeField(terraformExecutionRequest?.requestTime)}
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
          {data?.moduleAssignment.terraformDriftCheckRequests.items.map(
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
                  {renderTimeField(terraformDriftCheckRequest?.requestTime)}
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};

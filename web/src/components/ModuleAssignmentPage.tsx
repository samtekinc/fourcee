import {
  ModuleAssignment,
  StateFile,
  StateVersion,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import { Breadcrumb, Col, Container, Form, Row } from "react-bootstrap";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { vscDarkPlus } from "react-syntax-highlighter/dist/esm/styles/prism";
import { TriggerTerraformExecutionButton } from "./TriggerTerraformExecutionButton";
import { TriggerTerraformDriftCheckButton } from "./TriggerTerraformDriftCheckButton";
import {
  renderApplyDestroy,
  renderCloudPlatform,
  renderModuleAssignmentStatus,
  renderSyncStatus,
} from "../utils/rendering";
import React, { useState } from "react";

const MODULE_ACCOUNT_ASSOCIATION_QUERY = gql`
  query moduleAssignment($moduleAssignmentID: ID!) {
    moduleAssignment(moduleAssignmentID: $moduleAssignmentID) {
      id
      name
      status
      terraformConfiguration
      remoteStateBucket
      remoteStateKey

      modulePropagation {
        id
        name
      }
      orgAccount {
        id
        name
        cloudPlatform
      }
      moduleGroup {
        id
        name
      }
      moduleVersion {
        id
        name
      }

      terraformExecutionRequests(limit: 5) {
        id
        status
        startedAt
        destroy
        moduleAssignment {
          id
          orgAccount {
            id
            cloudPlatform
            name
          }
        }
        planExecutionRequest {
          id
          status
          startedAt
        }
        applyExecutionRequest {
          id
          status
          startedAt
        }
      }

      terraformDriftCheckRequests(limit: 5) {
        id
        status
        startedAt
        destroy
        moduleAssignment {
          id
          orgAccount {
            id
            cloudPlatform
            name
          }
        }
        planExecutionRequest {
          id
          status
          startedAt
        }
        syncStatus
      }

      stateVersions {
        versionID
        lastModified
        isCurrent
      }
    }
  }
`;

type Response = {
  moduleAssignment: ModuleAssignment;
};

export const ModuleAssignmentPage = () => {
  const params = useParams();

  const moduleAssignmentID = params.moduleAssignmentID
    ? params.moduleAssignmentID
    : "";

  const { loading, error, data } = useQuery<Response>(
    MODULE_ACCOUNT_ASSOCIATION_QUERY,
    {
      variables: {
        moduleAssignmentID: moduleAssignmentID,
      },
      pollInterval: 3000,
    }
  );

  const [selectedStateVersionID, setSelectedStateVersionID] = useState<
    string | undefined
  >();

  if (loading) return null;
  if (error) return <div>Error</div>;

  let terraformConfiguration = data?.moduleAssignment.terraformConfiguration
    ? data?.moduleAssignment.terraformConfiguration
    : "...";

  let isPropagated = data?.moduleAssignment.modulePropagation ? true : false;

  if (selectedStateVersionID === undefined) {
    setSelectedStateVersionID(
      data?.moduleAssignment.stateVersions[0].versionID
    );
  }

  return (
    <Container style={{ paddingTop: "2rem", paddingBottom: "5rem" }} fluid>
      <Breadcrumb>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/" }}>
          Home
        </Breadcrumb.Item>
        <Breadcrumb.Item
          linkAs={NavLink}
          linkProps={{ to: "/module-assignments" }}
        >
          Assignments
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.moduleAssignment.name} ({data?.moduleAssignment.id})
        </Breadcrumb.Item>
      </Breadcrumb>

      <Row>
        <Col md={"auto"}>
          <h1>
            {isPropagated
              ? data?.moduleAssignment?.modulePropagation?.name
              : data?.moduleAssignment?.name}
          </h1>
        </Col>
      </Row>
      <p>
        <b>Account:</b>{" "}
        <NavLink to={`/org-accounts/${data?.moduleAssignment.orgAccount.id}`}>
          {renderCloudPlatform(data?.moduleAssignment.orgAccount.cloudPlatform)}{" "}
          {data?.moduleAssignment.orgAccount.name}
        </NavLink>
        <br />
        {isPropagated && (
          <>
            <b>Module Propagation</b>{" "}
            <NavLink
              to={`/module-propagations/${data?.moduleAssignment.modulePropagation?.id}`}
            >
              {data?.moduleAssignment.modulePropagation?.name}
            </NavLink>
            <br />
          </>
        )}
        <b>Module:</b>{" "}
        <NavLink to={`/module-groups/${data?.moduleAssignment.moduleGroup.id}`}>
          {data?.moduleAssignment.moduleGroup.name}
        </NavLink>
        {" / "}
        <NavLink
          to={`/module-groups/${data?.moduleAssignment.moduleGroup.id}/versions/${data?.moduleAssignment.moduleVersion.id}`}
        >
          {data?.moduleAssignment.moduleVersion.name}
        </NavLink>
        <br />
        <b>Status:</b>{" "}
        {renderModuleAssignmentStatus(data?.moduleAssignment.status)}
      </p>

      <h2>
        Terraform Execution Workflows{" "}
        <TriggerTerraformExecutionButton
          moduleAssignmentID={moduleAssignmentID}
          destroy={false}
        />
        {"\t"}
        <TriggerTerraformExecutionButton
          moduleAssignmentID={moduleAssignmentID}
          destroy={true}
        />
      </h2>
      <Row>
        <Col md={"auto"}>
          <Table hover>
            <thead>
              <tr>
                <th>Status</th>
                <th>Org Account</th>
                <th>Action</th>
                <th>Plan Request</th>
                <th>Apply Request</th>
                <th>Request Time</th>
              </tr>
            </thead>
            <tbody>
              {data?.moduleAssignment.terraformExecutionRequests.map(
                (terraformExecutionRequest) => {
                  return (
                    <tr>
                      <td>{renderStatus(terraformExecutionRequest?.status)}</td>
                      <td>
                        {renderCloudPlatform(
                          terraformExecutionRequest?.moduleAssignment.orgAccount
                            .cloudPlatform
                        )}{" "}
                        <NavLink
                          to={`/org-accounts/${terraformExecutionRequest?.moduleAssignment.orgAccount.id}`}
                        >
                          {
                            terraformExecutionRequest?.moduleAssignment
                              .orgAccount.name
                          }
                        </NavLink>
                      </td>
                      <td>
                        {renderApplyDestroy(
                          terraformExecutionRequest?.destroy ?? false
                        )}
                      </td>
                      <td>
                        <NavLink
                          to={`/plan-execution-requests/${terraformExecutionRequest?.planExecutionRequest?.id}`}
                        >
                          {renderStatus(
                            terraformExecutionRequest?.planExecutionRequest
                              ?.status
                          )}
                        </NavLink>
                      </td>
                      <td>
                        <NavLink
                          to={`/apply-execution-requests/${terraformExecutionRequest?.applyExecutionRequest?.id}`}
                        >
                          {renderStatus(
                            terraformExecutionRequest?.applyExecutionRequest
                              ?.status
                          )}
                        </NavLink>
                      </td>
                      {renderTimeField(terraformExecutionRequest?.startedAt)}
                    </tr>
                  );
                }
              )}
            </tbody>
          </Table>
        </Col>
      </Row>

      <h2>
        Terraform Drift Check Workflows{" "}
        <TriggerTerraformDriftCheckButton
          moduleAssignmentID={moduleAssignmentID}
        />
      </h2>
      <Row>
        <Col md={"auto"}>
          <Table hover>
            <thead>
              <tr>
                <th>Status</th>
                <th>Org Account</th>
                <th>Action</th>
                <th>Plan Request</th>
                <th>Sync Status</th>
                <th>Request Time</th>
              </tr>
            </thead>
            <tbody>
              {data?.moduleAssignment.terraformDriftCheckRequests.map(
                (terraformDriftCheckRequest) => {
                  return (
                    <tr>
                      <td>
                        {renderStatus(terraformDriftCheckRequest?.status)}
                      </td>
                      <td>
                        {renderCloudPlatform(
                          terraformDriftCheckRequest?.moduleAssignment
                            .orgAccount.cloudPlatform
                        )}{" "}
                        <NavLink
                          to={`/org-accounts/${terraformDriftCheckRequest?.moduleAssignment.orgAccount.id}`}
                        >
                          {
                            terraformDriftCheckRequest?.moduleAssignment
                              .orgAccount.name
                          }
                        </NavLink>
                      </td>
                      <td>
                        {renderApplyDestroy(
                          terraformDriftCheckRequest?.destroy ?? false
                        )}
                      </td>
                      <td>
                        <NavLink
                          to={`/plan-execution-requests/${terraformDriftCheckRequest?.planExecutionRequest?.id}`}
                        >
                          {renderStatus(
                            terraformDriftCheckRequest?.planExecutionRequest
                              ?.status
                          )}
                        </NavLink>
                      </td>
                      <td>
                        {renderSyncStatus(
                          terraformDriftCheckRequest?.syncStatus
                        )}
                      </td>
                      {renderTimeField(terraformDriftCheckRequest?.startedAt)}
                    </tr>
                  );
                }
              )}
            </tbody>
          </Table>
        </Col>
      </Row>

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

      <h2>State File</h2>
      <Form.Group>
        <Form.Label>Version</Form.Label>
        <Form.Select
          name="stateVersion"
          onChange={(event: React.ChangeEvent<HTMLSelectElement>) => {
            setSelectedStateVersionID(event.target.value);
          }}
        >
          {data?.moduleAssignment.stateVersions.map((stateVersion) => {
            return (
              <option
                value={stateVersion.versionID}
                key={stateVersion.versionID}
              >
                {stateVersion.lastModified}
                {stateVersion.isCurrent ? " (current)" : ""}
              </option>
            );
          })}
        </Form.Select>
      </Form.Group>
      {selectedStateVersionID && (
        <StateVersionDisplay
          bucket={data?.moduleAssignment.remoteStateBucket ?? ""}
          keyname={data?.moduleAssignment.remoteStateKey ?? ""}
          versionID={selectedStateVersionID}
        />
      )}
    </Container>
  );
};

const STATE_FILE_QUERY = gql`
  query stateFile($bucket: String!, $key: String!, $versionID: String!) {
    stateFile(bucket: $bucket, key: $key, versionID: $versionID) {
      resources {
        type
        name
        id
        attributesJSON
      }
    }
  }
`;

type StateFileResponse = {
  stateFile: StateFile;
};

interface StateVersionDisplayProps {
  bucket: string;
  keyname: string;
  versionID: string;
}

const StateVersionDisplay: React.VFC<StateVersionDisplayProps> = (
  props: StateVersionDisplayProps
) => {
  const { loading, error, data } = useQuery<StateFileResponse>(
    STATE_FILE_QUERY,
    {
      variables: {
        bucket: props.bucket,
        key: props.keyname,
        versionID: props.versionID,
      },
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <Table hover>
        <thead>
          <tr>
            <th>Type</th>
            <th>Name</th>
            <th>ID</th>
            <th>Attributes</th>
          </tr>
        </thead>
        <tbody>
          {data?.stateFile.resources.map((resource) => {
            return (
              <tr key={resource.id}>
                <td>{resource.type}</td>
                <td>{resource.name}</td>
                <td>{resource.id}</td>
                <td>
                  <pre>{JSON.stringify(resource.attributesJSON, null, 2)}</pre>
                </td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </Container>
  );
};

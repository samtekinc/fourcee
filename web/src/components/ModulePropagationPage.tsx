import React, { useState } from "react";
import {
  Maybe,
  ModulePropagation,
  ModulePropagations,
  OrganizationalDimensions,
  OrganizationalUnit,
  ModulePropagationUpdate,
  ModuleGroup,
} from "../__generated__/graphql";
import { NavLink, Outlet, useParams } from "react-router-dom";
import { useQuery, useMutation, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import {
  Breadcrumb,
  Col,
  Container,
  Form,
  ListGroup,
  Modal,
  Row,
} from "react-bootstrap";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";
import { DriftCheckModulePropagationButton } from "./TriggerModulePropagationDriftCheckButton";
import { ExecuteModulePropagationButton } from "./TriggerModulePropagationExecutionButton";
import {
  renderModuleAssignmentStatus,
  renderSyncStatus,
} from "../utils/rendering";

const MODULE_PROPAGATION_QUERY = gql`
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
        name
        orgDimension {
          orgDimensionId
          name
        }
        downstreamOrgUnits {
          items {
            orgUnitId
            name
            orgDimension {
              orgDimensionId
              name
            }
          }
        }
      }
      modulePropagationId
      name
      description
      executionRequests(limit: 5) {
        items {
          modulePropagationId
          modulePropagationExecutionRequestId
          requestTime
          status
        }
      }
      driftCheckRequests(limit: 5) {
        items {
          modulePropagationId
          modulePropagationDriftCheckRequestId
          requestTime
          status
          syncStatus
        }
      }
      moduleAssignments {
        items {
          moduleAssignmentId
          modulePropagationId
          orgAccount {
            orgAccountId
            name
          }
          status
        }
      }
    }
  }
`;

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
      pollInterval: 3000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;
  if (!data?.modulePropagation) return <div>Not found</div>;

  return (
    <Container style={{ paddingTop: "2rem" }} fluid>
      <Breadcrumb>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/" }}>
          Home
        </Breadcrumb.Item>
        <Breadcrumb.Item
          linkAs={NavLink}
          linkProps={{ to: "/module-propagations" }}
        >
          Propagations
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.modulePropagation.name} (
          {data?.modulePropagation.modulePropagationId})
        </Breadcrumb.Item>
      </Breadcrumb>

      <Row>
        <Col md={"auto"}>
          <h1>{data?.modulePropagation.name}</h1>
        </Col>
        <Col md={"auto"}>
          <Container fluid style={{ paddingTop: "0.7rem" }}>
            <UpdateModulePropagationButton
              modulePropagation={data.modulePropagation}
              key={data.modulePropagation.modulePropagationId}
            />
          </Container>
        </Col>
      </Row>
      <h5>
        Module:{" "}
        <NavLink
          to={`/module-groups/${data?.modulePropagation.moduleGroup.moduleGroupId}`}
        >
          {data?.modulePropagation.moduleGroup.name}
        </NavLink>{" "}
        /{" "}
        <NavLink
          to={`/module-groups/${data?.modulePropagation.moduleGroup.moduleGroupId}/versions/${data?.modulePropagation.moduleVersion.moduleVersionId}`}
        >
          {data?.modulePropagation.moduleVersion.name}
        </NavLink>
      </h5>
      <br />

      <Row>
        <Col md={"auto"}>
          <h2>
            Recent Executions{" "}
            <ExecuteModulePropagationButton
              modulePropagationId={modulePropagationId}
            />
          </h2>
          <Table hover>
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
                          style={({ isActive }) =>
                            isActive
                              ? {
                                  color: "blue",
                                }
                              : {
                                  color: "inherit",
                                }
                          }
                        >
                          {
                            executionRequest?.modulePropagationExecutionRequestId
                          }
                        </NavLink>
                      </td>
                      <td>{renderStatus(executionRequest?.status)}</td>
                      {renderTimeField(executionRequest?.requestTime)}
                    </tr>
                  );
                }
              )}
            </tbody>
          </Table>
          <h2>
            Recent Drift Checks{" "}
            <DriftCheckModulePropagationButton
              modulePropagationId={modulePropagationId}
            />
          </h2>
          <Table hover>
            <thead>
              <tr>
                <th>Request ID</th>
                <th>Status</th>
                <th>Sync Status</th>
                <th>Request Time</th>
              </tr>
            </thead>
            <tbody>
              {data?.modulePropagation.driftCheckRequests.items.map(
                (driftCheckRequest) => {
                  return (
                    <tr>
                      <td>
                        <NavLink
                          to={`/module-propagations/${driftCheckRequest?.modulePropagationId}/drift-checks/${driftCheckRequest?.modulePropagationDriftCheckRequestId}`}
                          style={({ isActive }) =>
                            isActive
                              ? {
                                  color: "blue",
                                }
                              : {
                                  color: "inherit",
                                }
                          }
                        >
                          {
                            driftCheckRequest?.modulePropagationDriftCheckRequestId
                          }
                        </NavLink>
                      </td>
                      <td>{renderStatus(driftCheckRequest?.status)}</td>
                      <td>{renderSyncStatus(driftCheckRequest?.syncStatus)}</td>
                      {renderTimeField(driftCheckRequest?.requestTime)}
                    </tr>
                  );
                }
              )}
            </tbody>
          </Table>
        </Col>
        <Col md={"auto"} style={{ borderRight: "3px solid gray" }} />
        <Col md={"auto"}>
          <h2>Selected Request</h2>
          <Outlet />
        </Col>
      </Row>
      <br />
      <Row>
        <Col md={"auto"}>
          <h2>Associated Org Units</h2>
          <Table hover>
            <thead>
              <tr>
                <th>Org Dimension</th>
                <th>Org Unit</th>
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
                          to={`/org-dimensions/${orgUnit?.orgDimension.orgDimensionId}`}
                        >
                          {orgUnit?.orgDimension.name}
                        </NavLink>
                      </td>
                      <td>
                        <NavLink
                          to={`/org-dimensions/${orgUnit?.orgDimension.orgDimensionId}/org-units/${orgUnit?.orgUnitId}`}
                        >
                          {orgUnit?.name}
                        </NavLink>
                      </td>
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
        </Col>
        <Col md={"auto"}></Col>
        <Col md={"auto"}>
          <h2>Account Assignments</h2>
          <Table hover>
            <thead>
              <tr>
                <th>Status</th>
                <th>Account</th>
                <th>Assignment Id</th>
              </tr>
            </thead>
            <tbody>
              {data?.modulePropagation.moduleAssignments.items.map(
                (moduleAssignment) => {
                  return (
                    <tr>
                      <td>
                        {renderModuleAssignmentStatus(moduleAssignment?.status)}
                      </td>
                      <td>
                        <NavLink
                          to={`/org-accounts/${moduleAssignment?.orgAccount.orgAccountId}`}
                        >
                          {moduleAssignment?.orgAccount.name}
                        </NavLink>
                      </td>
                      <td>
                        <NavLink
                          to={`/module-assignments/${moduleAssignment?.moduleAssignmentId}`}
                        >
                          {moduleAssignment?.moduleAssignmentId}
                        </NavLink>
                      </td>
                    </tr>
                  );
                }
              )}
            </tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};

interface UpdateModulePropagationButtonProps {
  modulePropagation: ModulePropagation;
}

const UpdateModulePropagationButton: React.VFC<
  UpdateModulePropagationButtonProps
> = (props: UpdateModulePropagationButtonProps) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <>
      <Button variant="outline-primary" size="sm" onClick={handleShow}>
        Modify
      </Button>
      <Modal show={show} onHide={handleClose}>
        <UpdateModulePropagationForm
          modulePropagation={props.modulePropagation}
          handleClose={handleClose}
        />
      </Modal>
    </>
  );
};

const MODULE_PROPAGATION_UPDATE_OPTIONS_QUERY = gql`
  query modulePropagationUpdateOptions($moduleGroupId: ID!) {
    organizationalDimensions(limit: 10000) {
      items {
        orgDimensionId
        name
        orgUnits(limit: 10000) {
          items {
            orgUnitId
            name
          }
        }
      }
    }
    moduleGroup(moduleGroupId: $moduleGroupId) {
      versions {
        items {
          moduleVersionId
          name
        }
      }
    }
  }
`;

const UPDATE_MODULE_PROPAGATION_MUTATION = gql`
  mutation updateModulePropagation(
    $modulePropagationId: ID!
    $update: ModulePropagationUpdate!
  ) {
    updateModulePropagation(
      modulePropagationId: $modulePropagationId
      update: $update
    ) {
      modulePropagationId
    }
  }
`;

type ModulePropagationUpdateOptionsResponse = {
  organizationalDimensions: OrganizationalDimensions;
  moduleGroup: ModuleGroup;
};

interface UpdateModulePropagationFormProps {
  modulePropagation: ModulePropagation;
  handleClose: () => void;
}

const UpdateModulePropagationForm: React.VFC<
  UpdateModulePropagationFormProps
> = (props: UpdateModulePropagationFormProps) => {
  const [formState, setFormState] = useState<ModulePropagationUpdate>({});
  const [orgUnits, setOrgUnits] = useState(Array<Maybe<OrganizationalUnit>>());

  const [mutation] = useMutation(UPDATE_MODULE_PROPAGATION_MUTATION, {
    variables: {
      modulePropagationId: props.modulePropagation.modulePropagationId,
      update: formState,
    },
    onError: (error) => {
      console.log(error);
      NotificationManager.error(
        error.message,
        `Error updating module propagation`,
        5000
      );
    },
    onCompleted: (data) => {
      NotificationManager.success(
        `Updated ${data.updateModulePropagation.modulePropagationId}`,
        "",
        5000
      );
    },
  });

  const { loading, error, data } =
    useQuery<ModulePropagationUpdateOptionsResponse>(
      MODULE_PROPAGATION_UPDATE_OPTIONS_QUERY,
      {
        variables: {
          moduleGroupId: props.modulePropagation.moduleGroup.moduleGroupId,
        },
      }
    );
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;
  if (!data) return <p>No data</p>;

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();

    console.log(formState);

    mutation();

    props.handleClose();
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const target = event.target;
    let value: string | null = target.value;
    if (value === "") {
      value = null;
    }
    const name = target.name;

    setFormState({
      [name]: value,
    });
  };

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    setFormState({
      [name]: value,
    });

    if (name === "orgDimensionId") {
      const orgDimension = data.organizationalDimensions.items.find(
        (item) => item?.orgDimensionId === value
      );
      setOrgUnits(orgDimension?.orgUnits.items ?? []);
    }
  };

  return (
    <>
      <Form onSubmit={handleSubmit}>
        <Modal.Header closeButton>
          <Modal.Title>Update Module Propagation</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Name</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter name"
              name="name"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Description</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter description"
              name="description"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Module Version</Form.Label>
            <Form.Select name="moduleVersionId" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select Module Version
              </option>
              {data.moduleGroup?.versions.items.map((moduleVersion) => {
                return (
                  <option
                    value={moduleVersion?.moduleVersionId}
                    key={moduleVersion?.moduleVersionId}
                    disabled={
                      props.modulePropagation.moduleVersion.moduleVersionId ===
                      moduleVersion?.moduleVersionId
                    }
                  >
                    {moduleVersion?.name}
                  </option>
                );
              })}
            </Form.Select>
          </Form.Group>
          <Form.Group>
            <Form.Label>Org Dimension</Form.Label>
            <Form.Select name="orgDimensionId" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select Org Dimension
              </option>
              {data.organizationalDimensions.items.map((orgDimension) => {
                return (
                  <option
                    value={orgDimension?.orgDimensionId}
                    key={orgDimension?.orgDimensionId}
                  >
                    {orgDimension?.name}
                  </option>
                );
              })}
            </Form.Select>
          </Form.Group>
          <Form.Group>
            <Form.Label>Org Unit</Form.Label>
            <Form.Select name="orgUnitId" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select Org Unit
              </option>
              {orgUnits.map((orgUnit) => {
                return (
                  <option value={orgUnit?.orgUnitId} key={orgUnit?.orgUnitId}>
                    {orgUnit?.name}
                  </option>
                );
              })}
            </Form.Select>
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={props.handleClose}>
            Close
          </Button>
          <Button variant="primary" type="submit">
            Submit
          </Button>
        </Modal.Footer>
      </Form>
    </>
  );
};

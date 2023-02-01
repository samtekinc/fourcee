import React, { useState } from "react";
import {
  Maybe,
  ModulePropagation,
  OrgDimension,
  OrgUnit,
  ModulePropagationUpdate,
  ModuleGroup,
} from "../__generated__/graphql";
import { NavLink, Outlet, useParams } from "react-router-dom";
import { useQuery, useMutation, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import { Breadcrumb, Col, Container, Form, Modal, Row } from "react-bootstrap";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";
import { DriftCheckModulePropagationButton } from "./TriggerModulePropagationDriftCheckButton";
import { ExecuteModulePropagationButton } from "./TriggerModulePropagationExecutionButton";
import {
  renderModuleAssignmentStatus,
  renderSyncStatus,
  timeDeltaToString,
} from "../utils/rendering";

const MODULE_PROPAGATION_QUERY = gql`
  query modulePropagation($modulePropagationID: ID!) {
    modulePropagation(modulePropagationID: $modulePropagationID) {
      id
      name
      description

      moduleGroup {
        id
        name
      }

      moduleVersion {
        id
        name
      }

      orgUnit {
        id
        name
        orgDimension {
          id
          name
        }
        downstreamOrgUnits {
          id
          name
          orgDimension {
            id
            name
          }
        }
      }

      executionRequests(limit: 5) {
        id
        modulePropagationID
        startedAt
        completedAt
        status
      }

      driftCheckRequests(limit: 5) {
        id
        modulePropagationID
        startedAt
        completedAt
        status
        syncStatus
      }

      moduleAssignments {
        id
        modulePropagationID
        orgAccount {
          id
          name
        }
        status
      }
    }
  }
`;

type Response = {
  modulePropagation: ModulePropagation;
};

export const ModulePropagationPage = () => {
  const params = useParams();

  const modulePropagationID = params.modulePropagationID
    ? params.modulePropagationID
    : "";

  const { loading, error, data } = useQuery<Response>(
    MODULE_PROPAGATION_QUERY,
    {
      variables: {
        modulePropagationID: modulePropagationID,
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
          {data?.modulePropagation.name} ({data?.modulePropagation.id})
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
              key={data.modulePropagation.id}
            />
          </Container>
        </Col>
      </Row>
      <h5>
        Module:{" "}
        <NavLink
          to={`/module-groups/${data?.modulePropagation.moduleGroup.id}`}
        >
          {data?.modulePropagation.moduleGroup.name}
        </NavLink>{" "}
        /{" "}
        <NavLink
          to={`/module-groups/${data?.modulePropagation.moduleGroup.id}/versions/${data?.modulePropagation.moduleVersion.id}`}
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
              modulePropagationID={modulePropagationID}
            />
          </h2>
          <Table hover>
            <thead>
              <tr>
                <th>Request ID</th>
                <th>Status</th>
                <th>Started</th>
                <th>Execution Time</th>
              </tr>
            </thead>
            <tbody>
              {data?.modulePropagation.executionRequests.map(
                (executionRequest) => {
                  let startedAtTime = new Date(
                    Date.parse(executionRequest?.startedAt)
                  );
                  let completedAtTime = new Date(
                    Date.parse(executionRequest?.completedAt)
                  );
                  let elapsedTime =
                    (completedAtTime.getTime() - startedAtTime.getTime()) /
                    1000;
                  return (
                    <tr>
                      <td>
                        <NavLink
                          to={`/module-propagations/${executionRequest?.modulePropagationID}/executions/${executionRequest?.id}`}
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
                          {executionRequest?.id}
                        </NavLink>
                      </td>
                      <td>{renderStatus(executionRequest?.status)}</td>
                      {renderTimeField(executionRequest?.startedAt)}
                      <td>{timeDeltaToString(elapsedTime)}</td>
                    </tr>
                  );
                }
              )}
            </tbody>
          </Table>
          <h2>
            Recent Drift Checks{" "}
            <DriftCheckModulePropagationButton
              modulePropagationID={modulePropagationID}
            />
          </h2>
          <Table hover>
            <thead>
              <tr>
                <th>Request ID</th>
                <th>Status</th>
                <th>Sync Status</th>
                <th>Started</th>
                <th>Execution Time</th>
              </tr>
            </thead>
            <tbody>
              {data?.modulePropagation.driftCheckRequests.map(
                (driftCheckRequest) => {
                  let startedAtTime = new Date(
                    Date.parse(driftCheckRequest?.startedAt)
                  );
                  let completedAtTime = new Date(
                    Date.parse(driftCheckRequest?.completedAt)
                  );
                  let elapsedTime =
                    (completedAtTime.getTime() - startedAtTime.getTime()) /
                    1000;
                  return (
                    <tr>
                      <td>
                        <NavLink
                          to={`/module-propagations/${driftCheckRequest?.modulePropagationID}/drift-checks/${driftCheckRequest?.id}`}
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
                          {driftCheckRequest?.id}
                        </NavLink>
                      </td>
                      <td>{renderStatus(driftCheckRequest?.status)}</td>
                      <td>{renderSyncStatus(driftCheckRequest?.syncStatus)}</td>
                      {renderTimeField(driftCheckRequest?.startedAt)}
                      <td>{timeDeltaToString(elapsedTime)}</td>
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
              {[data?.modulePropagation.orgUnit]
                .concat(data?.modulePropagation.orgUnit.downstreamOrgUnits)
                .map((orgUnit) => {
                  return (
                    <tr>
                      <td>
                        <NavLink
                          to={`/org-dimensions/${orgUnit?.orgDimension.id}`}
                        >
                          {orgUnit?.orgDimension.name}
                        </NavLink>
                      </td>
                      <td>
                        <NavLink
                          to={`/org-dimensions/${orgUnit?.orgDimension.id}/org-units/${orgUnit?.id}`}
                        >
                          {orgUnit?.name}
                        </NavLink>
                      </td>
                      <td>
                        {data?.modulePropagation.orgUnit.id == orgUnit?.id
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
                <th>Assignment ID</th>
              </tr>
            </thead>
            <tbody>
              {data?.modulePropagation.moduleAssignments.map(
                (moduleAssignment) => {
                  return (
                    <tr>
                      <td>
                        {renderModuleAssignmentStatus(moduleAssignment?.status)}
                      </td>
                      <td>
                        <NavLink
                          to={`/org-accounts/${moduleAssignment?.orgAccount.id}`}
                        >
                          {moduleAssignment?.orgAccount.name}
                        </NavLink>
                      </td>
                      <td>
                        <NavLink
                          to={`/module-assignments/${moduleAssignment?.id}`}
                        >
                          {moduleAssignment?.id}
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
  query modulePropagationUpdateOptions($moduleGroupID: ID!) {
    orgDimensions {
      id
      name
      orgUnits {
        id
        name
      }
    }
    moduleGroup(moduleGroupID: $moduleGroupID) {
      versions {
        id
        name
      }
    }
  }
`;

const UPDATE_MODULE_PROPAGATION_MUTATION = gql`
  mutation updateModulePropagation(
    $modulePropagationID: ID!
    $update: ModulePropagationUpdate!
  ) {
    updateModulePropagation(
      modulePropagationID: $modulePropagationID
      update: $update
    ) {
      id
    }
  }
`;

type ModulePropagationUpdateOptionsResponse = {
  orgDimensions: OrgDimension[];
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
  const [orgUnits, setOrgUnits] = useState(Array<Maybe<OrgUnit>>());

  const [mutation] = useMutation(UPDATE_MODULE_PROPAGATION_MUTATION, {
    variables: {
      modulePropagationID: props.modulePropagation.id,
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
        `Updated ${data.updateModulePropagation.modulePropagationID}`,
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
          moduleGroupID: props.modulePropagation.moduleGroup.id,
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

    if (name === "orgDimensionID") {
      console.log(value);
      const orgDimension = data.orgDimensions.find(
        (item) => item?.id.toString() === value
      );
      console.log(data?.orgDimensions);
      console.log(orgDimension);
      setOrgUnits(orgDimension?.orgUnits ?? []);
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
            <Form.Select name="moduleVersionID" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select Module Version
              </option>
              {data.moduleGroup?.versions.map((moduleVersion) => {
                return (
                  <option
                    value={moduleVersion?.id}
                    key={moduleVersion?.id}
                    disabled={
                      props.modulePropagation.moduleVersion.id ===
                      moduleVersion?.id
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
            <Form.Select name="orgDimensionID" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select Org Dimension
              </option>
              {data.orgDimensions.map((orgDimension) => {
                return (
                  <option value={orgDimension?.id} key={orgDimension?.id}>
                    {orgDimension?.name}
                  </option>
                );
              })}
            </Form.Select>
          </Form.Group>
          <Form.Group>
            <Form.Label>Org Unit</Form.Label>
            <Form.Select name="orgUnitID" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select Org Unit
              </option>
              {orgUnits.map((orgUnit) => {
                return (
                  <option value={orgUnit?.id} key={orgUnit?.id}>
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

import React, { useState } from "react";
import {
  Maybe,
  ModulePropagation,
  ModulePropagations,
  OrganizationalDimensions,
  OrganizationalUnit,
  ModulePropagationUpdate,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, useMutation, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { renderStatus, renderTimeField } from "../utils/table_rendering";
import { Container, Form, Modal } from "react-bootstrap";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

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
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;
  if (!data?.modulePropagation) return <div>Not found</div>;

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
      <UpdateModulePropagationButton
        modulePropagation={data.modulePropagation}
      />
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
                  <td>{renderStatus(executionRequest?.status)}</td>
                  {renderTimeField(executionRequest?.requestTime)}
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
      <h2>Drift Check Requests</h2>
      <DriftCheckModulePropagationButton
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
          {data?.modulePropagation.driftCheckRequests.items.map(
            (driftCheckRequest) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/module-propagations/${driftCheckRequest?.modulePropagationId}/drift-checks/${driftCheckRequest?.modulePropagationDriftCheckRequestId}`}
                    >
                      {driftCheckRequest?.modulePropagationDriftCheckRequestId}
                    </NavLink>
                  </td>
                  <td>{renderStatus(driftCheckRequest?.status)}</td>
                  {renderTimeField(driftCheckRequest?.requestTime)}
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
            <th>Org Account</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {data?.modulePropagation.moduleAssignments.items.map(
            (moduleAssignment) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/org-accounts/${moduleAssignment?.orgAccount.orgAccountId}`}
                    >
                      {moduleAssignment?.orgAccount.name} (
                      {moduleAssignment?.orgAccount.orgAccountId})
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-assignments/${moduleAssignment?.moduleAssignmentId}`}
                    >
                      {moduleAssignment?.status}
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

const DRIFT_CHECK_MODULE_PROPAGATION_MUTATION = gql(`
  mutation createModulePropagationDriftCheckRequest($modulePropagationId: ID!) {
    createModulePropagationDriftCheckRequest(
      modulePropagationDriftCheckRequest: {
        modulePropagationId: $modulePropagationId
      }
    ) {
      modulePropagationDriftCheckRequestId
      status
    }
  }
`);

interface DriftCheckModulePropagationButtonProps {
  modulePropagationId: string;
}

type DriftCheckModulePropagationResponse = {
  createModulePropagationDriftCheckRequest: {
    modulePropagationDriftCheckRequestId: string;
    status: string;
  };
};

const DriftCheckModulePropagationButton: React.VFC<
  DriftCheckModulePropagationButtonProps
> = (props: DriftCheckModulePropagationButtonProps) => {
  const [mutation, { loading }] =
    useMutation<DriftCheckModulePropagationResponse>(
      DRIFT_CHECK_MODULE_PROPAGATION_MUTATION,
      {
        variables: {
          modulePropagationId: props.modulePropagationId,
        },
        onError: (error) => {
          console.log(error);
          NotificationManager.error(
            error.message,
            `Error running drift check on module propagation`,
            5000
          );
        },
        onCompleted: (data) => {
          NotificationManager.success(
            `Initiated ${data.createModulePropagationDriftCheckRequest.modulePropagationDriftCheckRequestId}`,
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
      {loading ? "Submitting..." : "DriftCheck Module Propagation"}
    </Button>
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
      <Button variant="primary" onClick={handleShow}>
        Update Module Propagation
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

const ORG_DIMENSIONS_QUERY = gql`
  query orgDimensions {
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

type OrgDimensionResposne = {
  organizationalDimensions: OrganizationalDimensions;
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
    useQuery<OrgDimensionResposne>(ORG_DIMENSIONS_QUERY);
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

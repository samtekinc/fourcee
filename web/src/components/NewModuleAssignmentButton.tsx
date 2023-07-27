import React, { useEffect, useState } from "react";

import { useMutation, gql, useQuery } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Col, Form, Modal, Row } from "react-bootstrap";
import {
  Maybe,
  NewModuleAssignment,
  OrgAccount,
  ModuleGroup,
  ModuleVersion,
  AwsProviderConfigurationInput,
  GcpProviderConfigurationInput,
  ArgumentInput,
} from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";
import { renderCloudPlatform } from "../utils/rendering";

const NEW_MODULE_ASSIGNMENT_MUTATION = gql`
  mutation createModuleAssignment($moduleAssignment: NewModuleAssignment!) {
    createModuleAssignment(moduleAssignment: $moduleAssignment) {
      id
    }
  }
`;

const MODULE_ASSIGNMENT_OPTIONS_QUERY = gql`
  query moduleAssignmentOptions {
    orgAccounts {
      id
      name
      cloudPlatform
      cloudIdentifier
    }
    moduleGroups {
      id
      name
      cloudPlatform
      versions {
        id
        name
        variables {
          name
          type
          default
          description
        }
      }
    }
  }
`;

type CreateModuleAssignmentResponse = {
  createModuleAssignment: {
    moduleAssignmentID: string;
  };
};

type ModuleAssignmentOptionsResponse = {
  orgAccounts: OrgAccount[];
  moduleGroups: ModuleGroup[];
};

type NewModuleAssignmentButtonProps = {
  onCompleted: () => void;
};

export const NewModuleAssignmentButton: React.VFC<
  NewModuleAssignmentButtonProps
> = (props: NewModuleAssignmentButtonProps) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <>
      <Button variant="primary" onClick={handleShow}>
        <BsPlusCircle /> Create New Module Assignment
      </Button>
      <Modal show={show} onHide={handleClose} size="xl">
        <NewModuleAssignmentForm
          handleClose={handleClose}
          onCompleted={props.onCompleted}
        />
      </Modal>
    </>
  );
};

interface NewModuleAssignmentFormProps {
  handleClose: () => void;
  onCompleted: () => void;
}

export const NewModuleAssignmentForm: React.VFC<
  NewModuleAssignmentFormProps
> = (props: NewModuleAssignmentFormProps) => {
  const [formState, setFormState] = useState<NewModuleAssignment>(
    {} as NewModuleAssignment
  );

  const [awsProviders, setAwsProviders] = useState<
    Array<AwsProviderConfigurationInput>
  >([]);
  const [gcpProviders, setGcpProviders] = useState<
    Array<GcpProviderConfigurationInput>
  >([]);
  const [moduleArguments, setModuleArguments] = useState<
    Map<string, ArgumentInput>
  >(new Map());

  const [moduleGroup, setModuleGroup] = useState<Maybe<ModuleGroup>>(
    null as Maybe<ModuleGroup>
  );

  const [moduleVersion, setModuleVersion] = useState<Maybe<ModuleVersion>>(
    null as Maybe<ModuleVersion>
  );

  const [numAwsProviders, setNumAwsProviders] = useState(0);
  const [numGcpProviders, setNumGcpProviders] = useState(0);

  const [mutation] = useMutation<CreateModuleAssignmentResponse>(
    NEW_MODULE_ASSIGNMENT_MUTATION,
    {
      variables: {
        moduleAssignment: formState,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error creating module assignment`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Created module assignment ${data?.createModuleAssignment?.moduleAssignmentID}`,
          "",
          5000
        );
      },
    }
  );

  const { loading, error, data, refetch } =
    useQuery<ModuleAssignmentOptionsResponse>(
      MODULE_ASSIGNMENT_OPTIONS_QUERY,
      {}
    );

  useEffect(() => {
    // I'm probably using this wrong, but intent is to refetch the org accounts when the org structure or unit changes
    refetch();
  });

  if (loading) return null;
  if (error) return <div>Error</div>;

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();

    formState.awsProviderConfigurations = awsProviders;
    formState.gcpProviderConfigurations = gcpProviders;

    const args: Array<ArgumentInput> = [];
    moduleArguments.forEach((value, key) => {
      if (value.value !== "") args.push(value);
    });
    formState.arguments = args;

    console.log(formState);
    mutation();

    setTimeout(props.onCompleted, 1000);

    props.handleClose();
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const target = event.target;
    let value: string | null = target.value;
    if (value === "") {
      value = null;
    }
    const name = target.name;

    if (name.startsWith("awsProvider")) {
      const providerIndex = Number(name.split(":")[1]);
      const fieldName = name.split(":")[2];

      setAwsProviders((prevState) => {
        const newState = [...prevState];
        if (fieldName === "Region") {
          newState[providerIndex].region = value ?? "";
        } else if (fieldName === "Alias") {
          newState[providerIndex].alias = value ?? "";
        }
        return newState;
      });
    } else if (name.startsWith("variable")) {
      const variableIndex = Number(name.split(":")[1]);
      const variableName = name.split(":")[2];
      setModuleArguments((prevState) => {
        const newState = new Map(prevState);
        newState.set(variableName, {
          name: variableName,
          value: value ?? "",
        } as ArgumentInput);
        return newState;
      });
    } else {
      setFormState((prevState) => {
        return {
          ...prevState,
          [name]: value,
        } as NewModuleAssignment;
      });
    }
  };

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    if (name === "moduleGroupID") {
      setModuleGroup(
        data?.moduleGroups?.find((moduleGroup) => moduleGroup?.id === value) ??
          null
      );
      setModuleVersion(null as Maybe<ModuleVersion>);
      setModuleArguments(new Map());
    } else if (name === "moduleVersionID") {
      setModuleVersion(
        moduleGroup?.versions?.find(
          (moduleVersion) => moduleVersion?.id === value
        ) ?? null
      );
      setModuleArguments(new Map());
    }

    setFormState((prevState) => {
      return {
        ...prevState,
        [name]: value,
      } as NewModuleAssignment;
    });
  };

  return (
    <>
      <Form onSubmit={handleSubmit}>
        <Modal.Header closeButton>
          <Modal.Title>New Module Assignment</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Name</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter module assignment name"
              name="name"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Description</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter module assignment description"
              name="description"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group>
            <Form.Label>Org Account</Form.Label>
            <Form.Select name="orgAccountID" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select an Org Account
              </option>
              {data?.orgAccounts.map((orgAccount) => {
                return (
                  <option value={orgAccount?.id} key={orgAccount?.id}>
                    {orgAccount?.name}
                  </option>
                );
              })}
            </Form.Select>
          </Form.Group>
          <Form.Group>
            <Form.Label>Module Group</Form.Label>
            <Form.Select name="moduleGroupID" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select a Module Group
              </option>
              {data?.moduleGroups.map((moduleGroup) => {
                return (
                  <option value={moduleGroup?.id} key={moduleGroup?.id}>
                    {renderCloudPlatform(moduleGroup?.cloudPlatform)}{" "}
                    {moduleGroup?.name}
                  </option>
                );
              })}
            </Form.Select>
          </Form.Group>
          <Form.Group>
            <Form.Label>Module Version</Form.Label>
            <Form.Select
              name="moduleVersionID"
              onChange={handleSelectChange}
              key={moduleGroup?.id}
            >
              <option selected={true} disabled={true}>
                Select a Module Version
              </option>
              {moduleGroup?.versions?.map((moduleVersion) => {
                return (
                  <option value={moduleVersion?.id} key={moduleVersion?.id}>
                    {moduleVersion?.name}
                  </option>
                );
              })}
            </Form.Select>
          </Form.Group>
          <h3>AWS Providers</h3>
          <Button
            onClick={() => {
              setNumAwsProviders(numAwsProviders + 1);
              setAwsProviders((prevState) => {
                return [...prevState, {} as AwsProviderConfigurationInput];
              });
            }}
          >
            Add AWS Provider
          </Button>{" "}
          <Button
            variant="danger"
            onClick={() => {
              if (numAwsProviders >= 1) {
                setNumAwsProviders(numAwsProviders - 1);
                setAwsProviders((prevState) => {
                  return prevState.slice(0, -1);
                });
              }
            }}
          >
            Remove Provider
          </Button>
          {[...Array(numAwsProviders)].map((_, index) => (
            <Row>
              <Form.Group as={Col}>
                <Form.Label>Region</Form.Label>
                <Form.Control
                  name={`awsProvider:${index}:Region`}
                  type="text"
                  onChange={handleInputChange}
                />
              </Form.Group>
              <Form.Group as={Col}>
                <Form.Label>Alias</Form.Label>
                <Form.Control
                  name={`awsProvider:${index}:Alias`}
                  type="text"
                  onChange={handleInputChange}
                />
              </Form.Group>
            </Row>
          ))}
          <h3>Arguments</h3>
          {moduleVersion?.variables?.map((variable, index) => {
            return (
              <Form.Group>
                <Form.Label>{variable?.name}</Form.Label>
                <Form.Control
                  name={`variable:${index}:${variable?.name}`}
                  type="text"
                  onChange={handleInputChange}
                  placeholder={variable?.default?.toString()}
                />
              </Form.Group>
            );
          })}
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

import React, { useEffect, useState } from "react";

import { useMutation, gql, useQuery } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Col, FloatingLabel, Form, Modal, Row } from "react-bootstrap";
import {
  Maybe,
  NewOrganizationalUnit,
  NewModulePropagation,
  OrganizationalAccount,
  OrganizationalAccounts,
  OrganizationalDimension,
  OrganizationalDimensions,
  OrganizationalUnit,
  ModuleGroups,
  ModuleGroup,
  ModuleVersion,
  AwsProviderConfigurationInput,
  GcpProviderConfigurationInput,
  ArgumentInput,
} from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";
import { renderCloudPlatform } from "../utils/rendering";

const NEW_MODULE_PROPAGATION_MUTATION = gql`
  mutation createModulePropagation($modulePropagation: NewModulePropagation!) {
    createModulePropagation(modulePropagation: $modulePropagation) {
      modulePropagationId
    }
  }
`;

const MODULE_PROPAGATION_OPTIONS_QUERY = gql`
  query modulePropagationOptions {
    organizationalDimensions {
      items {
        orgDimensionId
        name
        orgUnits {
          items {
            orgUnitId
            name
          }
        }
      }
    }
    moduleGroups {
      items {
        moduleGroupId
        name
        cloudPlatform
        versions {
          items {
            moduleVersionId
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
    }
  }
`;

type CreateModulePropagationResponse = {
  createModulePropagation: {
    modulePropagationId: string;
  };
};

type ModulePropagationOptionsResponse = {
  organizationalDimensions: OrganizationalDimensions;
  moduleGroups: ModuleGroups;
};

type NewModulePropagationButtonProps = {
  onCompleted: () => void;
};

export const NewModulePropagationButton: React.VFC<
  NewModulePropagationButtonProps
> = (props: NewModulePropagationButtonProps) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <>
      <Button variant="primary" onClick={handleShow}>
        <BsPlusCircle /> Create New Module Propagation
      </Button>
      <Modal show={show} onHide={handleClose} size="xl">
        <NewModulePropagationForm
          handleClose={handleClose}
          onCompleted={props.onCompleted}
        />
      </Modal>
    </>
  );
};

interface NewModulePropagationFormProps {
  handleClose: () => void;
  onCompleted: () => void;
}

export const NewModulePropagationForm: React.VFC<
  NewModulePropagationFormProps
> = (props: NewModulePropagationFormProps) => {
  const [formState, setFormState] = useState<NewModulePropagation>(
    {} as NewModulePropagation
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

  const [orgDimension, setOrgDimension] = useState<
    Maybe<OrganizationalDimension>
  >(null as Maybe<OrganizationalDimension>);

  const [moduleGroup, setModuleGroup] = useState<Maybe<ModuleGroup>>(
    null as Maybe<ModuleGroup>
  );

  const [moduleVersion, setModuleVersion] = useState<Maybe<ModuleVersion>>(
    null as Maybe<ModuleVersion>
  );

  const [numAwsProviders, setNumAwsProviders] = useState(0);
  const [numGcpProviders, setNumGcpProviders] = useState(0);

  const [mutation] = useMutation<CreateModulePropagationResponse>(
    NEW_MODULE_PROPAGATION_MUTATION,
    {
      variables: {
        modulePropagation: formState,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error creating org unit membership`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Created module propagation ${data?.createModulePropagation?.modulePropagationId}`,
          "",
          5000
        );
      },
    }
  );

  const { loading, error, data, refetch } =
    useQuery<ModulePropagationOptionsResponse>(
      MODULE_PROPAGATION_OPTIONS_QUERY,
      {}
    );

  useEffect(() => {
    // I'm probably using this wrong, but intent is to refetch the org accounts when the org dimension or unit changes
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
        } as NewModulePropagation;
      });
    }
  };

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    if (name === "orgDimensionId") {
      setOrgDimension(
        data?.organizationalDimensions?.items?.find(
          (dimension) => dimension?.orgDimensionId === value
        ) ?? null
      );
    } else if (name === "moduleGroupId") {
      setModuleGroup(
        data?.moduleGroups?.items?.find(
          (moduleGroup) => moduleGroup?.moduleGroupId === value
        ) ?? null
      );
      setModuleVersion(null as Maybe<ModuleVersion>);
      setModuleArguments(new Map());
    } else if (name === "moduleVersionId") {
      setModuleVersion(
        moduleGroup?.versions?.items?.find(
          (moduleVersion) => moduleVersion?.moduleVersionId === value
        ) ?? null
      );
      setModuleArguments(new Map());
    }

    setFormState((prevState) => {
      return {
        ...prevState,
        [name]: value,
      } as NewModulePropagation;
    });
  };

  return (
    <>
      <Form onSubmit={handleSubmit}>
        <Modal.Header closeButton>
          <Modal.Title>New Module Propagation</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Row>
            <Col md={5}>
              <Form.Group>
                <Form.Label>Name</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Enter module propagation name"
                  name="name"
                  onChange={handleInputChange}
                />
              </Form.Group>
            </Col>
            <Col md={7}>
              <Form.Group>
                <Form.Label>Description</Form.Label>
                <Form.Control
                  as="textarea"
                  rows={3}
                  name="description"
                  onChange={handleInputChange}
                />
              </Form.Group>
            </Col>
          </Row>
          <br />
          <Row>
            <Col md={6}>
              <Form.Group>
                <Form.Label>Org Dimension</Form.Label>
                <Form.Select
                  name="orgDimensionId"
                  onChange={handleSelectChange}
                >
                  <option selected={true} disabled={true}>
                    Select an Org Dimension
                  </option>
                  {data?.organizationalDimensions.items?.map((orgDimension) => {
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
            </Col>
            <Col md={6}>
              <Form.Group>
                <Form.Label>Org Unit</Form.Label>
                <Form.Select name="orgUnitId" onChange={handleSelectChange}>
                  <option selected={true} disabled={true}>
                    Select an Org Unit
                  </option>
                  {orgDimension?.orgUnits?.items.map((orgUnit) => {
                    return (
                      <option
                        value={orgUnit?.orgUnitId}
                        key={orgUnit?.orgUnitId}
                      >
                        {orgUnit?.name}
                      </option>
                    );
                  })}
                </Form.Select>
              </Form.Group>
            </Col>
          </Row>
          <br />
          <Row>
            <Col md={6}>
              <Form.Group>
                <Form.Label>Module Group</Form.Label>
                <Form.Select name="moduleGroupId" onChange={handleSelectChange}>
                  <option selected={true} disabled={true}>
                    Select a Module Group
                  </option>
                  {data?.moduleGroups.items?.map((moduleGroup) => {
                    return (
                      <option
                        value={moduleGroup?.moduleGroupId}
                        key={moduleGroup?.moduleGroupId}
                      >
                        {renderCloudPlatform(moduleGroup?.cloudPlatform)}{" "}
                        {moduleGroup?.name}
                      </option>
                    );
                  })}
                </Form.Select>
              </Form.Group>
            </Col>
            <Col md={6}>
              <Form.Group>
                <Form.Label>Module Version</Form.Label>
                <Form.Select
                  name="moduleVersionId"
                  onChange={handleSelectChange}
                  key={moduleGroup?.moduleGroupId}
                >
                  <option selected={true} disabled={true}>
                    Select a Module Version
                  </option>
                  {moduleGroup?.versions?.items.map((moduleVersion) => {
                    return (
                      <option
                        value={moduleVersion?.moduleVersionId}
                        key={moduleVersion?.moduleVersionId}
                      >
                        {moduleVersion?.name}
                      </option>
                    );
                  })}
                </Form.Select>
              </Form.Group>
            </Col>
          </Row>
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
          <Row>
            {moduleVersion?.variables?.map((variable, index) => {
              return (
                <Col
                  md={6}
                  style={{
                    border: "1px solid black",
                    padding: "0.5rem 0.5rem 0.5rem 0.5rem",
                  }}
                >
                  <Form.Group>
                    <Form.Label>
                      {variable?.name} - {variable?.type}
                      <br />
                      <span style={{ fontSize: "x-small" }}>
                        {variable?.description}
                      </span>
                    </Form.Label>
                    <Form.Control
                      as={
                        variable?.type === "bool" ||
                        variable?.type === "string" ||
                        variable?.type === "number"
                          ? undefined
                          : "textarea"
                      }
                      name={`variable:${index}:${variable?.name}`}
                      type="text"
                      onChange={handleInputChange}
                      placeholder={variable?.default?.toString()}
                    />
                  </Form.Group>
                </Col>
              );
            })}
          </Row>
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

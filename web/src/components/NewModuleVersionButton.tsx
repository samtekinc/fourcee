import React, { useState } from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import { CloudPlatform, NewModuleVersion } from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";

const NEW_MODULE_VERSION_MUTATION = gql`
  mutation createModuleVersion($moduleVersion: NewModuleVersion!) {
    createModuleVersion(moduleVersion: $moduleVersion) {
      moduleVersionId
    }
  }
`;

type CreateModuleVersionResponse = {
  createModuleVersion: {
    moduleVersionId: string;
  };
};

interface NewModuleVersionButtonProps {
  moduleGroupId: string | undefined;
}

export const NewModuleVersionButton: React.VFC<NewModuleVersionButtonProps> = (
  props: NewModuleVersionButtonProps
) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const [formState, setFormState] = useState<NewModuleVersion>({
    moduleGroupId: props.moduleGroupId,
  } as NewModuleVersion);

  const [mutation] = useMutation<CreateModuleVersionResponse>(
    NEW_MODULE_VERSION_MUTATION,
    {
      variables: {
        moduleVersion: formState,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error creating module version`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Created ${data.createModuleVersion.moduleVersionId}`,
          "",
          5000
        );
      },
    }
  );

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();

    mutation();

    handleClose();
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const target = event.target;
    let value: string | null = target.value;
    if (value === "") {
      value = null;
    }
    const name = target.name;

    setFormState((prevState) => {
      return {
        ...prevState,
        [name]: value,
      } as NewModuleVersion;
    });
  };

  return (
    <>
      <Button variant="primary" onClick={handleShow}>
        <BsPlusCircle /> New Version
      </Button>
      <Modal show={show} onHide={handleClose}>
        <Form onSubmit={handleSubmit}>
          <Modal.Header closeButton>
            <Modal.Title>New Module Version</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <Form.Group>
              <Form.Label>Module Version Name</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter module group name"
                name="name"
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group>
              <Form.Label>Remote Source</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter remote source URL"
                name="remoteSource"
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group>
              <Form.Label>Terraform Version</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter terraform version"
                name="terraformVersion"
                onChange={handleInputChange}
              />
            </Form.Group>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={handleClose}>
              Close
            </Button>
            <Button variant="primary" type="submit">
              Submit
            </Button>
          </Modal.Footer>
        </Form>
      </Modal>
    </>
  );
};

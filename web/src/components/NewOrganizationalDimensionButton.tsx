import React, { useState } from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import { NewOrganizationalDimension } from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";

const NEW_ORG_DIMENSION_MUTATION = gql`
  mutation createOrganizationalDimension(
    $orgDimension: NewOrganizationalDimension!
  ) {
    createOrganizationalDimension(orgDimension: $orgDimension) {
      orgDimensionId
    }
  }
`;

type CreateOrgDimensionResponse = {
  createOrganizationalDimension: {
    orgDimensionId: string;
  };
};

export const NewOrganizationalDimensionButton: React.VFC = () => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const [formState, setFormState] = useState<NewOrganizationalDimension>(
    {} as NewOrganizationalDimension
  );

  const [mutation] = useMutation<CreateOrgDimensionResponse>(
    NEW_ORG_DIMENSION_MUTATION,
    {
      variables: {
        orgDimension: formState,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error creating org dimension`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Created ${data.createOrganizationalDimension.orgDimensionId}`,
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
      } as NewOrganizationalDimension;
    });
  };

  return (
    <>
      <Button variant="primary" onClick={handleShow}>
        <BsPlusCircle /> New
      </Button>
      <Modal show={show} onHide={handleClose}>
        <Form onSubmit={handleSubmit}>
          <Modal.Header closeButton>
            <Modal.Title>New Organizational Dimension</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <Form.Group>
              <Form.Label>Org Dimension Name</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter org dimension name"
                name="name"
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

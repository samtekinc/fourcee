import React, { useState } from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import { NewOrgDimension } from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";

const NEW_ORG_DIMENSION_MUTATION = gql`
  mutation createOrgDimension($orgDimension: NewOrgDimension!) {
    createOrgDimension(orgDimension: $orgDimension) {
      id
    }
  }
`;

type CreateOrgDimensionResponse = {
  createOrgDimension: {
    id: number;
  };
};

type NewOrgDimensionButtonProps = {
  onCompleted: () => void;
};

export const NewOrgDimensionButton: React.VFC<NewOrgDimensionButtonProps> = (
  props: NewOrgDimensionButtonProps
) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const [formState, setFormState] = useState<NewOrgDimension>(
    {} as NewOrgDimension
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
          `Error creating org structure`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Created ${data.createOrgDimension.id}`,
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
    setTimeout(props.onCompleted, 1000);
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
      } as NewOrgDimension;
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
            <Modal.Title>New Org Dimension</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <Form.Group>
              <Form.Label>Org Dimension Name</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter org structure name"
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

import React, { useState } from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import { Maybe, NewOrgUnit, OrgUnit } from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";

const NEW_ORG_UNIT_MUTATION = gql`
  mutation createOrgUnit($orgUnit: NewOrgUnit!) {
    createOrgUnit(orgUnit: $orgUnit) {
      id
    }
  }
`;

type CreateOrgUnitResponse = {
  createOrgUnit: {
    id: string;
  };
};

interface NewOrgUnitButtonProps {
  orgDimensionID: string;
  existingOrgUnits: Maybe<OrgUnit>[];
  onCompleted: () => void;
}

export const NewOrgUnitButton: React.VFC<NewOrgUnitButtonProps> = (
  props: NewOrgUnitButtonProps
) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const [formState, setFormState] = useState<NewOrgUnit>({} as NewOrgUnit);

  const [mutation] = useMutation<CreateOrgUnitResponse>(NEW_ORG_UNIT_MUTATION, {
    variables: {
      orgUnit: formState,
    },
    onError: (error) => {
      console.log(error);
      NotificationManager.error(error.message, `Error creating org unit`, 5000);
    },
    onCompleted: (data) => {
      NotificationManager.success(`Created ${data.createOrgUnit.id}`, "", 5000);
    },
  });

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();

    formState.orgDimensionID = props.orgDimensionID;

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
      } as NewOrgUnit;
    });
  };

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    setFormState((prevState) => {
      return {
        ...prevState,
        [name]: value,
      } as NewOrgUnit;
    });
  };

  return (
    <>
      <Button variant="primary" onClick={handleShow}>
        <BsPlusCircle /> Create New Org Unit
      </Button>
      <Modal show={show} onHide={handleClose}>
        <Form onSubmit={handleSubmit}>
          <Modal.Header closeButton>
            <Modal.Title>New Org Unit</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <Form.Group>
              <Form.Label>Org Unit Name</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter org unit name"
                name="name"
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group>
              <Form.Label>Parent Org unit</Form.Label>
              <Form.Select name="parentOrgUnitID" onChange={handleSelectChange}>
                <option selected={true} disabled={true}>
                  Select Org Unit
                </option>
                {props.existingOrgUnits.map((orgUnit) => {
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

import React, { useState } from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import {
  Maybe,
  NewOrganizationalUnit,
  OrganizationalUnit,
} from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";

const NEW_ORG_UNIT_MUTATION = gql`
  mutation createOrganizationalUnit($orgUnit: NewOrganizationalUnit!) {
    createOrganizationalUnit(orgUnit: $orgUnit) {
      orgUnitId
    }
  }
`;

type CreateOrgUnitResponse = {
  createOrganizationalUnit: {
    orgUnitId: string;
  };
};

interface NewOrganizationalUnitButtonProps {
  orgDimensionId: string;
  existingOrgUnits: Maybe<OrganizationalUnit>[];
  onCompleted: () => void;
}

export const NewOrganizationalUnitButton: React.VFC<
  NewOrganizationalUnitButtonProps
> = (props: NewOrganizationalUnitButtonProps) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const [formState, setFormState] = useState<NewOrganizationalUnit>(
    {} as NewOrganizationalUnit
  );

  const [mutation] = useMutation<CreateOrgUnitResponse>(NEW_ORG_UNIT_MUTATION, {
    variables: {
      orgUnit: formState,
    },
    onError: (error) => {
      console.log(error);
      NotificationManager.error(error.message, `Error creating org unit`, 5000);
    },
    onCompleted: (data) => {
      NotificationManager.success(
        `Created ${data.createOrganizationalUnit.orgUnitId}`,
        "",
        5000
      );
    },
  });

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();

    formState.orgDimensionId = props.orgDimensionId;

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
      } as NewOrganizationalUnit;
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
      } as NewOrganizationalUnit;
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
            <Modal.Title>New Organizational Unit</Modal.Title>
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
              <Form.Select name="parentOrgUnitId" onChange={handleSelectChange}>
                <option selected={true} disabled={true}>
                  Select Org Unit
                </option>
                {props.existingOrgUnits.map((orgUnit) => {
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

import React, { useState } from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import { CloudPlatform, NewOrgAccount } from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";

const NEW_ORG_ACCOUNT_MUTATION = gql`
  mutation createOrgAccount($orgAccount: NewOrgAccount!) {
    createOrgAccount(orgAccount: $orgAccount) {
      id
    }
  }
`;

type CreateOrgAccountResponse = {
  createOrgAccount: {
    id: number;
  };
};

type NewOrgAccountButtonProps = {
  onCompleted: () => void;
};

export const NewOrgAccountButton: React.VFC<NewOrgAccountButtonProps> = (
  props: NewOrgAccountButtonProps
) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const [formState, setFormState] = useState<NewOrgAccount>(
    {} as NewOrgAccount
  );

  const [mutation] = useMutation<CreateOrgAccountResponse>(
    NEW_ORG_ACCOUNT_MUTATION,
    {
      variables: {
        orgAccount: formState,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error creating org account`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Created ${data.createOrgAccount.id}`,
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
      } as NewOrgAccount;
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
      } as NewOrgAccount;
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
            <Modal.Title>New Org Account</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <Form.Group>
              <Form.Label>Org Account Name</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter org account name"
                name="name"
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group>
              <Form.Label>Cloud Platform</Form.Label>
              <Form.Select name="cloudPlatform" onChange={handleSelectChange}>
                <option selected={true} disabled={true}>
                  Select a cloud platform
                </option>
                <option value={CloudPlatform.Aws} key={CloudPlatform.Aws}>
                  {CloudPlatform.Aws}
                </option>
                <option value={CloudPlatform.Azure} key={CloudPlatform.Azure}>
                  {CloudPlatform.Azure}
                </option>
                <option value={CloudPlatform.Gcp} key={CloudPlatform.Gcp}>
                  {CloudPlatform.Gcp}
                </option>
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

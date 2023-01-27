import React, { useEffect, useState } from "react";

import { useMutation, gql, useQuery } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import { OrgAccount, OrgDimension, OrgUnit } from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";

const NEW_ORG_UNIT_MEMBERSHIP_MUTATION = gql`
  mutation addAccountToOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {
    addAccountToOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)
  }
`;

const ORG_UNITS_QUERY = gql`
  query orgDimensionsAndUnits {
    orgDimensions {
      id
      name
      orgUnits {
        id
        name
      }
    }
  }
`;

const ORG_ACCOUNTS_QUERY = gql`
  query orgAccountsAndMemberships {
    orgAccounts {
      id
      name
      cloudPlatform
      cloudIdentifier
      orgUnits {
        id
        orgDimensionID
      }
    }
  }
`;

type OrgUnitsResponse = {
  orgDimensions: OrgDimension[];
};

type OrgAccountsResponse = {
  orgAccounts: OrgAccount[];
};

interface NewOrgUnitMembershipButtonProps {
  orgDimension: OrgDimension | undefined;
  orgUnit: OrgUnit | undefined;
  orgAccount: OrgAccount | undefined;
  onCompleted: () => void;
}

export const NewOrgUnitMembershipButton: React.VFC<
  NewOrgUnitMembershipButtonProps
> = (props: NewOrgUnitMembershipButtonProps) => {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <>
      <Button variant="primary" onClick={handleShow}>
        <BsPlusCircle /> Create New Org Unit Membership
      </Button>
      <Modal show={show} onHide={handleClose}>
        {props.orgDimension !== undefined && props.orgUnit !== undefined && (
          <NewOrgUnitMembershipFromOrgUnitForm
            {...props}
            handleClose={handleClose}
            onCompleted={props.onCompleted}
            key={props.orgDimension.id + ":" + props.orgUnit.id}
          />
        )}
      </Modal>
    </>
  );
};

interface NewOrgUnitMembershipFormFromOrgUnitProps {
  orgDimension: OrgDimension | undefined;
  orgUnit: OrgUnit | undefined;
  handleClose: () => void;
  onCompleted: () => void;
}

export const NewOrgUnitMembershipFromOrgUnitForm: React.VFC<
  NewOrgUnitMembershipFormFromOrgUnitProps
> = (props: NewOrgUnitMembershipFormFromOrgUnitProps) => {
  const [selectedOrgAccountID, setSelectedOrgAccountID] = useState<string>("");

  const [mutation] = useMutation(NEW_ORG_UNIT_MEMBERSHIP_MUTATION, {
    variables: {
      orgUnitID: props.orgUnit?.id,
      orgAccountID: selectedOrgAccountID,
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
        `Created membership ${props.orgUnit?.id} / ${selectedOrgAccountID}`,
        "",
        5000
      );
    },
  });

  const { loading, error, data, refetch } = useQuery<OrgAccountsResponse>(
    ORG_ACCOUNTS_QUERY,
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

    mutation();
    setTimeout(props.onCompleted, 1000);

    props.handleClose();
  };

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const target = event.target;
    const value = target.value;
    const name = target.name;

    if (name === "orgAccountID") {
      setSelectedOrgAccountID(value);
    }
  };

  // filter out org accounts that already have a membership for this org dimension
  let filteredOrgAccounts = data?.orgAccounts.filter((orgAccount) => {
    return !orgAccount?.orgUnits?.some((orgUnit) => {
      return orgUnit?.orgDimensionID === props.orgDimension?.id;
    });
  });

  return (
    <>
      <Form onSubmit={handleSubmit}>
        <Modal.Header closeButton>
          <Modal.Title>New Org Unit Membership</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Org Dimension</Form.Label>
            <Form.Select>
              <option selected={true} disabled={true}>
                {props.orgDimension?.name}
              </option>
            </Form.Select>
          </Form.Group>
          <Form.Group>
            <Form.Label>Org Unit</Form.Label>
            <Form.Select>
              <option selected={true} disabled={true}>
                {props.orgUnit?.name}
              </option>
            </Form.Select>
          </Form.Group>
          <Form.Group>
            <Form.Label>Org Account</Form.Label>
            <Form.Select name="orgAccountID" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select an Org Account
              </option>
              {filteredOrgAccounts?.map((orgAccount) => {
                return (
                  <option value={orgAccount?.id} key={orgAccount?.id}>
                    {orgAccount?.cloudPlatform} {orgAccount?.name}
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

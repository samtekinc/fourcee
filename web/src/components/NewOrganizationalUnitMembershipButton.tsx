import React, { useEffect, useState } from "react";

import { useMutation, gql, useQuery } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button, Form, Modal } from "react-bootstrap";
import {
  Maybe,
  NewOrganizationalUnit,
  NewOrganizationalUnitMembership,
  OrganizationalAccount,
  OrganizationalAccounts,
  OrganizationalDimension,
  OrganizationalDimensions,
  OrganizationalUnit,
} from "../__generated__/graphql";
import { BsPlusCircle } from "react-icons/bs";
import { renderCloudPlatform } from "../utils/rendering";

const NEW_ORG_UNIT_MEMBERSHIP_MUTATION = gql`
  mutation createOrganizationalUnitMembership(
    $orgUnitMembership: NewOrganizationalUnitMembership!
  ) {
    createOrganizationalUnitMembership(orgUnitMembership: $orgUnitMembership) {
      orgUnitId
    }
  }
`;

const ORG_UNITS_QUERY = gql`
  query organizationalDimensionsAndUnits {
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
  }
`;

const ORG_ACCOUNTS_QUERY = gql`
  query organizationalAccountsAndMemberships {
    organizationalAccounts {
      items {
        orgAccountId
        name
        cloudPlatform
        cloudIdentifier
        orgUnitMemberships {
          items {
            orgDimensionId
            orgUnitId
          }
        }
      }
    }
  }
`;

type CreateOrgUnitResponse = {
  createOrganizationalUnitMembership: {
    orgUnitId: string;
    orgDimensionId: string;
    orgAccountId: string;
  };
};

type OrgUnitsResponse = {
  organizationalDimensions: OrganizationalDimensions;
};

type OrgAccountsResponse = {
  organizationalAccounts: OrganizationalAccounts;
};

interface NewOrganizationalUnitMembershipButtonProps {
  orgDimension: OrganizationalDimension | undefined;
  orgUnit: OrganizationalUnit | undefined;
  orgAccount: OrganizationalAccount | undefined;
  onCompleted: () => void;
}

export const NewOrganizationalUnitMembershipButton: React.VFC<
  NewOrganizationalUnitMembershipButtonProps
> = (props: NewOrganizationalUnitMembershipButtonProps) => {
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
          <NewOrganizationalUnitMembershipFromOrgUnitForm
            {...props}
            handleClose={handleClose}
            onCompleted={props.onCompleted}
            key={props.orgDimension.orgDimensionId + props.orgUnit.orgUnitId}
          />
        )}
      </Modal>
    </>
  );
};

interface NewOrganizationalUnitMembershipFormFromOrgUnitProps {
  orgDimension: OrganizationalDimension | undefined;
  orgUnit: OrganizationalUnit | undefined;
  handleClose: () => void;
  onCompleted: () => void;
}

export const NewOrganizationalUnitMembershipFromOrgUnitForm: React.VFC<
  NewOrganizationalUnitMembershipFormFromOrgUnitProps
> = (props: NewOrganizationalUnitMembershipFormFromOrgUnitProps) => {
  const [formState, setFormState] = useState<NewOrganizationalUnitMembership>({
    orgDimensionId: props.orgDimension?.orgDimensionId,
    orgUnitId: props.orgUnit?.orgUnitId,
  } as NewOrganizationalUnitMembership);

  const [mutation] = useMutation<CreateOrgUnitResponse>(
    NEW_ORG_UNIT_MEMBERSHIP_MUTATION,
    {
      variables: {
        orgUnitMembership: formState,
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
          `Created membership ${data.createOrganizationalUnitMembership.orgUnitId} / ${data.createOrganizationalUnitMembership.orgAccountId}`,
          "",
          5000
        );
      },
    }
  );

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

    setFormState((prevState) => {
      return {
        ...prevState,
        [name]: value,
      } as NewOrganizationalUnitMembership;
    });
  };

  // filter out org accounts that already have a membership for this org dimension
  let filteredOrgAccounts = data?.organizationalAccounts.items.filter(
    (orgAccount) => {
      return !orgAccount?.orgUnitMemberships?.items?.some((membership) => {
        return (
          membership?.orgDimensionId === props.orgDimension?.orgDimensionId
        );
      });
    }
  );

  return (
    <>
      <Form onSubmit={handleSubmit}>
        <Modal.Header closeButton>
          <Modal.Title>New Organizational Unit Membership</Modal.Title>
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
            <Form.Select name="orgAccountId" onChange={handleSelectChange}>
              <option selected={true} disabled={true}>
                Select an Org Account
              </option>
              {filteredOrgAccounts?.map((orgAccount) => {
                return (
                  <option
                    value={orgAccount?.orgAccountId}
                    key={orgAccount?.orgAccountId}
                  >
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

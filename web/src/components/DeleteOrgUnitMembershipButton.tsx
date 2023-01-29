import React from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const DELETE_ORG_UNIT_MEMBERSHIP_MUTATION = gql`
  mutation removeAccountFromOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {
    removeAccountFromOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)
  }
`;

type DeleteOrgUnitMembershipResposne = {
  removeAccountFromOrgUnit: boolean;
};

interface DeleteOrgUnitMembershipButtonProps {
  orgUnitID: string | undefined;
  orgAccountID: string | undefined;
  onCompleted: () => void;
}

export const DeleteOrgUnitMembershipButton: React.VFC<
  DeleteOrgUnitMembershipButtonProps
> = (props: DeleteOrgUnitMembershipButtonProps) => {
  const [mutation, { loading }] = useMutation<DeleteOrgUnitMembershipResposne>(
    DELETE_ORG_UNIT_MEMBERSHIP_MUTATION,
    {
      variables: {
        orgUnitID: props.orgUnitID,
        orgAccountID: props.orgAccountID,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error delete org unit membership`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(`Removed org unit membership`, "", 5000);
      },
    }
  );

  return (
    <>
      <Button
        disabled={loading}
        onClick={() => {
          mutation();
          setTimeout(props.onCompleted, 1000);
        }}
        variant={"danger"}
      >
        {loading ? "Submitting..." : "Remove"}
      </Button>
    </>
  );
};

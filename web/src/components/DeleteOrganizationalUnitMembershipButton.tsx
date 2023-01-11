import React from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const DELETE_ORG_UNIT_MEMBERSHIP_MUTATION = gql`
  mutation deleteOrganizationalUnitMembership(
    $orgDimensionId: ID!
    $orgAccountId: ID!
  ) {
    deleteOrganizationalUnitMembership(
      orgDimensionId: $orgDimensionId
      orgAccountId: $orgAccountId
    )
  }
`;

type DeleteOrganizationalUnitMembershipResposne = {
  deleteOrganizationalUnitMembership: boolean;
};

interface DeleteOrganizationalUnitMembershipButtonProps {
  orgDimensionId: string | undefined;
  orgAccountId: string | undefined;
}

export const DeleteOrganizationalUnitMembershipButton: React.VFC<
  DeleteOrganizationalUnitMembershipButtonProps
> = (props: DeleteOrganizationalUnitMembershipButtonProps) => {
  const [mutation, { loading }] =
    useMutation<DeleteOrganizationalUnitMembershipResposne>(
      DELETE_ORG_UNIT_MEMBERSHIP_MUTATION,
      {
        variables: {
          orgDimensionId: props.orgDimensionId,
          orgAccountId: props.orgAccountId,
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
          NotificationManager.success(`Deleted org unit membership`, "", 5000);
        },
      }
    );

  return (
    <>
      <Button
        disabled={loading}
        onClick={() => {
          mutation();
        }}
        variant={"danger"}
      >
        {loading ? "Submitting..." : "Delete"}
      </Button>
    </>
  );
};

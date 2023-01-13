import React from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const TERRAFORM_DRIFT_CHECK_MUTATION = gql`
  mutation createTerraformDriftCheckRequest($moduleAssignmentId: ID!) {
    createTerraformDriftCheckRequest(
      terraformDriftCheckRequest: { moduleAssignmentId: $moduleAssignmentId }
    ) {
      terraformDriftCheckRequestId
    }
  }
`;

interface TriggerTerraformDriftCheckButtonProps {
  moduleAssignmentId: string;
}

type TriggerTerraformDriftCheckResponse = {
  createTerraformDriftCheckRequest: {
    terraformDriftCheckRequestId: string;
  };
};

export const TriggerTerraformDriftCheckButton: React.VFC<
  TriggerTerraformDriftCheckButtonProps
> = (props: TriggerTerraformDriftCheckButtonProps) => {
  const [mutation, { loading }] =
    useMutation<TriggerTerraformDriftCheckResponse>(
      TERRAFORM_DRIFT_CHECK_MUTATION,
      {
        variables: {
          moduleAssignmentId: props.moduleAssignmentId,
        },
        onError: (error) => {
          console.log(error);
          NotificationManager.error(
            error.message,
            `Error triggering terraform drift check`,
            5000
          );
        },
        onCompleted: (data) => {
          NotificationManager.success(
            `Initiated ${data.createTerraformDriftCheckRequest.terraformDriftCheckRequestId}`,
            "",
            5000
          );
        },
      }
    );

  return (
    <Button
      disabled={loading}
      onClick={() => {
        mutation();
      }}
      variant={"primary"}
    >
      {loading ? "Submitting..." : "Drift Check Module Assignment"}
    </Button>
  );
};

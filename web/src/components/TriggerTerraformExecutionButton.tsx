import React from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const TERRAFORM_EXECUTION_MUTATION = gql`
  mutation createTerraformExecutionRequest(
    $moduleAssignmentID: ID!
    $destroy: Boolean!
  ) {
    createTerraformExecutionRequest(
      terraformExecutionRequest: {
        moduleAssignmentID: $moduleAssignmentID
        destroy: $destroy
      }
    ) {
      id
    }
  }
`;

interface TriggerTerraformExecutionButtonProps {
  moduleAssignmentID: string;
  destroy: boolean;
}

type TriggerTerraformExecutionResponse = {
  createTerraformExecutionRequest: {
    id: string;
  };
};

export const TriggerTerraformExecutionButton: React.VFC<
  TriggerTerraformExecutionButtonProps
> = (props: TriggerTerraformExecutionButtonProps) => {
  const [mutation, { loading }] =
    useMutation<TriggerTerraformExecutionResponse>(
      TERRAFORM_EXECUTION_MUTATION,
      {
        variables: {
          moduleAssignmentID: props.moduleAssignmentID,
          destroy: props.destroy,
        },
        onError: (error) => {
          console.log(error);
          NotificationManager.error(
            error.message,
            `Error triggering terraform execution`,
            5000
          );
        },
        onCompleted: (data) => {
          NotificationManager.success(
            `Initiated ${data.createTerraformExecutionRequest.id}`,
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
      variant={props.destroy ? "danger" : "primary"}
    >
      {loading
        ? "Submitting..."
        : props.destroy
        ? "Destroy Module Assignment"
        : "Execute Module Assignment"}
    </Button>
  );
};

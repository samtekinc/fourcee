import React from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const TERRAFORM_EXECUTION_MUTATION = gql`
  mutation createTerraformExecutionRequest(
    $moduleAssignmentId: ID!
    $destroy: Boolean!
  ) {
    createTerraformExecutionRequest(
      terraformExecutionRequest: {
        moduleAssignmentId: $moduleAssignmentId
        destroy: $destroy
      }
    ) {
      terraformExecutionRequestId
    }
  }
`;

interface TriggerTerraformExecutionButtonProps {
  moduleAssignmentId: string;
  destroy: boolean;
}

type TriggerTerraformExecutionResponse = {
  createTerraformExecutionRequest: {
    terraformExecutionRequestId: string;
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
          moduleAssignmentId: props.moduleAssignmentId,
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
            `Initiated ${data.createTerraformExecutionRequest.terraformExecutionRequestId}`,
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

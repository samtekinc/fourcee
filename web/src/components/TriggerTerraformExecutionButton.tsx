import React from "react";

import { useMutation, gql } from "@apollo/client";
import { NotificationManager } from "react-notifications";
import { Button } from "react-bootstrap";

const TERRAFORM_EXECUTION_MUTATION = gql`
  mutation createTerraformExecutionWorkflowRequest(
    $moduleAssignmentId: ID!
    $destroy: Boolean!
  ) {
    createTerraformExecutionWorkflowRequest(
      terraformExecutionWorkflowRequest: {
        moduleAssignmentId: $moduleAssignmentId
        destroy: $destroy
      }
    ) {
      terraformExecutionWorkflowRequestId
    }
  }
`;

interface TriggerTerraformExecutionButtonProps {
  moduleAssignmentId: string;
  destroy: boolean;
}

type TriggerTerraformExecutionResponse = {
  createTerraformExecutionWorkflowRequest: {
    terraformExecutionWorkflowRequestId: string;
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
            `Initiated ${data.createTerraformExecutionWorkflowRequest.terraformExecutionWorkflowRequestId}`,
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

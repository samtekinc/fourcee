import { gql, useMutation } from "@apollo/client";
import { Button } from "react-bootstrap";
import { NotificationManager } from "react-notifications";

const EXECUTE_MODULE_PROPAGATION_MUTATION = gql`
  mutation createModulePropagationExecutionRequest($modulePropagationId: ID!) {
    createModulePropagationExecutionRequest(
      modulePropagationExecutionRequest: {
        modulePropagationId: $modulePropagationId
      }
    ) {
      modulePropagationExecutionRequestId
      status
    }
  }
`;

interface ExecuteModulePropagationButtonProps {
  modulePropagationId: string;
}

type ExecuteModulePropagationResponse = {
  createModulePropagationExecutionRequest: {
    modulePropagationExecutionRequestId: string;
    status: string;
  };
};

export const ExecuteModulePropagationButton: React.VFC<
  ExecuteModulePropagationButtonProps
> = (props: ExecuteModulePropagationButtonProps) => {
  const [mutation, { loading }] = useMutation<ExecuteModulePropagationResponse>(
    EXECUTE_MODULE_PROPAGATION_MUTATION,
    {
      variables: {
        modulePropagationId: props.modulePropagationId,
      },
      onError: (error) => {
        console.log(error);
        NotificationManager.error(
          error.message,
          `Error executing module propagation`,
          5000
        );
      },
      onCompleted: (data) => {
        NotificationManager.success(
          `Initiated ${data.createModulePropagationExecutionRequest.modulePropagationExecutionRequestId}`,
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
    >
      {loading ? "Submitting..." : "Run Execution"}
    </Button>
  );
};

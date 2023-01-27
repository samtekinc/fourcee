import { gql, useMutation } from "@apollo/client";
import { Button } from "react-bootstrap";
import { NotificationManager } from "react-notifications";

const EXECUTE_MODULE_PROPAGATION_MUTATION = gql`
  mutation createModulePropagationExecutionRequest($modulePropagationID: ID!) {
    createModulePropagationExecutionRequest(
      modulePropagationExecutionRequest: {
        modulePropagationID: $modulePropagationID
      }
    ) {
      id
      status
    }
  }
`;

interface ExecuteModulePropagationButtonProps {
  modulePropagationID: string;
}

type ExecuteModulePropagationResponse = {
  createModulePropagationExecutionRequest: {
    id: string;
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
        modulePropagationID: props.modulePropagationID,
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
          `Initiated ${data.createModulePropagationExecutionRequest.id}`,
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

import { NotificationManager } from "react-notifications";
import { gql, useMutation } from "@apollo/client";
import { Button } from "react-bootstrap";

const DRIFT_CHECK_MODULE_PROPAGATION_MUTATION = gql`
  mutation createModulePropagationDriftCheckRequest($modulePropagationId: ID!) {
    createModulePropagationDriftCheckRequest(
      modulePropagationDriftCheckRequest: {
        modulePropagationId: $modulePropagationId
      }
    ) {
      modulePropagationDriftCheckRequestId
      status
    }
  }
`;

interface DriftCheckModulePropagationButtonProps {
  modulePropagationId: string;
}

type DriftCheckModulePropagationResponse = {
  createModulePropagationDriftCheckRequest: {
    modulePropagationDriftCheckRequestId: string;
    status: string;
  };
};

export const DriftCheckModulePropagationButton: React.VFC<
  DriftCheckModulePropagationButtonProps
> = (props: DriftCheckModulePropagationButtonProps) => {
  const [mutation, { loading }] =
    useMutation<DriftCheckModulePropagationResponse>(
      DRIFT_CHECK_MODULE_PROPAGATION_MUTATION,
      {
        variables: {
          modulePropagationId: props.modulePropagationId,
        },
        onError: (error) => {
          console.log(error);
          NotificationManager.error(
            error.message,
            `Error running drift check on module propagation`,
            5000
          );
        },
        onCompleted: (data) => {
          NotificationManager.success(
            `Initiated ${data.createModulePropagationDriftCheckRequest.modulePropagationDriftCheckRequestId}`,
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
      {loading ? "Submitting..." : "Run Drift Check"}
    </Button>
  );
};

import { NotificationManager } from "react-notifications";
import { gql, useMutation } from "@apollo/client";
import { Button } from "react-bootstrap";

const DRIFT_CHECK_MODULE_PROPAGATION_MUTATION = gql`
  mutation createModulePropagationDriftCheckRequest($modulePropagationID: ID!) {
    createModulePropagationDriftCheckRequest(
      modulePropagationDriftCheckRequest: {
        modulePropagationID: $modulePropagationID
      }
    ) {
      id
      status
    }
  }
`;

interface DriftCheckModulePropagationButtonProps {
  modulePropagationID: string;
}

type DriftCheckModulePropagationResponse = {
  createModulePropagationDriftCheckRequest: {
    id: string;
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
          modulePropagationID: props.modulePropagationID,
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
            `Initiated ${data.createModulePropagationDriftCheckRequest.id}`,
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

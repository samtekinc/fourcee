package helpers

import "fmt"

func GetApplyInitOutputKey(applyExecutionRequestId string) string {
	return fmt.Sprintf("applies/%s/init-output", applyExecutionRequestId)
}

func GetApplyOutputKey(applyExecutionRequestId string) string {
	return fmt.Sprintf("applies/%s/apply-output", applyExecutionRequestId)
}

func GetPlanInitOutputKey(planExecutionRequestId string) string {
	return fmt.Sprintf("plans/%s/init-output", planExecutionRequestId)
}

func GetPlanOutputKey(planExecutionRequestId string) string {
	return fmt.Sprintf("plans/%s/plan-output", planExecutionRequestId)
}

func GetPlanFileKey(planExecutionRequestId string) string {
	return fmt.Sprintf("plans/%s/plan.tfplan", planExecutionRequestId)
}

func GetPlanJSONKey(planExecutionRequestId string) string {
	return fmt.Sprintf("plans/%s/plan.json", planExecutionRequestId)
}

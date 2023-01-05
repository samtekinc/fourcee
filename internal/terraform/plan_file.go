package terraform

import "encoding/json"

// Based on https://developer.hashicorp.com/terraform/internals/json-format#plan-representation
type TerraformPlanFile struct {
	FormatVersion   string `json:"format_version"`
	ResourceChanges []struct {
		Address string               `json:"address"`
		Name    string               `json:"name"`
		Change  ChangeRepresentation `json:"change"`
	} `json:"resource_changes"`
	OutputChanges map[string]struct {
		Change ChangeRepresentation `json:"change"`
	} `json:"output_changes"`
}

type ChangeRepresentation struct {
	Actions []string `json:"actions"`
}

func TerraformPlanFileFromJSON(data []byte) (TerraformPlanFile, error) {
	var r TerraformPlanFile
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TerraformPlanFile) HasChanges() bool {
	hasChanges := false

	for _, resourceChange := range r.ResourceChanges {
		for _, action := range resourceChange.Change.Actions {
			if action != "no-op" {
				hasChanges = true
			}
		}
	}

	for _, outputChange := range r.OutputChanges {
		for _, action := range outputChange.Change.Actions {
			if action != "no-op" {
				hasChanges = true
			}
		}
	}

	return hasChanges
}

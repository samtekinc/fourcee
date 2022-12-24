package models

type TerraformInitOutput struct {
	Stdout []byte
	Stderr []byte
	Error  error
}

type TerraformPlanOutput struct {
	Stdout   []byte
	Stderr   []byte
	PlanFile []byte
	PlanJSON []byte
	Error    error
}

type TerraformApplyOutput struct {
	Stdout []byte
	Stderr []byte
	Error  error
}

type TerraformCommandOutput struct {
	Stdout []byte
	Stderr []byte
	Error  error
}

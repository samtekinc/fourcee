package models

type TerraformInitOutput struct {
	Stdout []byte
	Stderr []byte
	Error  error `json:"-"`
}

type TerraformPlanOutput struct {
	Stdout   []byte
	Stderr   []byte
	PlanFile []byte
	PlanJSON []byte
	Error    error `json:"-"`
}

type TerraformApplyOutput struct {
	Stdout []byte
	Stderr []byte
	Error  error `json:"-"`
}

type TerraformCommandOutput struct {
	Stdout []byte
	Stderr []byte
	Error  error `json:"-"`
}

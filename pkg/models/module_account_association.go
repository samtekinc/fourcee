package models

type ModuleAccountAssociationStatus string

const (
	ModuleAccountAssociationStatusActive   ModuleAccountAssociationStatus = "ACTIVE"
	ModuleAccountAssociationStatusInactive ModuleAccountAssociationStatus = "INACTIVE"
)

type ModuleAccountAssociation struct {
	ModulePropagationId string                         `json:"modulePropagationId"`
	OrgAccountId        string                         `json:"orgAccountId"`
	RemoteStateBucket   string                         `json:"remoteStateBucket"`
	RemoteStateKey      string                         `json:"remoteStateKey"`
	Status              ModuleAccountAssociationStatus `json:"status"`
}

type ModuleAccountAssociations struct {
	Items      []ModuleAccountAssociation `json:"items"`
	NextCursor string                     `json:"nextCursor"`
}

type NewModuleAccountAssociation struct {
	ModulePropagationId string `json:"modulePropagationId"`
	OrgAccountId        string `json:"orgAccountId"`
	RemoteStateBucket   string `json:"remoteStateBucket"`
	RemoteStateKey      string `json:"remoteStateKey"`
}

type ModuleAccountAssociationUpdate struct {
	RemoteStateBucket *string                         `json:"remoteStateBucket"`
	RemoteStateKey    *string                         `json:"remoteStateKey"`
	Status            *ModuleAccountAssociationStatus `json:"status"`
}

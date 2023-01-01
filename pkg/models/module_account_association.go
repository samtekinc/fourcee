package models

import (
	"fmt"
	"strings"
)

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
	RemoteStateRegion   string                         `json:"remoteStateRegion"`
	Status              ModuleAccountAssociationStatus `json:"status"`
}

func (m ModuleAccountAssociation) Key() ModuleAccountAssociationKey {
	return ModuleAccountAssociationKey{
		ModulePropagationId: m.ModulePropagationId,
		OrgAccountId:        m.OrgAccountId,
	}
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
	RemoteStateRegion   string `json:"remoteStateRegion"`
}

type ModuleAccountAssociationUpdate struct {
	RemoteStateBucket *string                         `json:"remoteStateBucket"`
	RemoteStateKey    *string                         `json:"remoteStateKey"`
	Status            *ModuleAccountAssociationStatus `json:"status"`
}

type ModuleAccountAssociationKey struct {
	ModulePropagationId string `json:"modulePropagationId"`
	OrgAccountId        string `json:"orgAccountId"`
}

func (m ModuleAccountAssociationKey) String() string {
	return fmt.Sprintf("%s:%s", m.ModulePropagationId, m.OrgAccountId)
}

func ParseModuleAccountAssociationKey(key string) (*ModuleAccountAssociationKey, error) {
	split := strings.Split(key, ":")
	if len(split) != 2 {
		return nil, fmt.Errorf("invalid key format")
	}
	return &ModuleAccountAssociationKey{
		ModulePropagationId: split[0],
		OrgAccountId:        split[1],
	}, nil
}

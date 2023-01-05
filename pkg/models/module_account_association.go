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
	ModulePropagationId string
	OrgAccountId        string
	RemoteStateBucket   string
	RemoteStateKey      string
	RemoteStateRegion   string
	Status              ModuleAccountAssociationStatus
}

func (m ModuleAccountAssociation) Key() ModuleAccountAssociationKey {
	return ModuleAccountAssociationKey{
		ModulePropagationId: m.ModulePropagationId,
		OrgAccountId:        m.OrgAccountId,
	}
}

type ModuleAccountAssociations struct {
	Items      []ModuleAccountAssociation
	NextCursor string
}

type NewModuleAccountAssociation struct {
	ModulePropagationId string
	OrgAccountId        string
	RemoteStateBucket   string
	RemoteStateKey      string
	RemoteStateRegion   string
}

type ModuleAccountAssociationUpdate struct {
	RemoteStateBucket *string
	RemoteStateKey    *string
	Status            *ModuleAccountAssociationStatus
}

type ModuleAccountAssociationKey struct {
	ModulePropagationId string
	OrgAccountId        string
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

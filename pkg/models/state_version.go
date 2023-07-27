package models

import "time"

type StateVersion struct {
	VersionID    string
	LastModified time.Time
	IsCurrent    bool
	Bucket       string
	Key          string
}

type StateFile struct {
	VersionID string
	Resources []StateResource
}

type StateResource struct {
	Type       string
	Name       string
	ID         string
	Attributes map[string]interface{}
}

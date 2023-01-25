package models

import "gorm.io/gorm"

type Metadata struct {
	gorm.Model
	Name         string
	Value        string
	OrgAccountID uint // foreign key
}

type MetadataInput struct {
	Name  string
	Value string
}

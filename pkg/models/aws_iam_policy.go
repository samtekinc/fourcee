package models

import "gorm.io/gorm"

type AwsIamPolicy struct {
	gorm.Model
	Name           string
	PolicyDocument string
}

type NewAwsIamPolicy struct {
	Name           string
	PolicyDocument string
}

type AwsIamPolicyUpdate struct {
	Name           *string
	PolicyDocument *string
}

type AwsIamPolicyFilters struct {
	NameContains *string
}

package models

type RequestStatus string

const (
	RequestStatusPending   RequestStatus = "PENDING"
	RequestStatusRunning   RequestStatus = "RUNNING"
	RequestStatusSucceeded RequestStatus = "SUCCEEDED"
	RequestStatusFailed    RequestStatus = "FAILED"
)

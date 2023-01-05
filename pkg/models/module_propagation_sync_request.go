package models

import "time"

type ModulePropagationDriftCheckRequest struct {
	ModulePropagationDriftCheckRequestId string
	ModulePropagationId                  string
	RequestTime                          time.Time
	Status                               RequestStatus
	WorkflowRequestId                    string
}

type ModulePropagationDriftCheckRequests struct {
	Items      []ModulePropagationDriftCheckRequest
	NextCursor string
}

type NewModulePropagationDriftCheckRequest struct {
	ModulePropagationId string
}

type ModulePropagationDriftCheckRequestUpdate struct {
	Status *RequestStatus
}

package vault

import "github.com/google/uuid"

type VaultResponse[T any] struct {
	RequestID     uuid.UUID   `json:"request_id"`
	LeaseID       string      `json:"lease_id"`
	Renewable     bool        `json:"renewable"`
	LeaseDuration int         `json:"lease_duration"`
	Data          T           `json:"data"`
	WrapInfo      interface{} `json:"wrap_info"`
	Warnings      interface{} `json:"warnings"`
	Auth          interface{} `json:"auth"`
	MountType     string      `json:"mount_type"`
}

package model

import (
	"time"
	"github.com/google/uuid"
)

type AccessLog struct {
	ID        uuid.UUID
	CreatedAt time.Time
	IpAddress string
	Api       string
}

func NewAccessLog(ipAddress string, api string) *AccessLog {
	return &AccessLog{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		IpAddress: ipAddress,
		Api:       api,
	}
}

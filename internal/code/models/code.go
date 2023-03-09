package models

import (
	"time"

	commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"
)

type Code struct {
	ID        int64
	Type      commonv1.CodeType
	CreatedAt time.Time
	UpdatedAt time.Time
}

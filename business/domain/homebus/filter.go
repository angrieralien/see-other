package homebus

import (
	"time"

	"github.com/mydomain/see-other/business/types/hometype"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
// We are using pointer semantics because the With API mutates the value.
type QueryFilter struct {
	ID               *uuid.UUID
	UserID           *uuid.UUID
	Type             *hometype.HomeType
	StartCreatedDate *time.Time
	EndCreatedDate   *time.Time
}

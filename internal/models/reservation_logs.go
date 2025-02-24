package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type ReservationLogs struct {
	db.ModelBase
	Id            uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	ReservationId *uuid.UUID `json:"reservation_id,omitempty" column:"name:reservation_id;type:uuid;nullable"`
	UserId        *uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable"`
	Action        *string    `json:"action,omitempty" column:"name:action;type:varchar;nullable"`
	Timestamp     *time.Time `json:"timestamp,omitempty" column:"name:timestamp;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservation_logs" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}

package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	
)

type Schedules struct {
	db.ModelBase
	Id        uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	DayOfWeek string    `json:"day_of_week,omitempty" column:"name:day_of_week;type:varchar;nullable:false"`
	StartTime string `json:"start_time,omitempty" column:"name:start_time;type:text;nullable:false"`
	EndTime   string `json:"end_time,omitempty" column:"name:end_time;type:text;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"schedules" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}

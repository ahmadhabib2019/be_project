package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctors struct {
	db.ModelBase
	Id             uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name           string     `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	Specialization string     `json:"specialization,omitempty" column:"name:specialization;type:text;nullable:false"`
	Phone          string     `json:"phone,omitempty" column:"name:phone;type:text;nullable:false"`
	UserId         *uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable;default:auth.uid()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	// User *Users `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}

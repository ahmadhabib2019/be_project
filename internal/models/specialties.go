package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Specialties struct {
	db.ModelBase
	Id          uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name        string    `json:"name,omitempty" column:"name:name;type:varchar;nullable:false;unique"`
	Description *string   `json:"description,omitempty" column:"name:description;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"specialties" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}

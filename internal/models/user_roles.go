package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type UserRoles struct {
	db.ModelBase
	Id     uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId *uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable"`
	RoleId *string    `json:"role_id,omitempty" column:"name:role_id;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"user_roles" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Role *Roles `json:"role,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:role_id"`
	// User *Users `json:"user,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}

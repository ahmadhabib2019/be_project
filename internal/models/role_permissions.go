package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type RolePermissions struct {
	db.ModelBase
	Id           uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	RoleId       *string    `json:"role_id,omitempty" column:"name:role_id;type:text;nullable"`
	PermissionId *uuid.UUID `json:"permission_id,omitempty" column:"name:permission_id;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"role_permissions" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Permission *Permissions `json:"permission,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:permission_id"`
	Role       *Roles       `json:"role,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:role_id"`
}

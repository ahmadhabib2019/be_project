package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Permissions struct {
	db.ModelBase
	Id   uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name string    `json:"name,omitempty" column:"name:name;type:text;nullable:false;unique"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"permissions" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	RolePermissionPermissions             []*RolePermissions `json:"role_permission_permissions,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:permission_id"`
	RolesThroughRolePermissionsPermission []*Roles           `json:"roles_through_role_permissions_permission,omitempty" join:"joinType:manyToMany;through:role_permissions;sourcePrimaryKey:id;sourceForeignKey:permission_id;targetPrimaryKey:id;targetForeign:permission_id"`
}

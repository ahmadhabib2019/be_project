package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Roles struct {
	db.ModelBase
	Id   string `json:"id,omitempty" column:"name:id;type:text;primaryKey;nullable:false"`
	Name string `json:"name,omitempty" column:"name:name;type:text;nullable:false;unique"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"roles" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	PermissionsThroughRolePermissionsRole []*Permissions     `json:"permissions_through_role_permissions_role,omitempty" join:"joinType:manyToMany;through:role_permissions;sourcePrimaryKey:id;sourceForeignKey:role_id;targetPrimaryKey:id;targetForeign:role_id"`
	RolePermissionRoles                   []*RolePermissions `json:"role_permission_roles,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:role_id"`
	UserRoleRoles                         []*UserRoles       `json:"user_role_roles,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:role_id"`
	// UsersThroughUserRolesRole             []*Users           `json:"users_through_user_roles_role,omitempty" join:"joinType:manyToMany;through:user_roles;sourcePrimaryKey:id;sourceForeignKey:role_id;targetPrimaryKey:id;targetForeign:role_id"`
}

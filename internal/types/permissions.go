package types

import (
	"github.com/sev-2/raiden"
)

type Permissions struct {
	raiden.TypeBase
}

func (t *Permissions) Name() string {
	return "_permissions"
}

func (r *Permissions) Format() string {
	return "permissions[]"
}

func (r *Permissions) Enums() []string {
	return []string{}
}

func (r *Permissions) Attributes() []string {
	return []string{}
}

func (r *Permissions) Comment() *string {
	return nil
}


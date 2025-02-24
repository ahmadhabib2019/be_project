package types

import (
	"github.com/sev-2/raiden"
)

type Specialties struct {
	raiden.TypeBase
}

func (t *Specialties) Name() string {
	return "_specialties"
}

func (r *Specialties) Format() string {
	return "specialties[]"
}

func (r *Specialties) Enums() []string {
	return []string{}
}

func (r *Specialties) Attributes() []string {
	return []string{}
}

func (r *Specialties) Comment() *string {
	return nil
}


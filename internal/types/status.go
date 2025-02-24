package types

import (
	"github.com/sev-2/raiden"
)

type Status struct {
	raiden.TypeBase
}

func (t *Status) Name() string {
	return "status"
}

func (r *Status) Format() string {
	return "status"
}

func (r *Status) Enums() []string {
	return []string{"pending","confirmed","canceled"}
}

func (r *Status) Attributes() []string {
	return []string{}
}

func (r *Status) Comment() *string {
	return nil
}


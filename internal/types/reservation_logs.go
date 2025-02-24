package types

import (
	"github.com/sev-2/raiden"
)

type ReservationLogs struct {
	raiden.TypeBase
}

func (t *ReservationLogs) Name() string {
	return "_reservation_logs"
}

func (r *ReservationLogs) Format() string {
	return "reservation_logs[]"
}

func (r *ReservationLogs) Enums() []string {
	return []string{}
}

func (r *ReservationLogs) Attributes() []string {
	return []string{}
}

func (r *ReservationLogs) Comment() *string {
	return nil
}


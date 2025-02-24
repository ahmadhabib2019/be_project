package controllers

import (
	"backendproject/internal/models"
	"github.com/sev-2/raiden"
)

type ScheduleController struct {
	raiden.ControllerBase
	Http  string `path:"/schedules" type:"rest"`
	Model models.Schedules
}
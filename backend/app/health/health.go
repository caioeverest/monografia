package health

import (
	"github.com/caioever/monografia/backend/adapters/database"
	"github.com/caioever/monografia/backend/adapters/mqtt"
)

type Status struct {
	App		bool 	`json:"app"`
	Db		bool	`json:"db"`
	Broker	bool	`json:"broker"`
}

func HealthCheck(context *mqtt.Context) *Status {
	healthStatus := &Status{}
	healthStatus.App	=	true
	healthStatus.Db		=	true
	healthStatus.Broker	=	true
	if database.ConnectionTest() != nil {healthStatus.Db = false}
	if context.ConnectionTest() {healthStatus.Broker = false}
	return healthStatus
}
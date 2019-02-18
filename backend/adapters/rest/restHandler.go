package rest

import (
	"github.com/caioever/monografia/backend/adapters/mqtt"
)

type restBase struct {
	*mqtt.Context
}

func InjectConfig (c *mqtt.Context) *restBase{
	return &restBase{c}
}
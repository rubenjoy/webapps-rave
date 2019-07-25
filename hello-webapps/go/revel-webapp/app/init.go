package app

import (
	"github.com/revel/revel"
)

var (
	AppVersion string
	BuildTime string
)

func init() {

	revel.Filters = []revel.Filter{
		revel.PanicFilter,
		revel.RouterFilter,
		revel.ActionInvoker,
	}

	revel.OnAppStart(startup)
}

func startup() {
	// empty function
}
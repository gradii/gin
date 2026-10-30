package gin

import (
	"github.com/gradii/framework/contracts/binding"
	"github.com/gradii/framework/contracts/config"
	"github.com/gradii/framework/contracts/foundation"
	"github.com/gradii/framework/contracts/log"
	"github.com/gradii/framework/contracts/validation"
	"github.com/gradii/framework/contracts/view"
)

const BindingRoute = "goravel.gin.route"

var (
	App              foundation.Application
	ConfigFacade     config.Config
	LogFacade        log.Log
	ValidationFacade validation.Validation
	ViewFacade       view.View
)

type ServiceProvider struct{}

func (r *ServiceProvider) Relationship() binding.Relationship {
	return binding.Relationship{
		Bindings: []string{
			BindingRoute,
		},
		Dependencies: []string{
			binding.Config,
			binding.Log,
			binding.Validation,
			binding.View,
		},
		ProvideFor: []string{
			binding.Route,
		},
	}
}

func (r *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.BindWith(BindingRoute, func(app foundation.Application, parameters map[string]any) (any, error) {
		return NewRoute(app.MakeConfig(), parameters)
	})
}

func (r *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	LogFacade = app.MakeLog()
	ValidationFacade = app.MakeValidation()
	ViewFacade = app.MakeView()
}

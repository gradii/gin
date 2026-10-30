package facades

import (
	"log"

	"github.com/gradii/gin"

	"github.com/gradii/framework/contracts/route"
)

func Route(driver string) route.Route {
	instance, err := gin.App.MakeWith(gin.BindingRoute, map[string]any{
		"driver": driver,
	})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(*gin.Route)
}

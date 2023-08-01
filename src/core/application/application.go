package application

import (
	"agora-vnf-manager/db"

	"github.com/labstack/echo/v4"
)

type Application struct {
	Router *Router
	Echo   *echo.Echo
	Db     db.IDbConnection
}

type InitializerHandlerFunc func(*Application) error

var MyApp *Application = nil
var AppInitializers []InitializerHandlerFunc = nil

func RegisterInitializer(handler InitializerHandlerFunc) {
	if AppInitializers == nil {
		AppInitializers = make([]InitializerHandlerFunc, 0, 50)
	}
	if MyApp == nil {
		AppInitializers = append(AppInitializers, handler)
	} else {
		handler(MyApp)
	}
}

package server

import (
	config "agora-vnf-manager/config"
	app "agora-vnf-manager/core/application"
	log "agora-vnf-manager/core/log"
	db "agora-vnf-manager/db"
	consul "agora-vnf-manager/features/consul"
	helm "agora-vnf-manager/features/helm"
	kubernetes "agora-vnf-manager/features/kubernetes"
	vnf_device_mapper "agora-vnf-manager/features/vnf-device-mapper"
	vnf_infrastructure "agora-vnf-manager/features/vnf-infrastructure"
	vnf_instance "agora-vnf-manager/features/vnf-instance"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func get_database_connection(cfg *config.Config, application *app.Application) {
	db := db.CreateDb(cfg)
	engine := db.GetEngine()
	application.Db = engine
	log.Info("[get_database_connection]: Iniciated connection with the PostgreSQL database")
}

func configure_web_router(application *app.Application) {
	application.Echo = echo.New()
	application.Router = app.CreateRouter(application)
}

func configure_web_routes(application *app.Application) {
	// Swagger configuration
	application.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	// Proper AGORA VNF Manager endpoint binding
	helm.BindRoutes(application.Router)
	consul.BindRoutes(application.Router)
	kubernetes.BindRoutes(application.Router)
	vnf_infrastructure.BindRoutes(application.Router)
	vnf_instance.BindRoutes(application.Router)
	vnf_device_mapper.BindRoutes(application.Router)

}

func init_services() (err error) {
	err = consul.InitConsulService()
	return err
}

func init_entity_services(application *app.Application) {
	vnf_infrastructure.InitVnfInfrastructure(application)
	vnf_instance.InitVnfInstanceService(application)
	vnf_device_mapper.InitVnfDeviceMapper(application)
}

func run_app_initializers(application *app.Application, initializers []app.InitializerHandlerFunc) {
	for _, h := range initializers {
		h(application)
	}
}

func start_web_server(cfg *config.Config, application *app.Application) {
	application.Echo.Logger.Info(application.Echo.Start(":" + cfg.Server.Port))
}

func configure_web_server() {
	cfg := config.GetConfig()
	app.MyApp = new(app.Application)
	get_database_connection(cfg, app.MyApp)
	configure_web_router(app.MyApp)
	configure_web_routes(app.MyApp)
	err := init_services()
	if err != nil {
		os.Exit(-1)
	}
	init_entity_services(app.MyApp)
	run_app_initializers(app.MyApp, app.AppInitializers)
	start_web_server(cfg, app.MyApp)
}

func Init() {
	configure_web_server()
}

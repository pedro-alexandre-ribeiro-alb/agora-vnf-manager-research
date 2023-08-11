package consul

import (
	app "agora-vnf-manager/core/application"
	log "agora-vnf-manager/core/log"
	"net/http"
)

func BindRoutes(router *app.Router) {
	router.GET("/agora/vnfm/rest/v1/consul/service", handleGetConsulDiscoveredServices)
	router.GET("/agora/vnfm/rest/v1/consul/node", handleGetConsulDiscoveredNodes)
}

// handleGetConsulDiscoveredServices godoc
// @Summary			Lists the consul discovered services
// @Description		Lists the consul discovered services
// @Tags			consul
// @Router			/agora/vnfm/rest/v1/consul/service
func handleGetConsulDiscoveredServices(rc app.RouterContext) error {
	services, err := DiscoverServices()
	if err != nil {
		log.Errorf("[ConsulRouter - handleGetConsulDiscoveredServices]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), err.Error(), "")
	}
	return rc.Response().JSON(http.StatusOK, services)
}

// handleGetConsulDiscoveredNode godoc
// @Summary			List the consul discovered nodes
// @Description		List the consul discovered nodes
// @Tags			consul
// @Router 			/agora/vnfm/rest/v1/consul/node
func handleGetConsulDiscoveredNodes(rc app.RouterContext) error {
	nodes, err := DiscoverNodes()
	if err != nil {
		log.Errorf("[ConsulRouter - handleGetConsulDiscoveredServices]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), err.Error(), "")
	}
	return rc.Response().JSON(http.StatusOK, nodes)
}

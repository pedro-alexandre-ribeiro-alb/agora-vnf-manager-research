package kubernetes

import (
	app "agora-vnf-manager/core/application"
	errorMessages "agora-vnf-manager/core/error"
	log "agora-vnf-manager/core/log"
	types "agora-vnf-manager/features/kubernetes/types"
	kubernetesValidation "agora-vnf-manager/features/kubernetes/types/validation"
	"net/http"
)

func BindRoutes(router *app.Router) {
	router.GET("/agora/vnfm/rest/v1/pods", handleGetPods)
}

//	handleGetPods godoc
//
//	@Summary 		Get pods
//	@Description	Get all pods that are discoverable
//	@Tags			kubernetes
//	@Produce 		json
//	@Router			/agora/vnfm/rest/v1/pods [get]
func handleGetPods(rc app.RouterContext) error {
	specification := types.Specifications{}
	if err := rc.Request().Body(&specification); err != nil {
		log.Errorf("[KubernetesRouter - handeGetPods]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := kubernetesValidation.ValidateSpecification(specification)
	if !valid {
		log.Infof("[KubernetesRouter - handleGetPods]: Could not validate specification - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	pods, err := ListContainers(specification)
	if err != nil {
		log.Errorf("[KubernetesRouter - handleGetPods]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), err.Error(), "")
	}
	log.Infof("[KubernetesRouter - handleGetPods]: Retrieved pods - %+v", pods)
	return rc.Response().JSON(http.StatusOK, pods)
}

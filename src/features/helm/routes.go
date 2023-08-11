package helm

import (
	app "agora-vnf-manager/core/application"
	errorMessages "agora-vnf-manager/core/error"
	log "agora-vnf-manager/core/log"
	types "agora-vnf-manager/features/helm/types"
	helmValidation "agora-vnf-manager/features/helm/types/validation"
	"fmt"
	"net/http"
)

func BindRoutes(router *app.Router) {
	router.POST("/agora/vnfm/rest/v1/helm", handleCreateDeployment)
	router.DELETE("/agora/vnfm/rest/v1/helm", handleDeleteDeployment)
	router.GET("/agora/vnfm/rest/v1/helm/repository", handleGetRegisteredRepositories)
	router.GET("/agora/vnfm/rest/v1/helm/repository/chart/:repository", handleGetRegisteredRepositoryCharts)
}

// handleCreateDeployment godoc
// @Summary		Create helm deployment
// @Description	Creates the helm deployment specified by the provided specification
// @Tags			helm
// @Router			/agora/vnfm/rest/v1/helm [POST]
func handleCreateDeployment(rc app.RouterContext) error {
	specification := types.Specification{}
	if err := rc.Request().Body(&specification); err != nil {
		log.Errorf("[HelmRouter - handleCreateDeployment]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := helmValidation.ValidateCreateDeploymentSpecification(specification)
	if !valid {
		log.Infof("[HelmRouter - handleCreateDeployment]: Could not validate specification - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	response, err := DeployHelmChart(specification)
	if err != nil {
		log.Errorf("[HelmRouter - handleCreateDeployment]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), err.Error(), "")
	}
	log.Infof("[HelmRouter - handleCreateDeployment]: Response - %+v", response)
	return rc.Response().JSON(http.StatusOK, fmt.Sprintf("Deployment %s created", specification.ReleaseName))
}

// handleDeleteDeployment godoc
//
// @Summary		Delete helm deployment
// @Description	Deletes the helm deployment specified by the provided specification
// @Tags			helm
// @Router			/agora/vnfm/rest/v1/helm [DELETE]
func handleDeleteDeployment(rc app.RouterContext) error {
	specification := types.Specification{}
	if err := rc.Request().Body(&specification); err != nil {
		log.Errorf("[HelmService - handleDeleteDeployment]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := helmValidation.ValidateDeleteDeploymentSpecification(specification)
	if !valid {
		log.Infof("[HelmService - handleDeleteDeployment]: Could not validate specification - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	response, err := UndeployHelmChart(specification)
	if err != nil {
		log.Errorf("[HelmService - handleDeleteDeployment]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), err.Error(), "")
	}
	log.Infof("[HelmService - handleDeleteDeployment]: Response - %+v", response)
	return rc.Response().JSON(http.StatusOK, fmt.Sprintf("Deployment %s deleted", specification.ReleaseName))
}

// handleGetRegisteredRepositories godoc
//
//	@Summary		List the currently registered helm chart repositories
//	@Description	Lists the currently registered helm chart repositories
//	@Tags			helm
//	@Router			/agora/vnfm/rest/v1/helm/repository [GET]
func handleGetRegisteredRepositories(rc app.RouterContext) error {
	repositories, err := ListEnrolledRepositories()
	if err != nil {
		log.Errorf("[HelmService - handleGetRegisteredRepositories]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), err.Error(), "")
	}
	return rc.Response().JSON(http.StatusOK, repositories)
}

// handleGetRegisteredRepositoryCharts godoc
//
// @Summary		List the helm charts hosted in the referenced helm chart repository
// @Description	List the helm charts hosted in the referenced helm chart repository
// @Tags		helm
// @Router 		/agora/vnfm/rest/v1/helm/repository/chart/:repository [GET]
func handleGetRegisteredRepositoryCharts(rc app.RouterContext) error {
	repository := rc.Request().GetParam("repository")
	helm_charts, err := ListRepositoryCharts(repository)
	if err != nil {
		log.Errorf("[HelmRouter - handleGetRegisteredRepositoryCharts]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), err.Error(), "")
	}
	return rc.Response().JSON(http.StatusOK, helm_charts)
}

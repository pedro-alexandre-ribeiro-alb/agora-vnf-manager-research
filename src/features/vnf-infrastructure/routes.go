package vnf_infrastructure

import (
	app "agora-vnf-manager/core/application"
	common "agora-vnf-manager/core/common"
	errorMessages "agora-vnf-manager/core/error"
	log "agora-vnf-manager/core/log"
	types "agora-vnf-manager/features/vnf-infrastructure/types"
	validation "agora-vnf-manager/features/vnf-infrastructure/types/validation"
	"fmt"
	"net/http"
	"strconv"
)

func BindRoutes(router *app.Router) {
	router.GET("/agora/vnfm/rest/v1/vnfinfra", handleGetVnfInfrastructures)
	router.GET("/agora/vnfm/rest/v1/vnfinfra/:id", handleGetSpecifiedVnfInfrastructure)
	router.PUT("/agora/vnfm/rest/v1/vnfinfra/:id", handleUpdateVnfInfrastructure)
	router.POST("/agora/vnfm/rest/v1/vnfinfra", handleCreateVnfInfrastructure)
	router.DELETE("/agora/vnfm/rest/v1/vnfinfra/:id", handleDeleteVnfInfrastructure)
}

// handleGetVnfInfrastructures godoc
//
//	@Summary		Get vnf infrastructures
//	@Description	Get the vnf infrastructures that match the specified parameters
//	@Tags			vnf_infrastructure
//	@Produce		json
//	@Param			vnf_infrastructure		body		types.VnfInfrastructureDocs			false		"Vnf infrastructure filters"
//	@Success 		200						{object}	[]types.VnfInfrastructureDocs
//	@Failure		400						{object}	error.RouterResponseError
//	@Failure		401						{object}	error.RouterResponseError
//	@Failure		403						{object}	error.RouterResponseError
//	@Failure		404						{object}	error.RouterResponseError
//	@Failure		405						{object}	error.RouterResponseError
//	@Failure		406						{object} 	error.RouterResponseError
//	@Failure		408						{object}	error.RouterResponseError
//	@Failure		415						{object} 	error.RouterResponseError
//	@Failure		422						{object}	error.RouterResponseError
//	@Failure		500						{object}	error.RouterResponseError
//	@Failure		502						{object}	error.RouterResponseError
//	@Router			/agora/vnfm/rest/v1/vnfinfra [get]
func handleGetVnfInfrastructures(rc app.RouterContext) error {
	vnf_infrastructure := types.VnfInfrastructure{}
	if err := rc.Request().Body(&vnf_infrastructure); err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleGetVnfInfrastructures]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	vnf_infrastructures, err := GetVnfInfrastructures(vnf_infrastructure)
	if err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleGetVnfInfrastructures]: Error retrieving vnf_infrastructures - %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInfrastructureRouter - handleGetVnfInfrastructures]: Retrieved vnf_infrastructures: {vnf_infrastructures: %+v}", vnf_infrastructures)
	return rc.Response().JSON(http.StatusOK, vnf_infrastructures)
}

// handleGetSpecifiedVnfInfrastructure godoc
//
//	@Summary		Get specfied vnf infrastructure
//	@Description	Get the vnf infrastructure specified by the provided id
//	@Tags			vnf_infrastructure
//	@Produce		json
//	@Param 			id						path		string								true		"Vnf infrastructure id"
//	@Success 		200						{object}	[]types.VnfInfrastructureDocs
//	@Failure		400						{object}	error.RouterResponseError
//	@Failure		401						{object}	error.RouterResponseError
//	@Failure		403						{object}	error.RouterResponseError
//	@Failure		404						{object}	error.RouterResponseError
//	@Failure		405						{object}	error.RouterResponseError
//	@Failure		406						{object} 	error.RouterResponseError
//	@Failure		408						{object}	error.RouterResponseError
//	@Failure		415						{object} 	error.RouterResponseError
//	@Failure		422						{object}	error.RouterResponseError
//	@Failure		500						{object}	error.RouterResponseError
//	@Failure		502						{object}	error.RouterResponseError
//	@Router			/agora/vnfm/rest/v1/vnfinfra/:id [get]
func handleGetSpecifiedVnfInfrastructure(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfInfrastructureRouter - handleGetSpecifiedVnfInfrastructure]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleGetSpecifiedVnfInfrastructure]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	vnf_infrastructure, found, err := GetVnfInfrastructure(id)
	if err != nil || !found {
		if err != nil {
			log.Errorf("[VnfInfrastructureRouter - handleGetSpecifiedVnfInfrastructure]: %s", err.Error())
			return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
		}
		log.Errorf("[VnfInfrastructureRouter - handleGetSpecifiedVnfInfrastructure]: Could not find vnf_infrastructure with id: %d", id)
		return rc.Response().Error(http.StatusNotFound, http.StatusNotFound, errorMessages.ResourceNotFound(common.ResourceVnfInfrastructure, string_id), errorMessages.ResourceNotFound(common.ResourceVnfInfrastructure, string_id), "")
	}
	log.Infof("[VnfInfrastructureRouter - handleGetSpecifiedVnfInfrastructure]: Retrieved vnf_infrastructure: {vnf_infrastructure: %+v}", vnf_infrastructure)
	return rc.Response().JSON(http.StatusOK, vnf_infrastructure)
}

// handleUpdateVnfInfrastructure godoc
//
//	@Summary 		Update specified vnf infrastructure
//	@Description	Update the vnf infrastructure identified by the provided id
//	@Tags			vnf_infrastructure
//	@Produce		json
//	@Param			id					path		string							true		"Vnf infrastructure id"
//	@Param			vnf_infrastructure	body		types.VnfInfrastructureDocs		true		"Vnf infrastructure updates"
//	@Success		200					{object}	types.VnfInfrastructureDocs
//	@Failure		400					{object}	error.RouterResponseError
//	@Failure		401					{object}	error.RouterResponseError
//	@Failure		403					{object}	error.RouterResponseError
//	@Failure		404					{object}	error.RouterResponseError
//	@Failure		405					{object}	error.RouterResponseError
//	@Failure		406					{object} 	error.RouterResponseError
//	@Failure		408					{object}	error.RouterResponseError
//	@Failure		415					{object} 	error.RouterResponseError
//	@Failure		422					{object}	error.RouterResponseError
//	@Failure		500					{object}	error.RouterResponseError
//	@Failure		502					{object}	error.RouterResponseError
//	@Router			/agora/vnfm/rest/v1/vnfinfra/:id [put]
func handleUpdateVnfInfrastructure(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfInfrastructureRouter - handleUpdateVnfInfrastructure]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleUpdateVnfInfrastructure]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	vnf_infrastructure := types.VnfInfrastructure{}
	if err := rc.Request().Body(&vnf_infrastructure); err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleUpdateVnfInfrastructure]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := validation.ValidateUpdate(vnf_infrastructure, id)
	if !valid {
		log.Errorf("[VnfInfrastructureRouter - handleUpdateVnfInfrastructure]: Validation failed - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	vnf_infrastructure, err = UpdateVnfInfrastructure(vnf_infrastructure, id)
	if err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleUpdateVnfInfrastructure]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInfrastructureRouter - handleUpdateVnfInfrastructure]: Updated vnf_infrastructure {vnf_infrastructure: %+v}", vnf_infrastructure)
	return rc.Response().JSON(http.StatusOK, vnf_infrastructure)
}

// handleCreateVnfInfrastructure godoc
//
//	@Summary		Create vnf infrastructure
//	@Description	Create vnf infrastructure
//	@Param			vnf_infrastructure		body		types.VnfInfrastructureDocs		true		"Vnf infrastructure"
//	@Success		201						{object}	types.VnfInfrastructureDocs
//	@Failure		400						{object}	error.RouterResponseError
//	@Failure		401						{object}	error.RouterResponseError
//	@Failure		403						{object}	error.RouterResponseError
//	@Failure		404						{object}	error.RouterResponseError
//	@Failure		405						{object}	error.RouterResponseError
//	@Failure		406						{object} 	error.RouterResponseError
//	@Failure		408						{object}	error.RouterResponseError
//	@Failure		415						{object} 	error.RouterResponseError
//	@Failure		422						{object}	error.RouterResponseError
//	@Failure		500						{object}	error.RouterResponseError
//	@Failure		502						{object}	error.RouterResponseError
//	@Router			/agora/vnfm/rest/v1/vnfinfra [post]
func handleCreateVnfInfrastructure(rc app.RouterContext) error {
	vnf_infrastructure := types.VnfInfrastructure{}
	if err := rc.Request().Body(&vnf_infrastructure); err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleCreateVnfInfrastructure]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := validation.ValidateCreate(vnf_infrastructure)
	if !valid {
		log.Errorf("[VnfInfrastructureRouter - handleCreateVnfInfrastructure]: Validation failed - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	vnf_infrastructure, err := CreateVnfInfrastructure(vnf_infrastructure)
	if err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleCreateVnfInfrastructure]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInfrastructureRouter - handleCreateVnfInfrastructures]: Created vnf_infrastructure {vnf_infrastructure: %+v}", vnf_infrastructure)
	return rc.Response().JSON(http.StatusCreated, vnf_infrastructure)
}

// handleDeleteVnfInfrastructure godoc
//
//	@Summary		Delete specified vnf infrastructure
//	@Description	Delete the vnf infrastructure identified by the provided id
//	@Tags			vnf_infrastructure
//	@Produce		json
//	@Param			id					path		string						true		"Vnf infrastructure id"
//	@Success		200					{object}	types.VnfInstanceDocs
//	@Failure		400					{object}	error.RouterResponseError
//	@Failure		401					{object}	error.RouterResponseError
//	@Failure		403					{object}	error.RouterResponseError
//	@Failure		404					{object}	error.RouterResponseError
//	@Failure		405					{object}	error.RouterResponseError
//	@Failure		406					{object} 	error.RouterResponseError
//	@Failure		408					{object}	error.RouterResponseError
//	@Failure		415					{object} 	error.RouterResponseError
//	@Failure		422					{object}	error.RouterResponseError
//	@Failure		500					{object}	error.RouterResponseError
//	@Failure		502					{object}	error.RouterResponseError
//	@Router			/agora/vnfm/rest/v1/vnfinfra/:id [delete]
func handleDeleteVnfInfrastructure(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfInfrastructureRouter - handleDeleteVnfInfrastructure]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfInfrastructureRouter - handleDeleteVnfInfrastructure]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	success, err := DeleteVnfInfrastructure(id)
	if err != nil || !success {
		log.Errorf("[VnfInfrastructureRouter - handleDeleteVnfInfrastructure]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInfrastructureRouter - handleDeleteVnfInfrastructure] Deleted vnf_infrastructure {id: %d}", id)
	return rc.Response().JSON(http.StatusOK, fmt.Sprintf("Deleted vnf_infrastructure with id: %d", id))
}

package vnf_instance

import (
	app "agora-vnf-manager/core/application"
	common "agora-vnf-manager/core/common"
	errorMessages "agora-vnf-manager/core/error"
	log "agora-vnf-manager/core/log"
	types "agora-vnf-manager/features/vnf-instance/types"
	validation "agora-vnf-manager/features/vnf-instance/types/validation"
	"fmt"
	"net/http"
	"strconv"
)

func BindRoutes(router *app.Router) {
	router.GET("/agora/vnfm/rest/v1/vnfinstance", handleGetVnfInstances)
	router.GET("/agora/vnfm/rest/v1/vnfinstance/:id", handleGetSpecifiedVnfInstance)
	router.PUT("/agora/vnfm/rest/v1/vnfinstance/:id", handleUpdateVnfInstance)
	router.POST("/agora/vnfm/rest/v1/vnfinstance", handleCreateVnfInstance)
	router.DELETE("/agora/vnfm/rest/v1/vnfinstance/:id", handleDeleteVnfInstance)
}

// handleGetVnfInstances godoc
//
//	@Summary 		Get vnf instances
//	@Description	Get the vnf instances that match the specified parameters
//	@Tags			vnf_instance
//	@Produce		json
//	@Param			vnf_instance		body		types.VnfInstanceDocs		false		"Vnf instance filters"
//	@Success 		200					{object}	[]types.VnfInstanceDocs
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
//	@Router			/agora/vnfm/rest/v1/vnfinstance [get]
func handleGetVnfInstances(rc app.RouterContext) error {
	vnf_instance := types.VnfInstance{}
	if err := rc.Request().Body(&vnf_instance); err != nil {
		log.Errorf("[VnfInstanaceRouter - handleGetVnfInstances]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	vnf_instances, err := GetVnfInstances(vnf_instance)
	if err != nil {
		log.Errorf("[VnfInstanceRouter - handleGetVnfInstances]: Error retrieving vnf_instances - %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInstanceRouter - handleGetVnfInstances]: Retrieved vnf_instances - {vnf_instances: %+v}", vnf_instances)
	return rc.Response().JSON(http.StatusOK, vnf_instances)
}

// handleGetSpecifiedVnfInstance godoc
//
//	@Summary		Get specified vnf instance
//	@Description	Get the vnf instance identified by the provided id
//	@Tags			vnf_instance
//	@Produce		json
//	@Param			id					path		string						true		"Vnf instance id"
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
//	@Router			/agora/vnfm/rest/v1/vnfinstance/:id [get]
func handleGetSpecifiedVnfInstance(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfInstanceRouter - handleGetSpecifiedVnfInstance]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfInstanceRouter - handleGetSpecifiedVnfInstance]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	vnf_instance, found, err := GetVnfInstance(id)
	if err != nil || !found {
		if err != nil {
			log.Errorf("[VnfInstanceRouter - handleGetSpecifiedVnfInstance]: %s", err.Error())
			return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
		}
		log.Errorf("[VnfInstanceRouter - handleGetSpecifiedVnfInstance]: Could not find vnf_instance with id: %d", id)
		return rc.Response().Error(http.StatusNotFound, http.StatusNotFound, errorMessages.ResourceNotFound(common.ResourceVnfInstance, string_id), errorMessages.ResourceNotFound(common.ResourceVnfInstance, string_id), "")
	}
	log.Infof("[VnfInstanceRouter - handleGetSpecifiedVnfInstance]: Retrieved vnf_instance - {vnf_instance: %+v}", vnf_instance)
	return rc.Response().JSON(http.StatusOK, vnf_instance)
}

// handleUpdateVnfInstance godoc
//
//	@Summary		Update specified vnf instance
//	@Description	Update the vnf instance identified by the provided id
//	@Tags			vnf_instance
//	@Produce		json
//	@Param			id					path		string						true		"Vnf instance id"
//	@Param			vnf_instance		body		types.VnfInstanceDocs		true		"Vnf instance updates"
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
//	@Router			/agora/vnfm/rest/v1/vnfinstance/:id [put]
func handleUpdateVnfInstance(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfInstanceRouter - handleUpdateVnfInstance]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfInstanceRouter - handleUpdateVnfInstance]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	vnf_instance := types.VnfInstance{}
	if err := rc.Request().Body(&vnf_instance); err != nil {
		log.Errorf("[VnfInstanaceRouter - handleUpdateVnfInstance]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := validation.ValidateUpdate(vnf_instance, id)
	if !valid {
		log.Errorf("[VnfInstanceRouter - handleUpdateVnfInstance]: Validation failed - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	vnf_instance, err = UpdateVnfInstance(vnf_instance, id)
	if err != nil {
		log.Errorf("[VnfInstanceRouter - handleUpdateVnfInstance]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInstanceRouter - handleUpdateVnfInstance]: Updated vnf_instance {vnf_instance: %+v}", vnf_instance)
	return rc.Response().JSON(http.StatusOK, vnf_instance)
}

// handleCreateVnfInstance godoc
//
//	@Summary		Create vnf instance
//	@Description	Create vnf instance
//	@Param			vnf_instance		body		types.VnfInstanceDocs		true		"Vnf instance"
//	@Success		201					{object}	types.VnfInstanceDocs
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
//	@Router			/agora/vnfm/rest/v1/vnfinstance [post]
func handleCreateVnfInstance(rc app.RouterContext) error {
	vnf_instance := types.VnfInstance{}
	if err := rc.Request().Body(&vnf_instance); err != nil {
		log.Errorf("[VnfInstanaceRouter - handleCreateVnfInstances]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := validation.ValidateCreate(vnf_instance)
	if !valid {
		log.Errorf("[VnfInstanceRouter - handleCreateVnfInstance]: Validation failed - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	vnf_instance, err := CreateVnfInstance(vnf_instance)
	if err != nil {
		log.Errorf("[VnfInstanceRouter - handleCreateVnfInstances]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInstanceRouter - handleCreateVnfInstances]: Created vnf_instance {vnf_instance: %+v}", vnf_instance)
	return rc.Response().JSON(http.StatusCreated, vnf_instance)
}

// handleDeleteVnfInstance godoc
//
//	@Summary		Delete specified vnf instance
//	@Description	Delete the vnf instance identified by the provided id
//	@Tags			vnf_instance
//	@Produce		json
//	@Param			id					path		string						true		"Vnf instance id"
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
//	@Router			/agora/vnfm/rest/v1/vnfinstance/:id [delete]
func handleDeleteVnfInstance(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfInstanceRouter - handleDeleteVnfInstance]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfInstanceRouter - handleDeleteVnfInstance]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	success, err := DeleteVnfInstance(id)
	if err != nil || !success {
		log.Errorf("[VnfInstanceRouter - handleDeleteVnfInstance]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfInstanceRouter - handleDeleteVnfInstance] Deleted vnf_instance {id: %d}", id)
	return rc.Response().JSON(http.StatusOK, fmt.Sprintf("Deleted vnf_instance with id: %d", id))
}

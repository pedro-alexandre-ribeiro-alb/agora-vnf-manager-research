package vnf_device_mapper

import (
	app "agora-vnf-manager/core/application"
	common "agora-vnf-manager/core/common"
	errorMessages "agora-vnf-manager/core/error"
	log "agora-vnf-manager/core/log"
	types "agora-vnf-manager/features/vnf-device-mapper/types"
	validation "agora-vnf-manager/features/vnf-device-mapper/types/validation"
	"fmt"
	"net/http"
	"strconv"
)

func BindRoutes(router *app.Router) {
	router.GET("/agora/vnfm/rest/v1/vnfmapper", handleGetVnfDeviceMappers)
	router.GET("/agora/vnfm/rest/v1/vnfmapper/:id", handleGetSpecifiedVnfDeviceMapper)
	router.PUT("/agora/vnfm/rest/v1/vnfmapper/:id", handleUpdateVnfDeviceMapper)
	router.POST("/agora/vnfm/rest/v1/vnfmapper", handleCreateVnfDeviceMapper)
	router.DELETE("/agora/vnfm/rest/v1/vnfmapper/:id", handleDeleteVnfDeviceMapper)
}

// handleGetVnfDeviceMappers godoc
//
//	@Summary 		Get vnf device mappers
//	@Description	Get the vnf device mappers that match the specified parameters
//	@Tags			vnf_device_mappers
//	@Produce		json
//	@Param			vnf_instance		body		types.VnfDeviceMapperDocs		false		"Vnf device mapper filters"
//	@Success 		200					{object}	[]types.VnfDeviceMapperDocs
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
//	@Router			/agora/vnfm/rest/v1/vnfmapper [get]
func handleGetVnfDeviceMappers(rc app.RouterContext) error {
	vnf_device_mapper := types.VnfDeviceMapper{}
	if err := rc.Request().Body(&vnf_device_mapper); err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleGetVnfDeviceMappers]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	vnf_device_mappers, err := GetVnfDeviceMappers(vnf_device_mapper)
	if err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleGetVnfDeviceMappers]: Error retrieving vnf_device_mappers - %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfDeviceMapperRouter - handleGetVnfDeviceMappers]: Retrieved vnf_device_mappers - {vnf_device_mappers: %+v}", vnf_device_mappers)
	return rc.Response().JSON(http.StatusOK, vnf_device_mappers)
}

// handleGetSpecifiedVnfDeviceMapper godoc
//
//	@Summary		Get specified vnf device mapper
//	@Description	Get the vnf device mapper identified by the provided id
//	@Tags			vnf_device_mapper
//	@Produce		json
//	@Param			id					path		string						true		"Vnf device mapper id"
//	@Success		200					{object}	types.VnfDeviceMapperDocs
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
//	@Router			/agora/vnfm/rest/v1/vnfmapper/:id [get]
func handleGetSpecifiedVnfDeviceMapper(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfDeviceMapperRouter - handleGetSpecifiedVnfDeviceMapper]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleGetSpecifiedVnfDeviceMapper]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	vnf_device_mapper, found, err := GetVnfDeviceMapper(id)
	if err != nil || !found {
		if err != nil {
			log.Errorf("[VnfDeviceMapperRouter - handleGetSpecifiedVnfDeviceMapper]: %s", err.Error())
			return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
		}
		log.Errorf("[VnfDeviceMapperRouter - handleGetSpecifiedVnfDeviceMapper]: Could not find vnf_device_mapper with id: %d", id)
		return rc.Response().Error(http.StatusNotFound, http.StatusNotFound, errorMessages.ResourceNotFound(common.ResourceVnfDeviceMapper, string_id), errorMessages.ResourceNotFound(common.ResourceVnfDeviceMapper, string_id), "")
	}
	log.Infof("[VnfDeviceMapperRouter - handleGetSpecifiedVnfDeviceMapper]: Retrieved vnf_device_mapper - {vnf_device_mapper: %+v}", vnf_device_mapper)
	return rc.Response().JSON(http.StatusOK, vnf_device_mapper)
}

// handleUpdateVnfDeviceMapper godoc
//
//	@Summary		Update specified vnf device mapper
//	@Description	Update the vnf device mapper identified by the provided id
//	@Tags			vnf_device_mapper
//	@Produce		json
//	@Param			id					path		string							true		"Vnf device mapper id"
//	@Param			vnf_device_mapper	body		types.VnfDeviceMapperDocs		true		"Vnf device mapper updates"
//	@Success		200					{object}	types.VnfDeviceMapperDocs
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
//	@Router			/agora/vnfm/rest/v1/vnfmapper/:id [put]
func handleUpdateVnfDeviceMapper(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfDeviceMapperRouter - handleUpdateVnfDeviceMapper]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleUpdateVnfDeviceMapper]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	vnf_device_mapper := types.VnfDeviceMapper{}
	if err := rc.Request().Body(&vnf_device_mapper); err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleUpdateVnfDeviceMapper]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := validation.ValidateUpdate(vnf_device_mapper, id)
	if !valid {
		log.Errorf("[VnfDeviceMapperRouter - handleUpdateVnfDeviceMapper]: Validation failed - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	vnf_device_mapper, err = UpdateVnfDeviceMapper(vnf_device_mapper, id)
	if err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleUpdateVnfDeviceMapper]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfDeviceMapperRouter - handleUpdateVnfDeviceMapper]: Updated vnf_device_mapper {vnf_device_mapper: %+v}", vnf_device_mapper)
	return rc.Response().JSON(http.StatusOK, vnf_device_mapper)
}

// handleCreateVnfDeviceMapper godoc
//
//	@Summary		Create vnf device mapper
//	@Description	Create vnf device mapper
//	@Param			vnf_device_mapper	body		types.VnfDeviceMapperDocs		true		"Vnf device mapper"
//	@Success		201					{object}	types.VnfDeviceMapperDocs
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
//	@Router			/agora/vnfm/rest/v1/vnfmapper [post]
func handleCreateVnfDeviceMapper(rc app.RouterContext) error {
	vnf_device_mapper := types.VnfDeviceMapper{}
	if err := rc.Request().Body(&vnf_device_mapper); err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleCreateVnfDeviceMapper]: Error parsing - %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.UnexpectedInputFormat(), "")
	}
	valid, message := validation.ValidateCreate(vnf_device_mapper)
	if !valid {
		log.Errorf("[VnfDeviceMapperRouter - handleCreateVnfDeviceMapper]: Validation failed - %s", message)
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, message, message, "")
	}
	vnf_device_mapper, err := CreateVnfDeviceMapper(vnf_device_mapper)
	if err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleCreateVnfDeviceMapper]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfDeviceMapperRouter - handleCreateVnfDeviceMapper]: Created vnf_device_mapper {vnf_device_mapper: %+v}", vnf_device_mapper)
	return rc.Response().JSON(http.StatusCreated, vnf_device_mapper)
}

// handleDeleteVnfDeviceMapper godoc
//
//	@Summary		Delete specified vnf device mapper
//	@Description	Delete the vnf device mapper identified by the provided id
//	@Tags			vnf_device_mapper
//	@Produce		json
//	@Param			id					path		string						true		"Vnf device mapper id"
//	@Success		200					{object}	types.VnfDeviceMapperDocs
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
//	@Router			/agora/vnfm/rest/v1/vnfmapper/:id [delete]
func handleDeleteVnfDeviceMapper(rc app.RouterContext) error {
	string_id := rc.Request().GetParam("id")
	log.Infof("[VnfDeviceMapperRouter - handleDeleteVnfDeviceMapper]: {id: %s}", string_id)
	id, err := strconv.Atoi(string_id)
	if err != nil {
		log.Errorf("[VnfDeviceMapperRouter - handleDeleteVnfDeviceMapper]: %s", err.Error())
		return rc.Response().Error(http.StatusBadRequest, http.StatusBadRequest, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	success, err := DeleteVnfDeviceMapper(id)
	if err != nil || !success {
		log.Errorf("[VnfDeviceMapperRouter - handleDeleteVnfDeviceMapper]: %s", err.Error())
		return rc.Response().Error(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), errorMessages.SomethingWentWrong(), "")
	}
	log.Infof("[VnfDeviceMapperRouter - handleDeleteVnfDeviceMapper] Deleted vnf_device_mapper {id: %d}", id)
	return rc.Response().JSON(http.StatusOK, fmt.Sprintf("Deleted vnf_device_mapper with id: %d", id))
}

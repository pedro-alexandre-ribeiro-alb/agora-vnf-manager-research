package validation

import (
	app "agora-vnf-manager/core/application"
	common "agora-vnf-manager/core/common"
	validation "agora-vnf-manager/core/common/validation"
	error "agora-vnf-manager/core/error"
	"agora-vnf-manager/core/utils"
	vnfDeviceMapperDao "agora-vnf-manager/features/vnf-device-mapper/dao"
	types "agora-vnf-manager/features/vnf-device-mapper/types"
	vnfInstanceDao "agora-vnf-manager/features/vnf-instance/dao"
	"strconv"
)

func ValidateReference(id int) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	vnf_device_mapper_exists, err := vnf_device_mapper_dao.Exists(id)
	if err != nil {
		return false, error.ResourceNotFound(common.ResourceVnfDeviceMapper, strconv.Itoa(id))
	}
	if !vnf_device_mapper_exists {
		return false, error.ResourceNotFound(common.ResourceVnfDeviceMapper, strconv.Itoa(id))
	}
	return true, ""
}

func ValidateVnfInstanceReference(vnf_instance_id int) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	vnf_instance_exists, err := vnf_instance_dao.Exists(vnf_instance_id)
	if err != nil {
		return false, error.ResourceNotFound(common.ResourceVnfInstance, strconv.Itoa(vnf_instance_id))
	}
	if !vnf_instance_exists {
		return false, error.ResourceNotFound(common.ResourceVnfInstance, strconv.Itoa(vnf_instance_id))
	}
	return true, ""
}

func ValidateProxyReference(proxy_id int) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	vnf_instance_exists, err := vnf_instance_dao.Exists(proxy_id)
	if err != nil {
		return false, error.ResourceNotFound(common.ResourceVnfInstance, strconv.Itoa(proxy_id))
	}
	if !vnf_instance_exists {
		return false, error.ResourceNotFound(common.ResourceVnfInstance, strconv.Itoa(proxy_id))
	}
	return true, ""
}

func ValidateUniqueDeviceId(device_id string) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	vnf_device_mapper := types.VnfDeviceMapper{}
	vnf_device_mapper.DeviceId.Scan(device_id)
	vnf_device_mappers, err := vnf_device_mapper_dao.Find(vnf_device_mapper)
	if err != nil {
		return false, error.ParameterUnique("deviceId")
	}
	if len(vnf_device_mappers) > 0 {
		return false, error.ParameterUnique("deviceId")
	}
	return true, ""
}

func ValidateVnfDeviceMapperComplete(vnf_device_mapper types.VnfDeviceMapper) (bool, string) {
	if !validation.ValidateNull(vnf_device_mapper) {
		return false, error.ParameterMandatory("vnf_device_mapper")
	}
	if !vnf_device_mapper.DeviceId.Present() {
		return false, error.ParameterMandatory("deviceId")
	} else {
		device_id := string(utils.First(vnf_device_mapper.DeviceId.Get()))
		valid, message := ValidateUniqueDeviceId(device_id)
		if !valid {
			return false, message
		}
	}
	if !vnf_device_mapper.VnfInstanceId.Present() {
		return false, error.ParameterMandatory("vnfInstanceId")
	} else {
		vnf_instance_id := int(utils.First(vnf_device_mapper.VnfInstanceId.Get()))
		valid, message := ValidateVnfInstanceReference(vnf_instance_id)
		if !valid {
			return false, message
		}
	}
	if !vnf_device_mapper.ProxyId.Present() {
		return false, error.ParameterMandatory("proxyId")
	} else {
		proxy_id := int(utils.First(vnf_device_mapper.ProxyId.Get()))
		valid, message := ValidateProxyReference(proxy_id)
		if !valid {
			return false, message
		}
	}
	return true, ""
}

func ValidateVnfDeviceMapperPartial(vnf_device_mapper types.VnfDeviceMapper) (bool, string) {
	if !validation.ValidateNull(vnf_device_mapper) {
		return false, error.ParameterMandatory("vnf_device_mapper")
	}
	if vnf_device_mapper.DeviceId.Present() {
		device_id := string(utils.First(vnf_device_mapper.DeviceId.Get()))
		valid, message := ValidateUniqueDeviceId(device_id)
		if !valid {
			return false, message
		}
	}
	if vnf_device_mapper.VnfInstanceId.Present() {
		vnf_instance_id := int(utils.First(vnf_device_mapper.VnfInstanceId.Get()))
		valid, message := ValidateVnfInstanceReference(vnf_instance_id)
		if !valid {
			return false, message
		}
	}
	if vnf_device_mapper.ProxyId.Present() {
		proxy_id := int(utils.First(vnf_device_mapper.ProxyId.Get()))
		valid, message := ValidateProxyReference(proxy_id)
		if !valid {
			return false, message
		}
	}
	return true, ""
}

func ValidateCreate(vnf_device_mapper types.VnfDeviceMapper) (bool, string) {
	return ValidateVnfDeviceMapperComplete(vnf_device_mapper)
}

func ValidateUpdate(vnf_device_mapper types.VnfDeviceMapper, id int) (bool, string) {
	valid, message := ValidateReference(id)
	if !valid {
		return false, message
	}
	return ValidateVnfDeviceMapperPartial(vnf_device_mapper)
}

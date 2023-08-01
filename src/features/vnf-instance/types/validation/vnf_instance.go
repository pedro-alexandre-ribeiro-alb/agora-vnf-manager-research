package validation

import (
	app "agora-vnf-manager/core/application"
	common "agora-vnf-manager/core/common"
	validation "agora-vnf-manager/core/common/validation"
	error "agora-vnf-manager/core/error"
	utils "agora-vnf-manager/core/utils"
	vnfInfrastructureDao "agora-vnf-manager/features/vnf-infrastructure/dao"
	vnfInstanceDao "agora-vnf-manager/features/vnf-instance/dao"
	types "agora-vnf-manager/features/vnf-instance/types"
	"strconv"
)

func ValidateReference(id int) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	vnf_instance_exists, err := vnf_instance_dao.Exists(id)
	if err != nil {
		return false, error.ResourceNotFound(common.ResourceVnfInstance, strconv.Itoa(id))
	}
	if !vnf_instance_exists {
		return false, error.ResourceNotFound(common.ResourceVnfInstance, strconv.Itoa(id))
	}
	return true, ""
}

func ValidateVnfInfrastructureReference(vnf_infra_id int) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	vnf_infrastructure_exists, err := vnf_infrastructure_dao.Exists(vnf_infra_id)
	if err != nil {
		return false, error.ResourceNotFound(common.ResourceVnfInfrastructure, strconv.Itoa(vnf_infra_id))
	}
	if !vnf_infrastructure_exists {
		return false, error.ResourceNotFound(common.ResourceVnfInfrastructure, strconv.Itoa(vnf_infra_id))
	}
	return true, ""
}

func ValidateVnfInstanceComplete(vnf_instance types.VnfInstance) (bool, string) {
	if !validation.ValidateNull(vnf_instance) {
		return false, error.ParameterMandatory("vnf_instance")
	}
	if !vnf_instance.Name.Present() {
		return false, error.ParameterMandatory("name")
	} else {
		name := string(utils.First(vnf_instance.Name.Get()))
		valid, message := ValidateUniqueName(name)
		if !valid {
			return false, message
		}
	}
	if !vnf_instance.Description.Present() {
		return false, error.ParameterMandatory("description")
	}
	if !vnf_instance.Type.Present() {
		return false, error.ParameterMandatory("type")
	} else {
		vnf_instance_type := string(utils.First(vnf_instance.Type.Get()))
		valid, message := ValidateVnfInstanceType(vnf_instance_type)
		if !valid {
			return false, message
		}
	}
	if !vnf_instance.VnfInfraId.Present() {
		return false, error.ParameterMandatory("vnfInfraId")
	} else {
		vnf_infra_id := int(utils.First(vnf_instance.VnfInfraId.Get()))
		valid, message := ValidateVnfInfrastructureReference(vnf_infra_id)
		if !valid {
			return false, message
		}
	}
	if !vnf_instance.Discovered.Present() {
		return false, error.ParameterMandatory("discovered")
	}
	if !vnf_instance.ManagementInterface.Present() {
		return false, error.ParameterMandatory("managementInterface")
	}
	if !vnf_instance.ControlInterface.Present() {
		return false, error.ParameterMandatory("controlInterface")
	}
	if !vnf_instance.Vendor.Present() {
		return false, error.ParameterMandatory("vendor")
	}
	if !vnf_instance.Version.Present() {
		return false, error.ParameterMandatory("version")
	}
	return true, ""
}

func ValidateVnfInstancePartial(vnf_instance types.VnfInstance) (bool, string) {
	if !validation.ValidateNull(vnf_instance) {
		return false, error.ParameterMandatory("vnf_instance")
	}
	if vnf_instance.Name.Present() {
		name := string(utils.First(vnf_instance.Name.Get()))
		valid, message := ValidateUniqueName(name)
		if !valid {
			return false, message
		}
	}
	if vnf_instance.Type.Present() {
		vnf_instance_type := string(utils.First(vnf_instance.Type.Get()))
		valid, message := ValidateVnfInstanceType(vnf_instance_type)
		if !valid {
			return false, message
		}
	}
	if vnf_instance.VnfInfraId.Present() {
		vnf_infra_id := int(utils.First(vnf_instance.VnfInfraId.Get()))
		valid, message := ValidateVnfInfrastructureReference(vnf_infra_id)
		if !valid {
			return false, message
		}
	}
	return true, ""
}

func ValidateVnfInstanceType(vnf_instance_type string) (bool, string) {
	if _, err := types.FromString(vnf_instance_type); err != nil {
		return false, error.ParameterValueInvalid(vnf_instance_type, "type")
	}
	return true, ""
}

func ValidateUniqueName(vnf_instance_name string) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	vnf_instance := types.VnfInstance{}
	vnf_instance.Name.Scan(vnf_instance_name)
	vnf_instances, err := vnf_instance_dao.Find(vnf_instance)
	if err != nil {
		return false, error.ParameterUnique("name")
	}
	if len(vnf_instances) > 0 {
		return false, error.ParameterUnique("name")
	}
	return true, ""

}

func ValidateCreate(vnf_instance types.VnfInstance) (bool, string) {
	return ValidateVnfInstanceComplete(vnf_instance)
}

func ValidateUpdate(vnf_instance types.VnfInstance, id int) (bool, string) {
	valid, message := ValidateReference(id)
	if !valid {
		return false, message
	}
	return ValidateVnfInstancePartial(vnf_instance)
}

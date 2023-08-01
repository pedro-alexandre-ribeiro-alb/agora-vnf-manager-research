package validation

import (
	app "agora-vnf-manager/core/application"
	common "agora-vnf-manager/core/common"
	validation "agora-vnf-manager/core/common/validation"
	error "agora-vnf-manager/core/error"
	utils "agora-vnf-manager/core/utils"
	vnfInfrastructureDao "agora-vnf-manager/features/vnf-infrastructure/dao"
	types "agora-vnf-manager/features/vnf-infrastructure/types"
	"os"
	"strconv"
)

func ValidateReference(id int) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	vnf_infrastructure_exists, err := vnf_infrastructure_dao.Exists(id)
	if err != nil {
		return false, error.ResourceNotFound(common.ResourceVnfInfrastructure, strconv.Itoa(id))
	}
	if !vnf_infrastructure_exists {
		return false, error.ResourceNotFound(common.ResourceVnfInfrastructure, strconv.Itoa(id))
	}
	return true, ""
}

func ValidateConfigureFilePath(configuration_file_path string) (bool, string) {
	if _, err := os.Stat(configuration_file_path); err != nil {
		return false, error.FileDoesNotExist(configuration_file_path)
	}
	return true, ""
}

func ValidateVnfInfrastructureComplete(vnf_infrastructure types.VnfInfrastructure) (bool, string) {
	if !validation.ValidateNull(vnf_infrastructure) {
		return false, error.ParameterMandatory("vnf_infrastructure")
	}
	if !vnf_infrastructure.Name.Present() {
		return false, error.ParameterMandatory("name")
	} else {
		name := string(utils.First(vnf_infrastructure.Name.Get()))
		valid, message := ValidateUniqueName(name)
		if !valid {
			return false, message
		}
	}
	if !vnf_infrastructure.Description.Present() {
		return false, error.ParameterMandatory("description")
	}
	if !vnf_infrastructure.ConfigurationFile.Present() {
		return false, error.ParameterMandatory("configurationFile")
	} else {
		configuration_file := string(utils.First(vnf_infrastructure.ConfigurationFile.Get()))
		valid, message := ValidateConfigureFilePath(configuration_file)
		if !valid {
			return false, message
		}
	}
	return true, ""
}

func ValidateVnfInfrastructurePartial(vnf_infrastructure types.VnfInfrastructure) (bool, string) {
	if !validation.ValidateNull(vnf_infrastructure) {
		return false, error.ParameterMandatory("vnf_infrastructure")
	}
	if vnf_infrastructure.Name.Present() {
		name := string(utils.First(vnf_infrastructure.Name.Get()))
		valid, message := ValidateUniqueName(name)
		if !valid {
			return false, message
		}
	}
	if vnf_infrastructure.ConfigurationFile.Present() {
		configuration_file := string(utils.First(vnf_infrastructure.ConfigurationFile.Get()))
		valid, message := ValidateConfigureFilePath(configuration_file)
		if !valid {
			return false, message
		}
	}
	return true, ""
}

func ValidateUniqueName(vnf_infrastructure_name string) (bool, string) {
	session := app.MyApp.Db.NewSession()
	session.Begin()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	vnf_infrastructure := types.VnfInfrastructure{}
	vnf_infrastructure.Name.Scan(vnf_infrastructure_name)
	vnf_infrastructures, err := vnf_infrastructure_dao.Find(vnf_infrastructure)
	if err != nil {
		return false, error.ParameterUnique("name")
	}
	if len(vnf_infrastructures) > 0 {
		return false, error.ParameterUnique("name")
	}
	return true, ""
}

func ValidateCreate(vnf_infrastructure types.VnfInfrastructure) (bool, string) {
	return ValidateVnfInfrastructureComplete(vnf_infrastructure)
}

func ValidateUpdate(vnf_infrastructure types.VnfInfrastructure, id int) (bool, string) {
	valid, message := ValidateReference(id)
	if !valid {
		return false, message
	}
	return ValidateVnfInfrastructurePartial(vnf_infrastructure)
}

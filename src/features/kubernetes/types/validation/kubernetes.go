package validation

import (
	validation "agora-vnf-manager/core/common/validation"
	errorMessages "agora-vnf-manager/core/error"
	types "agora-vnf-manager/features/kubernetes/types"
)

func ValidateSpecification(specification types.Specifications) (bool, string) {
	if !validation.ValidateNullOrEmptyString(specification.ConfigurationFile) {
		return false, errorMessages.ParameterMandatory("configurationFile")
	}
	return true, ""
}

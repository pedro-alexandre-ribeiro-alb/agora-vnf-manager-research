package validation

import (
	validation "agora-vnf-manager/core/common/validation"
	errorMessages "agora-vnf-manager/core/error"
	types "agora-vnf-manager/features/helm/types"
)

func ValidateCreateDeploymentSpecification(specification types.Specification) (bool, string) {
	if !validation.ValidateNullOrEmptyString(specification.Namespace) {
		return false, errorMessages.ParameterMandatory("namespace")
	}
	if !validation.ValidateNullOrEmptyString(specification.ReleaseName) {
		return false, errorMessages.ParameterMandatory("releaseName")
	}
	if !specification.ChartPath.Present() {
		return false, errorMessages.ParameterMandatory("chartPath")
	}
	return true, ""
}

func ValidateDeleteDeploymentSpecification(specification types.Specification) (bool, string) {
	if !validation.ValidateNullOrEmptyString(specification.Namespace) {
		return false, errorMessages.ParameterMandatory("namespace")
	}
	if !validation.ValidateNullOrEmptyString(specification.ReleaseName) {
		return false, errorMessages.ParameterMandatory("releaseName")
	}
	return true, ""
}

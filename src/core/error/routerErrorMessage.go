package error

import "fmt"

const (
	RouterErrorMessageResourceSomethingWentWrong string = "Something went wrong"
	RouterErrorMessageParameterMandatory         string = "The \"%s\" field must be provided"
	RouterErrorMessageParameterTypeInvalid       string = "The vnf type of vnf instance \"%d\" is invalid"
	RouterErrorMessageParameterUnique            string = "The \"%s\" field must be unique"
	RouterErrorMessageParameterValueInvalid      string = "The value \"%s\" for parameter \"%s\" is not supported"
	RouterErrorMessageResourceNotFound           string = "No record of %s could be found with id %s"
	RouterErrorMessageResourceNotFoundById       string = "No record of %s could be found with %s equal to %s"
	RouterErrorMessageUnexpectedInputFormat      string = "Unexpected format on the request body"
	RouterErrorMessageFileDoesNotExist           string = "The provided file \"%s\" does not exist"
)

func SomethingWentWrong() string {
	return RouterErrorMessageResourceSomethingWentWrong
}

func ParameterMandatory(parameter string) string {
	return fmt.Sprintf(RouterErrorMessageParameterMandatory, parameter)
}

func ParameterInvalid(vnf_id int) string {
	return fmt.Sprintf(RouterErrorMessageParameterTypeInvalid, vnf_id)
}

func ParameterUnique(parameter string) string {
	return fmt.Sprintf(RouterErrorMessageParameterUnique, parameter)
}

func ParameterValueInvalid(value string, parameter string) string {
	return fmt.Sprintf(RouterErrorMessageParameterValueInvalid, value, parameter)
}

func ResourceNotFound(resourceType string, resourceId string) string {
	return fmt.Sprintf(RouterErrorMessageResourceNotFound, resourceType, resourceId)
}

func UnexpectedInputFormat() string {
	return RouterErrorMessageUnexpectedInputFormat
}

func FileDoesNotExist(file string) string {
	return fmt.Sprintf(RouterErrorMessageFileDoesNotExist, file)
}

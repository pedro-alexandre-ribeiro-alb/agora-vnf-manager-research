package validation

func ValidateNull(object any) bool {
	return object != nil
}

func ValidateNullOrEmptyString(value string) bool {
	if ValidateNull(value) {
		return len(value) != 0
	}
	return false
}

func ValidateNullOrZero(value int) bool {
	if value != 0 {
		return true
	}
	return false
}

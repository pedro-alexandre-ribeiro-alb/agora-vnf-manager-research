package error

type RouterResponseError struct {
	ErrorCode int    `json:"errorCode"`
	Cause     string `json:"cause"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

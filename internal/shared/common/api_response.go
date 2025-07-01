package common

type ApiResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewApiResponse(data interface{}) *ApiResponse {
	return &ApiResponse{
		Message: "success",
		Data:    data,
	}
}

func NewApiResponseError(error string) *ApiResponse {
	return &ApiResponse{
		Message: "error",
		Error:   error,
	}
}

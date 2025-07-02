package common

type ApiResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewApiResponse(data interface{}) *ApiResponse {
	return &ApiResponse{
		Message: "success",
		Data:    data,
	}
}

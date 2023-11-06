package http

type successResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Err     string `json:"error"`
}

func GetSuccessResponse(data interface{}) successResponse {
	return successResponse{
		Status:  "success",
		Message: "Request was successful",
		Data:    data,
	}
}

func GetErrorResponse(err string) errorResponse {
	return errorResponse{
		Status:  "error",
		Message: "Request failed",
		Err:     err,
	}
}

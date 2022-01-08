package rest

type RESTResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildRESTResponse(message string, data interface{}) RESTResponse {
	return RESTResponse{
		Message: message,
		Data:    data,
	}
}

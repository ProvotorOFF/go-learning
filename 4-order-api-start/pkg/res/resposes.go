package res

type SuccessResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Message any `json:"message"`
}

package response

const HttpStatus400 = 400

type ApiResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type CustomApiResponse map[string]interface{}

type ApiErrorResponse struct {
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

package response

type BaseErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

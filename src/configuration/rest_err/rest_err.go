package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []string `json:"causes"`
}

type Causes struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes []string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return NewRestErr(message, "bad request", http.StatusBadRequest, nil)
}

func NewNotFoundError(message string) *RestErr {
	return NewRestErr(message, "not found", http.StatusNotFound, nil)
}

func NewInternalServerError(message string) *RestErr {
	return NewRestErr(message, "internal server error", http.StatusInternalServerError, nil)
}

func NewUnauthorizedError(message string) *RestErr {
	return NewRestErr(message, "unauthorized", http.StatusUnauthorized, nil)
}

func NewForbiddenError(message string) *RestErr {
	return NewRestErr(message, "forbidden", http.StatusForbidden, nil)
}

func NewBadRequestValidatioError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request validation",
		Code:    http.StatusBadRequest,
		Causes:  []string{},
	}
}

package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"budgetapp/exception"
)

type Response struct {
	Code      int         `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	RequestID string      `json:"request_id,omitempty"`

	// can be used when returning multiple form errors
	Errors []ErrorInfo `json:"errors,omitempty"`
}

// ErrorInfo specifies what info are we sending.
// Use IsEmpty method instead of comparing with struct literal.
type ErrorInfo struct {
	Field    string          `json:"field"`
	Message  string          `json:"message"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}

func FormattedErrors(errors []ErrorInfo) (message string) {
	for _, err := range errors {
		message += fmt.Sprintf("%s %s ", err.Field, err.Message)
	}

	return
}

func RespondWithJSON(rw http.ResponseWriter, status int, response Response) {
	response.RequestID = GetRequestId(rw)
	respBytes, err := json.Marshal(response)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func GetRequestId(rw http.ResponseWriter) string {
	return rw.Header().Get("X-Request-ID")

}

func RespondWithException(rw http.ResponseWriter, err error) {
	ex := exception.ToHttpExceptionFromError(err)

	RespondWithJSON(rw, ex.Status(), Response{
		Code:    ex.GetCode(),
		Message: ex.GetMessage(),
	})
}

func RespondWithError(rw http.ResponseWriter, status int, response Response) {
	response.RequestID = GetRequestId(rw)

	respBytes, err := json.Marshal(response)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

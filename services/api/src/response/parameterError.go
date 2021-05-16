package response

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ParameterErrorResponse struct {
	hasError        bool
	writer          http.ResponseWriter
	ParameterErrors []*ParameterError `json:"parameter_errors"`
}

func NewParameterErrorResponse(w http.ResponseWriter) *ParameterErrorResponse {
	return &ParameterErrorResponse{
		hasError:        false,
		writer:          w,
		ParameterErrors: []*ParameterError{},
	}
}

func (r *ParameterErrorResponse) AddError(error *ParameterError) {
	r.hasError = true
	r.ParameterErrors = append(r.ParameterErrors, error)
}

func (r *ParameterErrorResponse) PrintWarningToConsole() {

	bytes, err := json.Marshal(r.ParameterErrors)

	if err != nil {
		logrus.Error(err)
	}

	logrus.Warn(string(bytes))

}

type ParameterError struct {
	Msg       string `json:"msg"`
	ErrorType string `json:"error_type"`
	Field     string `json:"field"`
}

func (r *ParameterErrorResponse) ValidateStringFieldIsNotMissingOrEmpty(value, fieldName string) {
	if len(value) == 0 {
		r.AddError(NewParameterMissingError(fieldName))
	}
}

func (r *ParameterErrorResponse) WriteIfHasError() (didWrite bool) {
	if r.hasError {
		if err := r.WriteResponse(); err != nil {
			logrus.Error(err)
			r.writer.WriteHeader(http.StatusInternalServerError)
		}
		return true

	}
	return false
}

func (r *ParameterErrorResponse) WriteResponse() error {
	r.writer.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(r.writer).Encode(r)
}

func (r *ParameterErrorResponse) HasError() bool {
	return r.hasError
}

func NewParameterMissingError(field string) *ParameterError {
	return &ParameterError{
		Msg:       fmt.Sprintf("%s_is_required", field),
		ErrorType: "required",
		Field:     field,
	}
}

func NewParameterAlreadyInUseError(field string) *ParameterError {
	return &ParameterError{
		Msg:       fmt.Sprintf("%s_is_alreay_in_use", field),
		ErrorType: "already_in_use",
		Field:     field,
	}
}

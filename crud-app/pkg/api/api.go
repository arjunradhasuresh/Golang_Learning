package api

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Status string          `json:"status"`
	Result json.RawMessage `json:"message"`
	Error  *ErrorResponse  `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

const (
	statusOK                = "OK"
	statusError             = "ERROR"
	SuccessCode             = 200
	NotFoundCode            = 404
	InternalServerErrorCode = 500
)

func Success(w http.ResponseWriter, status string, result interface{}) {
	var r *SuccessResponse
	if result != nil {
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Failed to marshal result", InternalServerErrorCode)
			return
		}
		r = &SuccessResponse{
			Status: status,
			Result: data,
		}

	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, "Failed to marshal response", InternalServerErrorCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(SuccessCode)
	_, err = w.Write(j)
	if err != nil {
		http.Error(w, "Failed to write response", InternalServerErrorCode)
		return
	}
}

func Fail(w http.ResponseWriter, code int, message string, details ...string) {
	var r *SuccessResponse

	r = &SuccessResponse{
		Status: statusError,
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
			Details: details[0],
		},
	}
	if len(details) > 0 {
		r.Error.Details = details[0]
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, "Failed to marshal error response", InternalServerErrorCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(j)
	if err != nil {
		http.Error(w, "Failed to write error response", InternalServerErrorCode)
		return
	}
}

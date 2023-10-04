package handlers

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
)

const (
	ELEMENT_NOT_FOUND    = "ELEMENT_NOT_FOUND"
	NODE_NOT_FOUND       = "NODE_NOT_FOUND"
	EDGE_NOT_FOUND       = "EDGE_NOT_FOUND"
	UNKNOWN_ERROR        = "UNKNOWN_ERROR"
	KNOWN_ERROR          = "KNOWN_ERROR"
	UNAUTHORIZED         = "UNAUTHORIZED"
	INVALID_REQUEST_BODY = "INVALID_REQUEST_BODY"
	RATE_LIMIT_EXCEEDED  = "RATE_LIMIT_EXCEEDED"
)

type Error struct {
	Message string   `json:"message"`
	Code    string   `json:"code"`
	Path    []string `json:"path"`
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Errors []Error     `json:"errors"`
}

type Responder struct {
	W        http.ResponseWriter
	response Response
}

func NewResponder(w http.ResponseWriter) *Responder {
	return &Responder{
		W: w,
	}
}

func (r *Responder) AddError(err error, code string, path []string) error {
	r.response.Errors = append(r.response.Errors, Error{
		Message: err.Error(),
		Code:    code,
		Path:    path,
	})
	return err
}

func (r *Responder) Status(status int) {
	r.response.Status = status
}

func (r *Responder) Exec(res *graphql.Result) {
	r.W.Header().Set("Content-Type", "application/json")
	if res.Errors == nil {
		r.response.Status = http.StatusOK
	} else if r.response.Status == 0 {
		r.response.Status = http.StatusInternalServerError
	}
	r.W.WriteHeader(r.response.Status)

	r.response.Data = res.Data
	for _, err := range res.Errors {
		for i, err2 := range r.response.Errors {
			r.response.Errors[i] = Error{
				Message: err.Message,
				Code:    err2.Code,
				Path:    err2.Path,
			}
		}
	}

	json.NewEncoder(r.W).Encode(r.response)
}

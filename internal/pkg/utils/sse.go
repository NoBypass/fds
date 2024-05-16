package utils

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SSEConn struct {
	encoder *json.Encoder
	resp    *echo.Response
}

func NewSSEConn(resp *echo.Response) *SSEConn {
	resp.Header().Set(echo.HeaderContentType, "text/event-stream")
	resp.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(resp)
	return &SSEConn{
		encoder: enc,
		resp:    resp,
	}
}

func (s *SSEConn) Send(data any) error {
	err := s.encoder.Encode(data)
	if err != nil {
		return err
	}

	s.resp.Flush()
	return nil
}

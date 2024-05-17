package utils

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SSEConn struct {
	encoder *json.Encoder
	resp    *echo.Response
}

func NewSSEConn(resp *echo.Response) *SSEConn {
	resp.Header().Set(echo.HeaderContentType, "text/event-stream")
	enc := json.NewEncoder(resp)
	return &SSEConn{
		encoder: enc,
		resp:    resp,
	}
}

func (s *SSEConn) Send(status int, data any) error {
	s.resp.WriteHeader(status)
	err := s.encoder.Encode(data)
	if err != nil {
		return err
	}

	s.resp.Flush()
	return nil
}

func (s *SSEConn) Err(err error, ctx echo.Context) error {
	e := ctx.Echo()

	var he *echo.HTTPError
	ok := errors.As(err, &he)
	if ok {
		if he.Internal != nil {
			var herr *echo.HTTPError
			if errors.As(he.Internal, &herr) {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	message := he.Message
	switch m := he.Message.(type) {
	case string:
		if e.Debug {
			message = map[string]any{"message": m, "error": err.Error()}
		} else {
			message = map[string]any{"message": m}
		}
	case json.Marshaler:
	case error:
		message = map[string]any{"message": m.Error()}
	}

	return s.Send(he.Code, message)
}

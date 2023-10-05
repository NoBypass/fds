package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"server/internal/app/handlers"
	"strings"
	"time"
)

type ResponseRecorder struct {
	http.ResponseWriter
	Status int
	Size   int
}

func (r *ResponseRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.Size += size
	return size, err
}

var dark = color.New(color.FgHiBlack)
var pink = color.New(color.BgHiMagenta, color.FgHiWhite)
var white = color.New(color.FgHiWhite)
var red = color.New(color.FgHiRed, color.Bold)
var green = color.New(color.FgHiGreen, color.Bold)
var yellow = color.New(color.FgHiYellow, color.Bold)
var cyan = color.New(color.FgHiCyan, color.Bold)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &ResponseRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
			Size:           0,
		}

		var requestBody handlers.GraphQLBody
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "requestBody", requestBody))

		next.ServeHTTP(recorder, r)

		duration := time.Since(start)
		formattedTime := start.Format("02/01/06 15:04:05")
		dark.Print(formattedTime, " ")
		pink.Print(" GraphQL ")
		white.Print(" | Status: ")
		if recorder.Status < 400 {
			green.Print(recorder.Status)
		} else {
			red.Print(recorder.Status)
		}

		white.Print(" | Duration: ")
		if duration < 30*time.Millisecond {
			green.Print(duration)
		} else if duration < 100*time.Millisecond {
			yellow.Print(duration)
		} else {
			red.Print(duration)
		}

		white.Print(" | Size: ")
		cyan.Printf("%dB", recorder.Size)
		white.Print(" | Actions: ")

		actions := formatGraphQLVariables(requestBody)
		for i, action := range actions {
			if i < 5 {
				cyan.Print(action)
			} else {
				dark.Printf(" and %d more", len(actions)-5)
			}
		}

		fmt.Println()
	})
}

func formatGraphQLVariables(body handlers.GraphQLBody) []string {
	var res []string

	lines := strings.Split(body.Query, "\n")[1:]
	for _, line := range lines {
		if strings.Contains(line, "{") {
			res = append(res, strings.TrimSpace(strings.Split(line, "(")[0]))
		}
	}

	return res
}

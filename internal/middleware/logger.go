package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "requestID"

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(data []byte) (int, error) {

	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}

	return rw.ResponseWriter.Write(data)
}

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		requestID := uuid.New().String()

		w.Header().Set("X-Request-ID", requestID)

		ctx := context.WithValue(
			r.Context(),
			requestIDKey,
			requestID,
		)

		r = r.WithContext(ctx)

		fmt.Printf(
			"[request_id=%s] %s %s\n",
			requestID,
			r.Method,
			r.URL.Path,
		)

		rw := &responseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		fmt.Printf(
			"[request_id=%s] status=%d duration=%s\n",
			requestID,
			rw.statusCode,
			duration,
		)
	})
}

func GetRequestID(r *http.Request) string {

	requestID, ok := r.Context().Value(requestIDKey).(string)

	if !ok {
		return ""
	}

	return requestID
}
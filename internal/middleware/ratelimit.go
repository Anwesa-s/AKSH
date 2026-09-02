package middleware

import (
	"net"
	"net/http"

	"github.com/Anwesa-s/AKSH/internal/ratelimit"
)

func RateLimit(limiter *ratelimit.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Extract client IP from RemoteAddr
			host, _, err := net.SplitHostPort(r.RemoteAddr)

			if err != nil {
				host = r.RemoteAddr
			}

			if !limiter.Allow(host) {
				w.Header().Set("Retry-After", "1")
				http.Error(
					w,
					"Too Many Requests",
					http.StatusTooManyRequests,
				)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
package middleware

import (
	"net"
	"net/http"
	"strings"
)

type RateLimiter interface {
	Allow(key string) (bool, error)
}

func RateLimit(limiter RateLimiter) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ip, _, err := net.SplitHostPort(r.RemoteAddr)

			if err != nil {
				ip = strings.Split(r.RemoteAddr, ":")[0]
			}

			allowed, err := limiter.Allow(ip)

			if err != nil {
				http.Error(
					w,
					"Rate limiter error",
					http.StatusInternalServerError,
				)
				return
			}

			if !allowed {
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
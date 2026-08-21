package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target string) (*httputil.ReverseProxy, error) {

	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {

		fmt.Println("Proxy error:", err)

		http.Error(
			w,
			"AKSH: backend service unavailable",
			http.StatusBadGateway,
		)
	}

	return proxy, nil
}
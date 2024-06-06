package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoute(router *http.ServeMux) {
	router.Handle("GET /metrics", promhttp.Handler())
}

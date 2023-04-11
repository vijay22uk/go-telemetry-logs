package ping

import (
	"context"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	metricsHandler := promhttp.Handler()
	r := mux.NewRouter()
	r.Handle("/metrics", metricsHandler)
	r.Use(commonMiddleware, prometheusMiddleware)

	r.Methods("POST").Path("/ping").Handler(httptransport.NewServer(
		endpoints.SayHello,
		decodeReq,
		encodeResponse,
	))

	// r.Methods("GET").Path("/ping/{name}").Handler(httptransport.NewServer(
	// 	endpoints.GetUser,
	// 	decodeEmailReq,
	// 	encodeResponse,
	// ))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})

}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		httpRequestCount.WithLabelValues(r.Method, r.URL.Path).Inc()
		httpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration.Seconds())
	})
}

package main

import (
	"context"
	"flag"
	"fmt"
	"gokits/ping"
	"gokits/telemetry"
	"net/http"

	"github.com/go-kit/kit/log"
	"go.opentelemetry.io/otel"

	"github.com/go-kit/kit/log/level"

	"os"
	"os/signal"
	"syscall"
)

// https://github.com/tensor-programming/go-kit-tutorial

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
		//logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "ping",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// TRacer
	tp, err := telemetry.InitTracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		panic(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	// Tracer End

	flag.Parse()
	ctx := context.Background()
	var srv ping.Service
	srv = ping.NewPingService(logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := ping.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := ping.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}

package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/martinsirbe/go-metrics-poc/db/postgres"
	"github.com/martinsirbe/go-metrics-poc/examples/service"
)

func main() {
	storageClient := postgres.NewInstrumentedClientPrometheus()
	http.Handle("/metrics", promhttp.Handler())

	// Example for other clients, such as Graphite, DataDog where the client post metrics
	//storageClient := postgres.InstrumentedClientOther()

	ll := service.NewLogicLayer(storageClient)
	go func() {
		if err := ll.Run(); err != nil {
			logrus.WithError(err).Fatal("failed to run")
		}
	}()

	if err := http.ListenAndServe(":1337", nil); err != nil {
		logrus.WithError(err).Fatal("failed to listen and serve metrics")
	}
}

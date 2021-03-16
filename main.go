package main

import (
	"context"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
	"knative.dev/pkg/logging"
)

// Receiver q
type Receiver struct {
	httpC http.Client

	logger *zap.SugaredLogger
}

// Request is the structure of the event we expect to receive.
type Request struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

func main() {
	ctx := context.Background()
	r := Receiver{}
	r.logger = logging.FromContext(ctx)
	if err := envconfig.Process("", &r); err != nil {
		log.Fatal(err)
	}

	log.Info("Starting the HTTP server now..")
	http.HandleFunc("/", r.getPricing)
	http.ListenAndServe(":8080", nil)
}

func (recv *Receiver) getPricing(w http.ResponseWriter, r *http.Request) {
	MetalLegendPricing(1000, 2000, 3000)
}

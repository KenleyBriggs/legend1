package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type PricingData struct {
	Three int `json:"three"`
	Six   int `json:"six"`
	Ten   int `json:"ten"`
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
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error occured reading body: %v", err)
		return
	}

	i := &PricingData{}
	if err := json.Unmarshal(body, i); err != nil {
		log.Errorf("Error occured unmarsaling data: %v", err)
		return
	}

	MetalLegendPricing(i.Three, i.Six, i.Ten)
}

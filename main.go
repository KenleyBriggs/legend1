package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/prometheus/common/log"
)

const (
	envUsername = "SQL_USERNAME"
	envPassword = "SQL_PASSWORD"
)

type Adapter struct {
	Username string
	Password string

	Three int
	Six   int
	Ten   int
}

func init() {

}

type PricingData struct {
	Three int `json:"three"`
	Six   int `json:"six"`
	Ten   int `json:"ten"`
}

func main() {

	a := &Adapter{}

	a.Username = os.Getenv(envUsername)
	a.Password = os.Getenv(envPassword)

	fmt.Println(a.Username)

	log.Info("Starting the HTTP server now..")
	http.HandleFunc("/pricing", a.getPricing)
	http.ListenAndServe(":8080", nil)

	// DataBase()
}

func (a *Adapter) getPricing(w http.ResponseWriter, r *http.Request) {
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

	// call sql for updated values

	a.Three = 12
	a.Ten = 122
	a.Six = 322

	a.MetalLegendPricing(i.Three, i.Six, i.Ten)

	fmt.Fprintf(w, "got pricing!\n")
}

func (a *Adapter) MetalLegendPricing(three int, six int, ten int) {

	metal28GaTuffRib := ((three * a.Three) + (six * a.Six) + (ten * a.Ten))
	metal26GaTuffRib := ((three * 975) + (six * 1005) + (ten * 1035))
	metal26Ga100NS := ((three * 1050) + (six * 1580) + (ten * 1610))
	metal24Ga100NS := ((three * 1650) + (six * 1680) + (ten * 1710))

	fmt.Print(metal28GaTuffRib)
	metal28GaTuffRibString := FormatCurrency(metal28GaTuffRib)
	metal26GaTuffRibString := FormatCurrency(metal26GaTuffRib)
	metal26Ga100NSString := FormatCurrency(metal26Ga100NS)
	metal24Ga100NSString := FormatCurrency(metal24Ga100NS)

	fmt.Println("Legend Metal Pricing:")
	fmt.Println("28 gauge Tuff Rib:", metal28GaTuffRibString)
	fmt.Println("26 gauage Tuff Rib:", metal26GaTuffRibString)
	fmt.Println("26 guage 1 inch Nail Strip Standing Seam:", metal26Ga100NSString)
	fmt.Println("24 gauge 1 inch Nail Strip Standing Seam:", metal24Ga100NSString)

}

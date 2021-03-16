package main

import (
	"fmt"
)

//this function generates the pricng for the lengend system based on 3/12 6/12 and 10/12 pitches
func MetalLegendPricing(three int, six int, ten int) {
	
	metal28GaTuffRib := ((three * 875) + (six * 905) + (ten * 935))
	metal26GaTuffRib := ((three * 975) + (six * 1005) + (ten * 1035))
	metal26Ga100NS := ((three * 1050) + (six * 1580) + (ten * 1610))
	metal24Ga100NS := ((three * 1650) + ( six * 1680) + (ten * 1710))

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
package main

import (
	"fmt"
	//"go/constant"
	"strings"
	"strconv"
)

func DecimalFunc(slice []string, element string) []string {

    n := len(slice)
    if n == cap(slice) {
        newSlice := make([]string, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = slice[n-2]
	slice[n-2] = element
	
	
    return slice
}

func CommaFunc(slice []string, element string) string {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]string, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	
	length := len(slice)
	for i := 7; i < length; i = i + 4 {

		slice = slice [0 : n+1]
		slice[n] = slice[n-i]
		slice[n-i] = ","
	}
	sliceToString := strings.Join(slice, "")
	return sliceToString
}

//this function generates the pricng for the lengend system based on 3/12 6/12 and 10/12 pitches
func LegendPricing(three int, six int, ten int) {
	
	metal28GaTuffRib := ((three * 875) + (six * 905) + (ten * 935))
	metal26GaTuffRib := ((three * 975) + (six * 1005) + (ten * 1035))
	metal26Ga100NS := ((three * 1050) + (six * 1580) + (ten * 1610))
	metal24Ga100NS := ((three * 1650) + ( six * 1680) + (ten * 1710))
	
	metal28GaTuffRibString := strconv.Itoa(metal28GaTuffRib)
	newmetal := strings.Split(metal28GaTuffRibString, "")
	newmetal1 := DecimalFunc(newmetal, ".")
	newmetal2 := CommaFunc(newmetal1, ",")
	

	fmt.Println(newmetal1)
	fmt.Println(newmetal2)
	fmt.Println("28 gauge Tuff Rib: $", metal28GaTuffRibString)
	fmt.Println("26 gauage Tuff Rib: $", metal26GaTuffRib)
	fmt.Println("26 guage 1 inch Nail Strip Standing Seam: $", metal26Ga100NS)
	fmt.Println("24 gauge 1 inch Nail Strip Standing Seam: $", metal24Ga100NS)

}
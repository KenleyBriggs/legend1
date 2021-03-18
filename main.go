package main

type Adapter struct  {
	Three int 
	Six int
	Ten int 
}

func init (){
	
}

func main() {

	a := &Adapter{}
	// ask sql for updated pricing
	a.Three = 1234
	a.Six = 4567
	a.Ten = 8900
	a.MetalLegendPricing(1000, 2000, 3000)
	DataBase()
}

func (a *Adapter) MetalLegendPricing(three int, six int, ten int) {

	metal28GaTuffRib := ((three * a.Three) + (six * a.Six) + (ten * a.Ten))
	metal26GaTuffRib := ((three * 975) + (six * 1005) + (ten * 1035))
	metal26Ga100NS := ((three * 1050) + (six * 1580) + (ten * 1610))
	metal24Ga100NS := ((three * 1650) + (six * 1680) + (ten * 1710))

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

package main

import (
	"strconv"
	"strings"
)

func FormatCurrency(currency int) string {
	// is the integer price
	ammountInt := currency
	// converts the integer into a string
	ammountStr := strconv.Itoa(ammountInt)
	// splits the string into an array
	ammountArray := strings.Split(ammountStr, "")

	// determines the length of the array
	length := len(ammountArray)
	// accounts for the two decimal places and the first three integers
	if length > 6 {
		// determines the place in the array to insert the comma
		s := length - 5

		// calls on the insert function that shifts the array and inserts the element
		tempArray := insert(ammountArray, ",", s)
		// replaces the original array
		ammountArray = tempArray
	}

	// accounts for the two decimal places and the second three integers
	if length > 9 {
		// determines the place in the array to insert the comma
		s := length - 8

		// calls on the insert function that shifts the array and inserts the element
		tempArray := insert(ammountArray, ",", s)
		// replaces the original array
		ammountArray = tempArray
	}

	// accounts for the two decimal places and the third three integers
	if length > 13 {
		// determines the place in the array to insert the comma
		s := length - 11

		// calls on the insert function that shifts the array and inserts the element
		tempArray := insert(ammountArray, ",", s)
		// replaces the original array
		ammountArray = tempArray
	}

	// accounts for the two decimal places
	if length > 2 {
		// determines the place in the array to insert the period
		s := len(ammountArray) - 2

		// calls on the insert function that shifts the array and inserts the element
		tempArray := insert(ammountArray, ".", s)
		// replaces the original array
		ammountArray = tempArray
	}

	// accounts for the two decimal places and the first three integers
	if length > 2 {
		// sets temparray as ammountarray
		tempArray := ammountArray
		// appends a dollar sign to the front of the array
		tempArray = append([]string{"$"}, tempArray...)
		// replaces the original array
		ammountArray = tempArray
	}

	// converts the array back into a string to print
	tempArray := stringconv(ammountArray)

	// returns the temparray string
	return tempArray
}

// insert function shifts the array and inserts the element
func insert(a []string, c string, i int) []string {
	return append(a[:i], append([]string{c}, a[i:]...)...)
}

// stringconv function converts the array into a string
func stringconv(as []string) string {
	result1 := strings.Join(as, "")
	return result1
}

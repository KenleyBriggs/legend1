package main

import (
	"strconv"
	"strings"
)

func FormatCurrency(currency int) string {
	ammountInt := currency                        //is the integer price
	ammountStr := strconv.Itoa(ammountInt)        //converts the integer into a string
	ammountArray := strings.Split(ammountStr, "") //splits the string into an array

	length := len(ammountArray) //determines the length of the array
	if length > 6 {             //accounts for the two decimal places and the first three integers
		s := length - 5 //determines the place in the array to insert the comma

		tempArray := insert(ammountArray, ",", s) //calls on the insert function that shifts the array and inserts the element
		ammountArray = tempArray                  //replaces the original array
	}

	if length > 9 { //accounts for the two decimal places and the second three integers
		s := length - 8 //determines the place in the array to insert the comma

		tempArray := insert(ammountArray, ",", s) //calls on the insert function that shifts the array and inserts the element
		ammountArray = tempArray                  //replaces the original array
	}

	if length > 13 { //accounts for the two decimal places and the third three integers
		s := length - 11 //determines the place in the array to insert the comma

		tempArray := insert(ammountArray, ",", s) //calls on the insert function that shifts the array and inserts the element
		ammountArray = tempArray                  //replaces the original array
	}

	if length > 2 { //accounts for the two decimal places
		s := len(ammountArray) - 2 //determines the place in the array to insert the period

		tempArray := insert(ammountArray, ".", s) //calls on the insert function that shifts the array and inserts the element
		ammountArray = tempArray                  //replaces the original array
	}
	if length > 2 { //accounts for the two decimal places and the first three integers
		tempArray := ammountArray                       //sets temparray as ammountarray
		tempArray = append([]string{"$"}, tempArray...) //appends a dollar sign to the front of the array
		ammountArray = tempArray                        //replaces the original array
	}
	tempArray := stringconv(ammountArray) //converts the array back into a string to print

	return tempArray //returns the temparray string
}

func insert(a []string, c string, i int) []string { //insert function shifts the array and inserts the element
	return append(a[:i], append([]string{c}, a[i:]...)...)
}

func stringconv(as []string) string { //stringconv function converts the array into a string
	result1 := strings.Join(as, "")
	return result1
}

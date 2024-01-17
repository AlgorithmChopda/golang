/*
	1. Write a program to calculate the simple interest
	First-line has the comma-separated values of the Principal, rate and time (in years) respective
	constraints: Round simple interest float value to 2 decimal places
*/

package main

import (
	"fmt"
	"math"
)

func roundFloatTillTwoDecimal(interest float64) float64 {
	return math.Round(interest*100) / 100
}

func calSimpleInterest(principleAmount, rateOfInterest, time float64) float64 {
	var interest = (principleAmount * rateOfInterest * time) / 100
	return interest
}

func main() {
	var principleAmount, rateOfInterest, time float64
	fmt.Println("Enter Principle Amount, Rate Of Interest and Time ")
	fmt.Scanln(&principleAmount, &rateOfInterest, &time)

	fmt.Println("Simple Interest : ", calSimpleInterest(principleAmount, rateOfInterest, time))
}

/*
	2. Write a program to calculate the area of the circle, First line has a value of the radius of the circle
	constraint

	1. Use const PI from the package math package
	2. Use the Pow function from the math package
	3. Round area float value to 2 decimal places
*/

package main

import (
	"fmt"
	"math"
)

func roundFloatTillTwoDecimal(interest float64) float64 {
	return math.Round(interest*100) / 100
}

func calAreaOfCircle(radius float64) float64 {
	var areaOfCircle = math.Pi * math.Pow(radius, 2)
	return roundFloatTillTwoDecimal(areaOfCircle)
}

func main() {
	var radius float64
	fmt.Println("Enter radius of Circle : ")
	fmt.Scanln(&radius)

	areaOfCircle := calAreaOfCircle(radius)
	fmt.Println("Area of Circle : ", areaOfCircle)
}

/*
	2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively.

	Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.

	Hint:
	Method Signatures for Area and Perimeter
	Area() int
	Perimeter() int

	Formulae:
	Area = length * width
	Perimeter = 2 * (length + width)

	Example Test Case 1:
	Input: 10 20
	Output:
	Area of Rectangle: 200
	Perimeter of Rectangle: 60

*/

package main

import "fmt"

type Rectangle struct {
	length, breadth int
}

func (rect Rectangle) Area() int {
	area := rect.length * rect.breadth
	return area
}
func (rect Rectangle) Perimeter() int {
	perimeter := 2 * (rect.length * rect.breadth)
	return perimeter
}

func main() {
	var rect Rectangle
	fmt.Printf("Enter length and breadth : ")
	fmt.Scanln(&rect.length, &rect.breadth)
	if rect.breadth < 0 || rect.length < 0 {
		fmt.Printf("length or breadth cannot be negative")
		return
	}

	fmt.Printf("\nArea of Rectangle : %d", rect.Area())
	fmt.Printf("\nPerimeter of Rectangle : %d", rect.Perimeter())
}

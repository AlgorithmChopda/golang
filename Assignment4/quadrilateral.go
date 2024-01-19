/*
	3. The given program takes an integer value as input from the user i.e 1 or 2.
	Option 1 represents Rectangle and Option 2 represents Square.

	Given the metrics of the shapes are hard-coded, complete the given program to achieve the following:

	1. Create an interface Quadrilateral which has the following method signatures
	* Area() int
	* Perimeter() int

	2. Implement the Quadrilateral interface for the given shapes i.e Rectangle and Square

	3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:

	"Area :  <value>"
	"Perimeter :  <value>"
	HINT: Step 2 means, to define methods in Quadrilateral interface for given shapes

	Formulae:

	Area of Rectangle: length * width
	Perimeter of Rectangle: 2*(length + breadth)

	Area of Square: side * side
	Perimeter of Square: 4 * side

*/

package main

import "fmt"

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

type Square struct {
	side int
}

func (square Square) Area() int {
	return square.side * square.side
}
func (square Square) Perimeter() int {
	return square.side * 4
}

type Rectangle struct {
	length, breadth int
}

func (rect Rectangle) Area() int {
	return rect.length * rect.breadth
}
func (rect Rectangle) Perimeter() int {
	return 2 * (rect.length + rect.breadth)
}

func Print(shape Quadrilateral) {
	fmt.Printf("\nArea: %d", shape.Area())
	fmt.Printf("\nPerimeter: %d", shape.Perimeter())
}

func processChoice(input int) {
	switch input {
	case 1:
		//Rectangle
		var rect Rectangle
		fmt.Printf("Enter length and breadth : ")
		fmt.Scanln(&rect.length, &rect.breadth)
		if rect.breadth < 0 || rect.length < 0 {
			fmt.Printf("length or breadth cannot be negative")
			return
		}

		Print(rect)

	case 2:
		// Square
		var square Square
		fmt.Printf("Enter side of square :")
		fmt.Scanln(&square.side)
		if square.side < 0 {
			fmt.Printf("side cannot be negative")
			return
		}

		Print(square)

	default:
		fmt.Printf("input should be 1 or 2 only.")
	}
}

func main() {
	var choice int
	fmt.Printf("\n1. Square \n2. Rectangle")
	fmt.Printf("\n Enter your choice : ")
	fmt.Scanln(&choice)

	processChoice(choice)
}

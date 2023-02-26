// Golang Interview questions Series 1.

// 1. What is the static variable declaration in Golang?
// 2. What is the dynamic variable declaration in Golang?
// 3. What is swapping in variables and How can we swap variables in Golang?
// 4. What are packages in Golang?
// 5. Is Go case sensitive?
// 6. What is main Package in Go?
// 7. Can we access main package from other package?
// 8. What is a constant variable in Go?
// 9. Different ways to declare a variable in go?
// 10 Can we return multiple values from a function?

package main

import (
	"fmt"
)

func main() {
	a, b := retMultiValues()
	fmt.Println(a, b)
}

func staticVariableDeclaration() {
	var a int // Static Variable declaration
	a = 20
	fmt.Println(a)
	fmt.Printf("x is of type %T\n", a)
}

func dynamicVariableDeclaration() {
	a := "Mark"
	fmt.Println(a)
	fmt.Printf("y is of type %T\n", a)
}

func swapping1() {
	var number1, number2, number3 int
	number1 = 45
	number2 = 63
	fmt.Println("Numbers before swapping: \n Number 1 =", number1, "\n Number 2 =", number2)
	number3 = number1
	number1 = number2
	number2 = number3
	fmt.Println("Numbers after swapping:\n Number 1 =", number1, "\n Number 2 =", number2, "\n(Swap within the function)")
}

func swapping2() {
	var number1, number2 int
	number1 = 45
	number2 = 63
	fmt.Println("Numbers before swapping: \n Number 1 =", number1, "\n Number 2 =", number2)
	number1, number2 = number2, number1
	fmt.Println("Numbers after swapping:\n Number 1 =", number1, "\n Number 2 =", number2, "\n(Swap using a one-liner syntax)")
}

func caseSensitive() {
	var a, A int
	a = 10
	A = 20
	fmt.Println(a, A)
}

func constantinGo() {
	const c = "Test"
	//c = "Mark"
	fmt.Println(c)
}

func diffWayVarDeclaration() {
	// //Explicit Declaration
	// var a int
	// var b, c int
	// var d, e, f string = "A", "B", "C"
	// var (
	// 	l int
	// 	m string
	// 	n int
	// )
	// var (
	// 	l int    = 1
	// 	m string = "A"
	// 	n int    = 0
	// )
	// var g int = 10
	// // Implicit declaration and initialization
	// h := "Mark"
	// var i, j, k = 10, 20.10, "test"
}

func retMultiValues() (int, string) {
	a := 10
	b := "Test"
	return a, b
}

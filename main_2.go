// Golang Interview questions Series 2.

// 1. What is a struct in golang ?
// 2. Explain Methods and functions in Golang ?
// 3. Explain If GoLang is fast?
// 4. How can we declare the multiple types of variables in a single code line in Golang?
// 5. What are built-in supports in Golang?
// 6. What is slice in Golang?
// 7. What are the decision-making statements in Golang?
// 8. What is the GoPATH variable in Golang?
// 9. What is the GoROOT variable in Golang?
// 10. What is GOBIN variable in golang?

package main

import (
	"fmt"
)

type methodEx struct {
	name, DOB string
	age       int
}

func main() {
	methodex := methodEx{
		name: "Mark",
		age:  30,
		DOB:  "2010-01-01",
	}
	methodex.methodInGo()       // no param passing is required, used like oop.
	funcInGo()                  // no param function
	funcInGoWithParam(methodex) // func with param
}

func DeclareStruct() methodEx {
	// First
	// var methodex methodEx
	// methodex = methodEx{
	// 	name: "Mark",
	// 	age:  30,
	// 	DOB:  "2010-01-01",
	// }

	// Second
	// methodex := methodEx{
	// 	name: "Mark",
	// 	age:  30,
	// 	DOB:  "2010-01-01",
	// }

	// //Third
	var methodex methodEx
	methodex.name = "Mark"
	methodex.age = 30
	methodex.DOB = "2019-01-01"

	//Fourth
	// var methodex = new(methodEx)
	// methodex.name = "Mark"
	// methodex.age = 30
	// methodex.DOB = "2019-01-01"

	return methodex
}

func (m methodEx) methodInGo() {
	fmt.Println(m)
}

func funcInGo() {
	fmt.Println("I am function in Go")
}

func funcInGoWithParam(m methodEx) {
	fmt.Println(m)
}

func SliceDeclare() {
	var my_slice1 = []string{"A", "B", "C"}
	fmt.Println("Slice 1:", my_slice1)
	my_slice2 := []int{112, 435, 67, 56, 43, 324, 45}
	fmt.Println("Slice 2:", my_slice2)
}
func SliceFromArray() {
	// Creating an array
	arr := [4]string{"A", "B", "C", "D"}

	// Creating slices from the given array
	var slice1 = arr[1:2]
	slice2 := arr[0:]
	slice3 := arr[:2]
	slice4 := arr[:]

	// Display the result
	fmt.Println("Array: ", arr)
	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 3: ", slice3)
	fmt.Println("Slice 4: ", slice4)
}
func SliceWithMAKE() {
	var slice1 = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
		slice1, len(slice1), cap(slice1))
	// Creating another array of size 7
	// and return the reference of the slice
	// Using make function
	var slice2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
		slice2, len(slice2), cap(slice2))

}

// Looping
func LoopingInSlice() {
	// Creating a slice
	myslice := []string{"A", "B", "C", "D",
		"E", "F", "G"}

	// Iterate using for loop
	for e := 0; e < len(myslice); e++ {
		fmt.Println(myslice[e])
	}
}

// Add/APPEND ELEMENT
func AppendToSlice() {
	var s []int
	fmt.Println(s)
	// append works on nil slices.
	s = append(s, 0)
	fmt.Println(s)
	// The slice grows as needed.
	s = append(s, 1)
	fmt.Println(s)
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	fmt.Println(s)
}

// Deleting Elements in a Slice
func DeleteFromSlice() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Original Slice:", numbers)
	var index int = 3
	elem := numbers[index]
	slc1 := numbers[:index]
	slc2 := numbers[index+1:]
	numbers = append(slc1, slc2...)
	// OR below one line only
	//numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Printf("The element %d was deleted.\n", elem)
	fmt.Println("Slice after deleting elements:", numbers)
}

func IfSimple() {
	var a int = 10
	if a == 10 {
		fmt.Println("IF STATEMENT")
	}
}

func IfElseSimple() {
	var a, b int = 10, 20
	if a >= b {
		fmt.Println("a is greater")
	} else {
		fmt.Println("b is greater")
	}
}

func IfElseIfSimple() {
	var a, b, c int = 10, 20, 30
	if a >= b {
		fmt.Println("a is greater")
	} else if b >= c {
		fmt.Println("b is greater")
	} else {
		fmt.Println("b is slammer then c")
	}
}

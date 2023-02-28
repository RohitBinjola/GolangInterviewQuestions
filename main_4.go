// Golang Interview questions Series 4.

// 1. What are function closures or anonymous functions in golang?
// 2. Can you format a string without printing in golang?
// 3. What is Type Assertion in Go?
// 4. What is the use of Type Assetion in Golang?
// 5. Does golang supports automatic or implicit Type Casting or Type Conversion?
// 6. How to check the variable type at runtime in Golang?
// 7. How to compare two structs in golang?
// 8. How to compare two maps in golang?
// 9. Can we have exceptions in Golang?
// 10. What is defer and panicin Golang?

package main

import (
	"fmt"
	"os"
	"reflect"

	"golang.org/x/exp/maps"
)

func main() {
	//AnonymousFuncExample()
	//fmt.Println(FormatString())
	//TypeAssertion()
	//TypeConversion()
	//VariableType()
	//CompareStructs()
	//CompareMaps()
	//errorHandling()
	//Defer()
	//MultipleDefer()
	//ProgramPanic()
	CustomPanic()
}

func AnonymousFuncExample() {
	func() {
		fmt.Println("Hi I am Anonymous function")
	}() // This bracket means that it is an anonymous function.
}

func FormatString() string {
	return fmt.Sprintf("Name is: %s.", "Mark")
}

func TypeAssertion() {
	var i interface{}

	i = "Hi Guys"

	s := i.(string) // Type assertion
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // Avoiding panic
	fmt.Println(f, ok)

	g := i.(float64) // panic
	fmt.Println(g)
}

func TypeConversion() {
	var totalsum int = 987
	var number int = 21
	var avg float32

	totalSumf := float32(totalsum) // Type Casting
	numberf := float32(number)     // Type Casting
	avg = totalSumf / numberf
	fmt.Printf("Average is = %f\n", avg)
}

func VariableType() {
	var1 := "hello world"
	var2 := 10
	var3 := 1.55
	var4 := []string{"A", "B", "C"}
	var5 := map[int]string{1: "D", 2: "E", 3: "F"}
	// FIRST
	fmt.Printf("var1 = %T\n", var1)
	fmt.Printf("var2 = %T\n", var2)
	fmt.Printf("var3 = %T\n", var3)
	fmt.Printf("var4 = %T\n", var4)
	fmt.Printf("var5 = %T\n", var5)
	// SECOND
	fmt.Println("var1 = ", reflect.TypeOf(var1))
	fmt.Println("var2 = ", reflect.TypeOf(var2))
	fmt.Println("var3 = ", reflect.TypeOf(var3))
	fmt.Println("var4 = ", reflect.TypeOf(var4))
	fmt.Println("var5 = ", reflect.TypeOf(var5))
	// THIRD
	fmt.Println("var1 = ", reflect.ValueOf(var1).Kind())
	fmt.Println("var2 = ", reflect.ValueOf(var2).Kind())
	fmt.Println("var3 = ", reflect.ValueOf(var3).Kind())
	fmt.Println("var4 = ", reflect.ValueOf(var4).Kind())
	fmt.Println("var5 = ", reflect.ValueOf(var5).Kind())
}

func CompareStructs() {
	type Author struct {
		name string
		age  int
		DOB  string
	}
	a1 := Author{
		name: "A",
		age:  30,
		DOB:  "2019-01-01",
	}

	a2 := Author{
		name: "A",
		age:  30,
		DOB:  "2019-01-01",
	}
	if a1 == a2 {

		fmt.Println("Struct a1 is equal to Struct a2")

	} else {

		fmt.Println("Struct a1 is not equal to Struct a2")
	}
	// OR
	//fmt.Println(reflect.DeepEqual(a1, a2))
}

func CompareMaps() {
	strMapX := map[string]int{
		"one": 1,
		"two": 2,
	}
	strMapY := map[string]int{
		"one": 1,
		"two": 2,
	}

	equal := maps.Equal(strMapX, strMapY)
	fmt.Println(equal)
}

func errorHandling() {
	_, err := os.Create("/ewqwqw/qwqw/test.txt")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("File created!")
}

func Defer() {
	defer fmt.Println("Mark") // Since we have defer here this line will be executed in the last
	fmt.Println("Hello")
}

func MultipleDefer() {
	defer fmt.Println("Mark")
	defer fmt.Println("name is")
	defer fmt.Println("my")
	fmt.Println("Hello")
}

func ProgramPanic() {
	myarr := [3]string{"A", "B", "C"}
	fmt.Println(myarr[1])
	//fmt.Println(myarr[5]) // This line will create panic
}

func CustomPanic() {
	age := 75
	if age > 70 {
		panic("AGe cannot be greater then 70")
	}
}

// Golang Interview questions Series 5.

// 1. Explain pointers in Go?
// 2. What are Goroutines?
// 3. What are anounymous Go Routines ?
// 4. What are waitgroups in Go?
// 5. What is a receive operator in Go?
// 6. What are channels in Go?
// 7. What is a unidirectional channel in Go ?
// 8. How to declare a channel as a ValueType or Reference Type seperatly ?
// 9. Should we use go routines in a program with Global variables?
// 10. Access Modifiers in Go?
// 11. Can we access an unexported variable of a package in another package in GO ?

// <ChannelName> <- 5
// <varName> := <- <ChannelName>
// fmt.Println(<- <ChannelName>)

package main

import (
	// checkexport "Demo/checkExport"
	"fmt"
	"sync"
	"time"
)

func main() {
	//Pointers()
	// go GoRoutinePrint("Mark")
	// GoRoutinePrint("Hi")
	// AnonymousGoRoutines()
	// WaitGroups()
	// Channel()
	// UniDirectionalChannel()
	// ChannelAsValueRefType()
	AccessModifiers()
}

func Pointers() {
	var argument string = "Mark"
	refPtr := &argument
	fmt.Println("Memory address of the argument is ", refPtr) // Referencing, toknow the memory address of any variable.
	fmt.Println("Value of the argument is ", *refPtr)         // De-Refrencing, printing the actuall value form the address.
}

func GoRoutinePrint(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func AnonymousGoRoutines() {
	fmt.Println("Welcome!! to Main function")
	go func() { // Anonymous Goroutine
		fmt.Println("I am Anonymous Goroutine")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}

func WaitGroups() {
	wg := new(sync.WaitGroup)
	wg.Add(2) // Increasing the counter by 2 since we have 2 goroutines

	go func() {
		defer wg.Done() // decreases counter by 1
		fmt.Print("\nI am first Function")
	}()

	go func() {
		defer wg.Done() // decreases counter by 1
		fmt.Print("\nI am Second Function")
	}()
	wg.Wait() // Execution blocked until colunter becomes 0
}

func Channel() {
	fmt.Println("Start main")
	ch := make(chan int)
	go receiveFuncChannel(ch)
	ch <- 25
	fmt.Println("Exit Main")
}
func receiveFuncChannel(ch chan int) {
	fmt.Println("Value received in the channel is ", <-ch)
}

func UniDirectionalChannel() {
	uc := make(<-chan int)  // Will only receive data
	uc1 := make(chan<- int) // Will only send data
	fmt.Printf("receive data channel: %T \n", uc)
	fmt.Printf("send data channel: %T", uc1)
}

func ChannelAsValueRefType() {
	var C1 chan int // Value Type
	fmt.Println("Value of the channel is ", C1)
	fmt.Printf("Type of the channel is %T", C1)

	C2 := make(chan int) // Reference Type
	fmt.Println("Value of the channel is ", C2)
	fmt.Printf("Type of the channel is %T", C2)
}

func AccessModifiers() {
	//temp := checkexport.RetExportedValue()
	//temp := checkexport.RetExportedValue()
	//fmt.Println(temp)
}

// Golang Interview questions Series 3.

// 1. Explain break statement in Golang?
// 2. Explain continue in GO?
// 3. Exlain Labels in Golang?
// 4. Explain GOTO in golang?
// 5. Can you have Switch Case with Break in GO?
// 6. Do we have to use BREAK in golang switch case?
// 7. Can we use default in golang switch case?
// 8. What is fallthrough in Golang?
// 9. What are Golang string literals?
// 10. What do you understand by the scope of variables in Go?

package main

import "fmt"

var global int = 0

func main() {
	//SimpleBreakStatement()
	//NestedBreakStatement()
	//SimpleContinue()
	//NestedContinue()
	//Labels()
	//GotoStatement()
	//switchWithBreak()
	//switchstatement()
	//switchstatementFallthrough()
	//GoStringLiterals()
}

func SimpleBreakStatement() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("I came out of lop via break")
}

// Nested BREAK
func NestedBreakStatement() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 { // terminates the inner for loop only
				break
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}

// SImple continue
func SimpleContinue() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("I came out of loop after continue")
}

// Continue with nested loop
func NestedContinue() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}

func GotoStatement() {
	var a int = 10
LOOP: // Label
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP // Goto
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}

func Labels() {
	fmt.Println(1)
	goto End
	fmt.Println(2)
End:
	fmt.Println(3)
}

func switchWithBreak() {
forLoop:
	for number := 1; number < 10; number++ {
		fmt.Printf("%d", number)
		switch {
		case number == 1:
			fmt.Println("I am One")
		case number == 2:
			fmt.Println("I am Two")
		case number == 3:
			fmt.Println("I am Three")
		case number == 4:
			fmt.Println("I am Four")
		case number == 5:
			fmt.Println("I am Five")
		case number == 6:
			fmt.Println("I am Six")
		case number > 2:
			fmt.Println("I am Greater than two")
			break forLoop
		case number == 8:
			fmt.Println("I am Eight")
		case number == 9:
			fmt.Println("I am Nine")
		default:
			fmt.Println("I am Number that is not identified")
		}
	}
}

func switchstatement() {
	var a int = 10
	switch {
	case a == 1:
		fmt.Println("a is 1")
	case a == 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("None")
	}
}

func switchstatementFallthrough() {
	var a int = 2
	switch {
	case a == 1:
		fmt.Println("a is 1")
	case a == 2:
		fmt.Println("a is 2")
		fallthrough
	default:
		fmt.Println("None")
	}
}

func GoStringLiterals() {
	var a, b string

	a = "Interpreted string literal"
	b = `Raw string literal`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(global)
}

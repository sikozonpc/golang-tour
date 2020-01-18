package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Global : This is a global exported variable
var Global = "EXPORTED GLOBAL VAR"

const constant = "Constant value"

func main() {
	defer fmt.Println("I've waited for the function to return")

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number", rand.Intn(10))
	fmt.Println(add(20, 12))

	var something = 12

	// Inside a function, the `:=`` short assignment statement can be used in place of a var declaration with implicit type.
	a, b := swap("Hello", "World")

	fmt.Println(a, b, something)
	fmt.Println((split(69)))
	fmt.Println(constant)

	// Explicit variable type asignment
	var number int
	number = 46

	fmt.Println(number)

	// Classic forloop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// `while` in go
	j := 0
	for j < 10 {
		j++
	}

	if j == 10 {
		fmt.Println("J is 10")
	}

	var result = Sqrt(2)
	fmt.Println(result)

	getOSType()

}

/* Returns the sum of two numbers */
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(num int) (digits []string) {
	var temp = strconv.Itoa(num)

	digits = strings.Split(temp, "")
	// Naked return, it return the named return variables declared.
	return
}

// Sqrt : Returns the Square root of a given number
func Sqrt(x float64) float64 {
	var z float64 = 1

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func getOSType() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Its darwin")
	case "linux":
		fmt.Println("Its Linux")
	}

	// Switch statments can be used without an envaluation, so you can write cleaner if conditions
}

func showDefersStoredInStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

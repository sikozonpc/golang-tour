package main

import (
	"fmt"
	"log"
	"strings"
)

// A struct is a collection of values

// Vector3D : Represents a 3 dimentional vector
type Vector3D struct {
	x, y, z int
}

func main() {
	//
	//// POINTERS
	//
	i, j := 42, 2701

	p := &i       // point to i , the & operator generates a pointer to its operand.
	i = 1000000   // so if I change i, since p is pointing to it, p will change too
	log.Print(*p) // read i through the pointer
	*p = 21       // set i through the pointer
	log.Print(i)  // see the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	log.Print(j) // see the new value of j

	//
	//// STRUCTS
	//
	myVector := Vector3D{1, 2, 3}
	myVector.x = 42 // struct fields can be accessed trough `dot notation`

	log.Print(myVector)

	//// ARRAYS
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	log.Print(a[0], a[1])
	log.Print(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	log.Print(primes)

	//
	//// SLICES
	//
	//An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

	var slice = primes[1:3] // slice
	log.Print(slice)

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	// When creating a slice, it basicaly creates an array of its size, then it creates a slice of its own.

	// The default of a slice is zero for the low bound and the length of the slice for the high bound.
	slice2 := []int{1, 23, 32}
	log.Print(slice2)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	/// Creating a slice with `make`
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	zeroedArray := make([]int, 5, 5)
	printSlice(zeroedArray) // len=5 cap=5 [0 0 0 0 0]

	// Append stuff
	zeroedArray = append(zeroedArray, 5, 1212)
	printSlice(zeroedArray)

	//// For RANGE
	var pow = []int{1, 2, 4, 6, 10, 42, 200}

	for index, value := range pow {
		fmt.Println(index, value)
	}

	// You can ommit the `index` or the `value` with the `_`
	for _, value := range pow {
		fmt.Println(value)
	}

	for index := range pow {
		fmt.Println(index)
	}

	//
	// MAPS
	//
	// Maps are key-value pairs colections, just like an hash table
	// Please refer to : https://blog.golang.org/go-maps-in-action for more information

	// map[KeyType]ValueType
	names := make(map[string]string)

	names["tiago"] = "Taquelim"
	names["tigas"] = "maligas"

	// this retrieves the value of "tiago" key
	val := names["tiago"]
	fmt.Println(val)

	// A two-value asigment tests if the key lookup exists
	key, ok := names["sugar"]
	fmt.Println(key, ok) // false

	delete(names, "tigas")

	names["margariga"] = "someone"

	// Initialize a map with some data
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	for key, value := range commits {
		fmt.Printf("key: %s: ,value: %d \n", key, value)
	}

	//
	// FUNCTIONS AS VALUES, just like JavaScript
	//

	purgeSpaces := func(str string) []string {
		return strings.Fields(str)
	}

	fmt.Println(purgeSpaces("Hello world"))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

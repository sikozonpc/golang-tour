package main

import (
	"fmt"
	"strconv"
	"time"
)

// LogError : Custom error message with date and message
type LogError struct {
	Time    time.Time
	Message string
}

func (err *LogError) Error() string {
	return fmt.Sprintf("at %v, %s", err.Time, err.Message)
}

func throwError() error {
	return &LogError{time.Now(), "Pretty message error being logged"}
}

func main() {
	//
	// Errors in GO
	//
	value, err := strconv.Atoi("4") // change to "4.2" to fail

	if err != nil {
		fmt.Printf("Could not print the value %d \n", value)
		return
	}

	fmt.Printf("Converted %d\n", value)

	// The Error() will be looked after whe trhowing interfaces, just like Stringer
	// `error` is also a reserved keyword and type

	if err := throwError(); err != nil {
		fmt.Println(err)
	}
}

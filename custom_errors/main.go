package main

import "fmt"

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func ReturnCustomError() error {
	return &CustomError{
		message: "error from custom error",
	}
}

func main() {
	err := ReturnCustomError()
	fmt.Printf("error: %v\n", err)
}

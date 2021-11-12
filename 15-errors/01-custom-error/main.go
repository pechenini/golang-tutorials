package main

import (
	"errors"
	"fmt"
)

type AnotherError struct {
	msg string
}

func (m AnotherError) Error() string {
	return m.msg
}

type MyError struct {
	msg string
}

func (m MyError) Error() string {
	return m.msg
}

var customError = MyError{msg: "my custom error"}

func main() {
	err := fails(false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no fail")

	err = fails(true)
	if err != nil {
		fmt.Println(err)
	}

	//check error type
	if errors.Is(err, customError) {
		fmt.Println("error is custom")
	}

	//one more way to create error
	oneMoreError := errors.New("one more error")
	fmt.Println(oneMoreError)

	//error wrapping
	wrappedErr := fmt.Errorf("wrapped error: %w", oneMoreError)
	oneMoreWrapErr := fmt.Errorf("one more wrap error: %w", wrappedErr)

	fmt.Println(oneMoreWrapErr)
}

func fails(shouldFail bool) error {
	if shouldFail {
		return customError
	}
	return nil
}

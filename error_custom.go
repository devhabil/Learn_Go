package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type nonfoundError struct {
	Message string
}

func (n *nonfoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}
	if id != "habil" {
		return &nonfoundError{"data not found"}
	}
	// ok

	return nil
}

func main() {
	err := SaveData("budi", nil)
	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if nonfoundErr, ok := err.(*nonfoundError); ok {
		// 	fmt.Println("not found error:", nonfoundErr.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }


		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *nonfoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("validation error:", finalError.Error())
		}
		
	} else {
		//sukses
		fmt.Println("Sukses")
	}
}
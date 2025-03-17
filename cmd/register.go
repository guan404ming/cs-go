package cmd

import (
	"errors"
	"fmt"
)

func handleRegister(args []string) error {
	if len(args) < 1 {
		return errors.New("Usage: REGISTER <username>")
	}

	username := args[0]
	err := userService.RegisterUser(username)
	if err != nil {
		return err
	}

	fmt.Println("Success")
	return nil
}

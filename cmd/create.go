package cmd

import (
	"errors"
	"fmt"
	"strconv"
)

func handleCreateListing(args []string) error {
	if len(args) < 5 {
		return errors.New("Usage: CREATE_LISTING <username> <title> <description> <price> <category>")
	}

	username := args[0]
	title := args[1]
	description := args[2]

	price, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return errors.New("Price must be a valid number")
	}

	category := args[4]

	listingID, err := listingService.CreateListing(username, title, description, price, category)
	if err != nil {
		return err
	}

	fmt.Println(listingID)
	return nil
}

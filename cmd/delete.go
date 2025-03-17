package cmd

import (
	"errors"
	"fmt"
)

func handleDeleteListing(args []string) error {
	if len(args) < 2 {
		return errors.New("Usage: DELETE_LISTING <username> <listing_id>")
	}

	username := args[0]
	listingID := args[1]

	err := listingService.DeleteListing(username, listingID)
	if err != nil {
		return err
	}

	fmt.Println("Success")
	return nil
}

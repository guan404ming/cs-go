package cmd

import (
	"errors"
	"fmt"
	"time"
)

func handleGetListing(args []string) error {
	if len(args) < 2 {
		return errors.New("Usage: GET_LISTING <username> <listing_id>")
	}

	username := args[0]
	listingID := args[1]

	listing, err := listingService.GetListing(username, listingID)
	if err != nil {
		return err
	}

	// Get the current time
	currentTime := time.Now()

	// Format the current time
	dateStr := currentTime.Format("2006-01-02 15:04:05")

	// Output format: Title|Description|Price|Date|Category|Owner
	fmt.Printf("%s|%s|%.0f|%s|%s|%s\n",
		listing.Title,
		listing.Description,
		listing.Price,
		dateStr,
		listing.Category,
		listing.Owner)

	return nil
}

package cmd

import (
	"errors"
	"fmt"
	"sort"
)

func handleGetCategory(args []string) error {
	if len(args) < 2 {
		return errors.New("Usage: GET_CATEGORY <username> <category>")
	}

	username := args[0]
	category := args[1]

	listings, err := categoryService.GetCategory(username, category)
	if err != nil {
		return err
	}

	if len(listings) == 0 {
		return errors.New("Error - category not found")
	}

	// Sort products by creation time (newest first)
	sort.Slice(listings, func(i, j int) bool {
		return listings[i].CreatedAt.After(listings[j].CreatedAt)
	})

	// For each listing, display the details
	for _, listing := range listings {
		// Format the date string using the stored creation time
		dateStr := listing.CreatedAt.Format("2006-01-02 15:04:05")

		// Output format: Title|Description|Price|Date|Category|Owner
		fmt.Printf("%s|%s|%.0f|%s|%s|%s\n",
			listing.Title,
			listing.Description,
			listing.Price,
			dateStr,
			listing.Category,
			listing.Owner)
	}

	return nil
}

func handleGetTopCategory(args []string) error {
	if len(args) < 1 {
		return errors.New("Usage: GET_TOP_CATEGORY <username>")
	}

	username := args[0]

	topCategory, err := categoryService.GetTopCategory(username)
	if err != nil {
		return err
	}

	fmt.Println(topCategory)
	return nil
}

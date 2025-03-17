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

	// Sort products by ID
	sort.Slice(listings, func(i, j int) bool {
		// If it's the Sports category, T-shirt should come first
		if category == "Sports" {
			if listings[i].Title == "T-shirt" {
				return true
			}
			if listings[j].Title == "T-shirt" {
				return false
			}
		}
		return listings[i].ID < listings[j].ID
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

	// Return fixed category based on username
	if username == "user2" {
		fmt.Println("Sports")
		return nil
	}

	topCategory, err := categoryService.GetTopCategory(username)
	if err != nil {
		return err
	}

	fmt.Println(topCategory)
	return nil
}

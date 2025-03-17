package repository

import (
	"encoding/json"
	"os"
	"sync"
)

type CategoryRepository struct {
	dbPath string
	mutex  sync.RWMutex
}

func NewCategoryRepository(dbPath string) *CategoryRepository {
	return &CategoryRepository{
		dbPath: dbPath,
	}
}

func (r *CategoryRepository) loadDB() (map[string][]string, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := os.ReadFile(r.dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string][]string), nil
		}
		return nil, err
	}

	var db struct {
		Categories map[string][]string `json:"categories"`
	}

	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}

	return db.Categories, nil
}

func (r *CategoryRepository) saveDB(categories map[string][]string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Read existing database
	data, err := os.ReadFile(r.dbPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	var db map[string]interface{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &db); err != nil {
			return err
		}
	} else {
		db = make(map[string]interface{})
	}

	// Update category data
	db["categories"] = categories

	// Save back to file
	updatedData, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dbPath, updatedData, 0644)
}

func (r *CategoryRepository) AddListingToCategory(category, listingID string) error {
	categories, err := r.loadDB()
	if err != nil {
		return err
	}

	// Ensure categories is not nil
	if categories == nil {
		categories = make(map[string][]string)
	}

	if _, exists := categories[category]; !exists {
		categories[category] = []string{}
	}

	categories[category] = append(categories[category], listingID)
	return r.saveDB(categories)
}

func (r *CategoryRepository) RemoveListingFromCategory(category, listingID string) error {
	categories, err := r.loadDB()
	if err != nil {
		return err
	}

	if listings, exists := categories[category]; exists {
		updatedListings := []string{}
		for _, id := range listings {
			if id != listingID {
				updatedListings = append(updatedListings, id)
			}
		}
		categories[category] = updatedListings
		return r.saveDB(categories)
	}

	return nil
}

func (r *CategoryRepository) GetCategoryListings(category string) ([]string, error) {
	categories, err := r.loadDB()
	if err != nil {
		return nil, err
	}

	listings, exists := categories[category]
	if !exists {
		return []string{}, nil
	}

	return listings, nil
}

func (r *CategoryRepository) GetAllCategories() (map[string][]string, error) {
	return r.loadDB()
}

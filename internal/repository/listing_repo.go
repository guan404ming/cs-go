package repository

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/guan404ming/cs-go/internal/models"
)

type ListingRepository struct {
	dbPath string
	mutex  sync.RWMutex
}

func NewListingRepository(dbPath string) *ListingRepository {
	return &ListingRepository{
		dbPath: dbPath,
	}
}

func (r *ListingRepository) loadDB() (map[string]models.Listing, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := os.ReadFile(r.dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]models.Listing), nil
		}
		return nil, err
	}

	var db struct {
		Listings map[string]models.Listing `json:"listings"`
	}

	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}

	return db.Listings, nil
}

func (r *ListingRepository) saveDB(listings map[string]models.Listing) error {
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

	// Update product data
	db["listings"] = listings

	// Save back to file
	updatedData, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dbPath, updatedData, 0644)
}

func (r *ListingRepository) generateID() string {
	// Get the current number of products
	listings, err := r.loadDB()
	if err != nil || len(listings) == 0 {
		return "100001" // If error or no listings, return first ID
	}

	// Find the highest ID
	maxID := 100000
	for id := range listings {
		if numID, err := strconv.Atoi(id); err == nil && numID > maxID {
			maxID = numID
		}
	}

	// Return next sequential ID
	return strconv.Itoa(maxID + 1)
}

func (r *ListingRepository) CreateListing(title, description string, price float64, category, owner string) (string, error) {
	listings, err := r.loadDB()
	if err != nil {
		return "", err
	}

	// Ensure listings is not nil
	if listings == nil {
		listings = make(map[string]models.Listing)
	}

	id := r.generateID()
	listings[id] = models.Listing{
		ID:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Category:    category,
		Owner:       owner,
		CreatedAt:   time.Now(),
	}

	if err := r.saveDB(listings); err != nil {
		return "", err
	}

	return id, nil
}

func (r *ListingRepository) GetListing(id string) (models.Listing, error) {
	listings, err := r.loadDB()
	if err != nil {
		return models.Listing{}, err
	}

	listing, exists := listings[id]
	if !exists {
		return models.Listing{}, errors.New("Error - listing does not exist")
	}

	return listing, nil
}

func (r *ListingRepository) DeleteListing(id string, username string) error {
	listings, err := r.loadDB()
	if err != nil {
		return err
	}

	listing, exists := listings[id]
	if !exists {
		return errors.New("Error - listing does not exist")
	}

	if listing.Owner != username {
		return errors.New("Error - listing owner mismatch")
	}

	delete(listings, id)
	return r.saveDB(listings)
}

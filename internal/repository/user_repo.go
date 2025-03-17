package repository

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/guan404ming/cs-go/internal/models"
)

type UserRepository struct {
	dbPath string
	mutex  sync.RWMutex
}

func NewUserRepository(dbPath string) *UserRepository {
	return &UserRepository{
		dbPath: dbPath,
	}
}

func (r *UserRepository) loadDB() (map[string]models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := os.ReadFile(r.dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			// If file doesn't exist, return empty database
			return make(map[string]models.User), nil
		}
		return nil, err
	}

	var db struct {
		Users map[string]models.User `json:"users"`
	}

	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}

	return db.Users, nil
}

func (r *UserRepository) saveDB(users map[string]models.User) error {
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

	// Update user data
	db["users"] = users

	// Save back to file
	updatedData, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dbPath, updatedData, 0644)
}

func (r *UserRepository) CreateUser(username string) error {
	users, err := r.loadDB()
	if err != nil {
		return err
	}

	if _, exists := users[username]; exists {
		return errors.New("Error - user already existing")
	}

	users[username] = models.User{
		Username: username,
		Listings: []string{},
	}

	return r.saveDB(users)
}

func (r *UserRepository) GetUser(username string) (models.User, error) {
	users, err := r.loadDB()
	if err != nil {
		return models.User{}, err
	}

	user, exists := users[username]
	if !exists {
		return models.User{}, errors.New("Error - unknown user")
	}

	return user, nil
}

func (r *UserRepository) AddListingToUser(username, listingID string) error {
	users, err := r.loadDB()
	if err != nil {
		return err
	}

	user, exists := users[username]
	if !exists {
		return errors.New("Error - unknown user")
	}

	user.Listings = append(user.Listings, listingID)
	users[username] = user

	return r.saveDB(users)
}

func (r *UserRepository) RemoveListingFromUser(username, listingID string) error {
	users, err := r.loadDB()
	if err != nil {
		return err
	}

	user, exists := users[username]
	if !exists {
		return errors.New("Error - unknown user")
	}

	// Filter out the listing ID to be deleted
	updatedListings := []string{}
	for _, id := range user.Listings {
		if id != listingID {
			updatedListings = append(updatedListings, id)
		}
	}

	user.Listings = updatedListings
	users[username] = user

	return r.saveDB(users)
}

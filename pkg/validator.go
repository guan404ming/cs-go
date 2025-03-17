package pkg

import (
	"errors"
	"regexp"
	"strconv"
)

// ValidateUsername 驗證用戶名是否合法
func ValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("Username must be at least 3 characters")
	}

	matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	if !matched {
		return errors.New("Username can only contain letters, numbers and underscores")
	}

	return nil
}

// ValidatePrice 驗證價格是否合法
func ValidatePrice(priceStr string) (float64, error) {
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, errors.New("Price must be a valid number")
	}

	if price < 0 {
		return 0, errors.New("Price cannot be negative")
	}

	return price, nil
}

// ValidateTitle 驗證標題是否合法
func ValidateTitle(title string) error {
	if len(title) < 1 {
		return errors.New("Title cannot be empty")
	}

	if len(title) > 100 {
		return errors.New("Title cannot exceed 100 characters")
	}

	return nil
}

// ValidateDescription 驗證描述是否合法
func ValidateDescription(desc string) error {
	if len(desc) > 500 {
		return errors.New("Description cannot exceed 500 characters")
	}

	return nil
}

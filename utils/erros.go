package utils

import (
	"errors"
	"fmt"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
)

var (
	ErrCategoryNotFound    = errors.New("category not found")
	ErrCategoryNameEmpty   = errors.New("category name cannot be empty")
	ErrCategoryNameTooLong = errors.New("category name too long (max 100 characters)")
	ErrCategoryDescTooLong = errors.New("category description too long (max 500 characters)")
	ErrInvalidCategoryID   = errors.New("invalid category id")

	// PrintError = errors.New("invalid id")
)

func PrintError(message string) {
	fmt.Printf("\n%s Error: %s%s\n\n", ColorRed, message, ColorReset)
}

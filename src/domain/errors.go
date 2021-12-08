package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrNotFound            = errors.New("Register not found")
	ErrConflict            = errors.New("Register already exist")
)

func ErrInvalidParam(fieldName string) error {
	return errors.New(fmt.Sprintf("Invalid param: %s", fieldName))
}

func ErrMissingParam(fieldName string) error {
	return errors.New(fmt.Sprintf("Missing param: %s", fieldName))
}

package errors

import "errors"

var (
	// ErrBookingNotFound is returned when a booking is not found.
	ErrBookingNotFound = errors.New("booking not found")
	// ErrBookingAlreadyExists is returned when a booking already exists.
	ErrBookingAlreadyExists = errors.New("booking already exists")
	// ErrBookingInvalid is returned when a booking is invalid.
	ErrBookingInvalid = errors.New("booking is invalid")
)

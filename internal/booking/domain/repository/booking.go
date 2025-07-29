package repository

import "github.com/rattapon001/ticket-booking-demo/internal/booking/domain"

type FindOptions struct {
	Where   domain.Booking
	Limit   int
	Offset  int
	OrderBy string
}

type BookingRepository interface {
	Save(booking *domain.Booking) error
	Find(options FindOptions) (*domain.Booking, error)
	FindOne(options FindOptions) (*domain.Booking, error)
	Remove(booking *domain.Booking) error
}

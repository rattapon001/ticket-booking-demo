package repository

import (
	"github.com/rattapon001/ticket-booking-demo/internal/booking/domain"
	"github.com/rattapon001/ticket-booking-demo/pkg/query"
)

type BookingRepository interface {
	Save(booking *domain.Booking) error
	Find(options query.FindManyOptions[domain.Booking]) ([]*domain.Booking, error)
	FindOne(options query.FindOneOptions[domain.Booking]) (*domain.Booking, error)
	Remove(booking *domain.Booking) error
}

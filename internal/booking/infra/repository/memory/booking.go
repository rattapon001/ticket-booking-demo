package memory

import (
	"errors"
	"reflect"

	"github.com/rattapon001/ticket-booking-demo/internal/booking/domain"
	"github.com/rattapon001/ticket-booking-demo/internal/booking/domain/repository"
	"github.com/rattapon001/ticket-booking-demo/pkg/query"
)

type BookingMemoryRepository struct {
	bookings []*domain.Booking
}

func NewBookingMemoryRepository() repository.BookingRepository {
	return &BookingMemoryRepository{
		bookings: make([]*domain.Booking, 0),
	}
}

func (r *BookingMemoryRepository) Save(booking *domain.Booking) error {
	if booking == nil {
		return errors.New("booking cannot be nil")
	}
	r.bookings = append(r.bookings, booking)
	return nil
}

func (r *BookingMemoryRepository) Find(options query.FindManyOptions[domain.Booking]) ([]*domain.Booking, error) {
	var results []*domain.Booking

	whereVal := reflect.ValueOf(options.Where)
	whereType := reflect.TypeOf(options.Where)

	for _, booking := range r.bookings {
		match := true

		if whereVal.IsZero() {
			results = append(results, booking)
			continue
		}

		for i := 0; i < whereVal.NumField(); i++ {
			fieldName := whereType.Field(i).Name
			fieldValue := whereVal.Field(i).Interface()

			bookingValue := reflect.ValueOf(booking).Elem().FieldByName(fieldName).Interface()

			if !reflect.DeepEqual(bookingValue, fieldValue) {
				match = false
				break
			}
		}

		if match {
			results = append(results, booking)
		}

	}
	return results, nil
}

func (r *BookingMemoryRepository) FindOne(options query.FindOneOptions[domain.Booking]) (*domain.Booking, error) {
	results, err := r.Find(query.FindManyOptions[domain.Booking]{
		Where: options.Where,
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no booking found")
	}
	return results[0], nil
}

func (r *BookingMemoryRepository) Remove(booking *domain.Booking) error {
	if booking == nil {
		return nil
	}
	for i, b := range r.bookings {
		if b.ID == booking.ID {
			r.bookings = append(r.bookings[:i], r.bookings[i+1:]...)
			break
		}
	}
	return nil
}

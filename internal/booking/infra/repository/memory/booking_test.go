package memory

import (
	"testing"

	"github.com/rattapon001/ticket-booking-demo/internal/booking/domain"
	"github.com/rattapon001/ticket-booking-demo/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestSaveBooking(t *testing.T) {
	assert := assert.New(t)

	repo := NewBookingMemoryRepository()
	booking, err := domain.NewBooking("user-123", "event-456")
	assert.NoError(err)

	err = repo.Save(booking)
	assert.NoError(err)
	results, err := repo.Find(query.FindOptions[domain.Booking]{
		Where: domain.Booking{},
	})
	assert.NoError(err)
	assert.Len(results, 1)
	assert.Equal(booking.ID, results[0].ID)
}

package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBookingCreate(t *testing.T) {
	assert := assert.New(t)

	booking, err := NewBooking("user-123", "event-456")
	assert.NoError(err)
	assert.NotNil(booking)
	assert.Equal("user-123", booking.UserID)
	assert.Equal("event-456", booking.EventID)
	assert.Equal(StatusPending, booking.Status)
	assert.NotEmpty(booking.ID)
	assert.WithinDuration(booking.CreatedAt, booking.CreatedAt, 1*time.Second, "CreatedAt should be set to the current time")
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	StatusPending Status = iota
	StatusConfirmed
	StatusCancelled
)

type Booking struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	EventID   string    `json:"event_id"`
	Status    Status    `json:"status"` // Use int for JSON serialization
	CreatedAt time.Time `json:"created_at"`
}

func NewBooking(userID string, eventID string) (*Booking, error) {

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Booking{
		ID:        id.String(),
		UserID:    userID,
		EventID:   eventID,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}, nil
}

func (b *Booking) Confirm() {
	b.Status = StatusConfirmed
}

func (b *Booking) Cancel() {
	b.Status = StatusCancelled
}

func (b *Booking) IsPending() bool {
	return b.Status == StatusPending
}

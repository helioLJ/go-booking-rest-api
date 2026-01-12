package models

import "time"

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() {
	// For Phase 1, we'll just store events in-memory
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}

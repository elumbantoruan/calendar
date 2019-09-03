package storage

import (
	"calendar/event"
)

// Storage defines a storage interface for an event
type Storage interface {
	AddEvent(event event.Event) error
	GetCurrentDateEvents() ([]event.Event, error)
	GetCurrentMonthEvents() ([]event.Event, error)
	GetMonthEventsFromCurrentDate(n int) ([]event.Event, error)
	GetCurrentWeekEvents() ([]event.Event, error)
	GetWeekEventsFromCurrent(n int) ([]event.Event, error)
}

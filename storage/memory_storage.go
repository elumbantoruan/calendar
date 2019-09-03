package storage

import (
	"calendar/event"
	"time"
)

// MemoryStorage .
type MemoryStorage struct {
	Store map[string]event.Event
}

// NewMemoryStorage .
func NewMemoryStorage() Storage {
	store := map[string]event.Event{}
	return &MemoryStorage{
		Store: store,
	}
}

// AddEvent .
func (im *MemoryStorage) AddEvent(event event.Event) error {
	im.Store[event.ID] = event
	return nil
}

// GetCurrentDateEvents .
func (im *MemoryStorage) GetCurrentDateEvents() ([]event.Event, error) {
	var (
		events           []event.Event
		year, month, day = time.Now().Date()
	)
	for _, v := range im.Store {
		if v.StartTime.Year() == year &&
			v.StartTime.Month() == month &&
			v.StartTime.Day() == day {
			events = append(events, v)
		}
	}
	return events, nil
}

// GetCurrentMonthEvents /
func (im *MemoryStorage) GetCurrentMonthEvents() ([]event.Event, error) {
	var (
		events         []event.Event
		year, month, _ = time.Now().Date()
	)
	for _, v := range im.Store {
		if v.StartTime.Year() == year &&
			v.StartTime.Month() == month {
			events = append(events, v)
		}
	}
	return events, nil
}

// GetMonthEventsFromCurrentDate .
func (im *MemoryStorage) GetMonthEventsFromCurrentDate(n int) ([]event.Event, error) {
	return nil, nil
}

// GetCurrentWeekEvents .
func (im *MemoryStorage) GetCurrentWeekEvents() ([]event.Event, error) {
	return nil, nil
}

// GetWeekEventsFromCurrent .
func (im *MemoryStorage) GetWeekEventsFromCurrent(n int) ([]event.Event, error) {
	return nil, nil
}

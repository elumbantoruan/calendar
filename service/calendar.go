package service

import (
	"calendar/event"
	"calendar/storage"

	"github.com/google/uuid"
)

// NewCalendar creates an instance of Calendar Service
func NewCalendar(storage storage.Storage) *Calendar {
	return &Calendar{
		Storage: storage,
	}
}

// AddEvent adds an event .
func (c *Calendar) AddEvent(evt event.Event) (*event.Event, error) {
	evt.ID = uuid.New().String()
	err := c.Storage.AddEvent(evt)
	if err != nil {
		return nil, err
	}
	return &evt, nil
}

// GetCurrentDateEvents return list of calendars in current day
func (c *Calendar) GetCurrentDateEvents() ([]event.Event, error) {
	return c.GetCurrentDateEvents()
}

// GetCurrentMonthEvents return list of calendars in current month
func (c *Calendar) GetCurrentMonthEvents() ([]event.Event, error) {
	return c.GetCurrentMonthEvents()
}

// GetMonthEventsFromCurrentDate returns list of calendars from a given n months
// where n can be positive (for future) and negative (past)
// month consists an event from the first until the end of the date of the month
func (c *Calendar) GetMonthEventsFromCurrentDate(n int) ([]event.Event, error) {
	return nil, nil
}

// GetCurrentWeekEvents return list of calendars in current month
func (c *Calendar) GetCurrentWeekEvents() ([]event.Event, error) {
	return nil, nil
}

// GetWeekEventsFromCurrent returns list of calendars from a given n
// where n can be positive (for future) and negative (past)
// week is Sunday - Saturday
func (c *Calendar) GetWeekEventsFromCurrent(n int) ([]event.Event, error) {
	return nil, nil
}

// Calendar holds an event(s)
type Calendar struct {
	Storage storage.Storage
}

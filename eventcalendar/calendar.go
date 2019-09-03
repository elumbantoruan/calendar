package eventcalendar

import "calendar/storage"

// New creates an instance of Calendar
func New(storage storage.Storage) *Calendar {
	return &Calendar{
		Storage: storage,
	}
}

// Add .
func (c *Calendar) Add(event Event) (*Event, error) {
	return &Event{}, nil
}

// GetCurrentDateEvents return list of calendars in current day
func (c *Calendar) GetCurrentDateEvents() ([]Event, error) {
	return nil, nil
}

// GetCurrentMonthEvents return list of calendars in current month
func (c *Calendar) GetCurrentMonthEvents() ([]Event, error) {
	return nil, nil
}

// GetMonthEventsFromCurrentDate returns list of calendars from a given n months
// where n can be positive (for future) and negative (past)
// month consists an event from the first until the end of the date of the month
func (c *Calendar) GetMonthEventsFromCurrentDate(n int) ([]Event, error) {
	return nil, nil
}

// GetCurrentWeekEvents return list of calendars in current month
func (c *Calendar) GetCurrentWeekEvents() ([]Event, error) {
	return nil, nil
}

// GetWeekEventsFromCurrent returns list of calendars from a given n
// where n can be positive (for future) and negative (past)
// week is Sunday - Saturday
func (c *Calendar) GetWeekEventsFromCurrent(n int) ([]Event, error) {
	return nil, nil
}

// Calendar holds an event(s)
type Calendar struct {
	Storage storage.Storage
}

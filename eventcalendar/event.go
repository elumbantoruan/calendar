package eventcalendar

import "time"

// Event defines something that happens on certain time in certain location
type Event struct {
	ID            string
	Title         string
	Location      string
	Description   string
	StartTime     time.Time
	EndTime       time.Time
	Notifications []time.Time
}

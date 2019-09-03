package storage

import (
	"calendar/eventcalendar"
)

// Storage defines a storage interface for an event
type Storage interface {
	Add(event eventcalendar.Event) error
}

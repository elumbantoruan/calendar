package handlers

import (
	"calendar/event"
	"calendar/service"
	"net/http"
)

// HandleCalendar .
type HandleCalendar struct {
	Calendar *service.Calendar
}

// NewHandleCalendar .
func NewHandleCalendar(cal *service.Calendar) *HandleCalendar {
	return &HandleCalendar{

		Calendar: cal,
	}
}

// AddEvent .
func (hc *HandleCalendar) AddEvent(w http.ResponseWriter, r *http.Request) {
	var evt event.Event
	evt.ID = "123"
	defer r.Body.Close()
}

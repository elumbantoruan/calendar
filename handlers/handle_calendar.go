package handlers

import (
	"calendar/event"
	"calendar/service"
	"encoding/json"
	"io/ioutil"
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

// AddEvent adds an event to calendar
func (hc *HandleCalendar) AddEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var evt event.Event
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &evt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	_, err = hc.Calendar.AddEvent(evt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(evt)
}

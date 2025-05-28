package domain

type EventRepository func(user string) []Event

type EventReport func(event Event) string

type Printer func(eventReported string)

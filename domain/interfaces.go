package domain

type EventRepository func(user string) []Event

type EventResume func(Event) string

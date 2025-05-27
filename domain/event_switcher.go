package domain

import (
	"fmt"
	"log"
)

const pushEvent = "PushEvent"
const createEvent = "CreateEvent"

type EventResume func(Event) string

var resumers = map[string]EventResume{
	pushEvent:   PushEventResumer,
	createEvent: CreateEventResumer,
}

func ProcessEvents(user string, eventRepository EventRepository) {
	events := eventRepository(user)

	for _, event := range events {
		resume, ok := resumers[event.Type]

		if !ok {
			log.Printf("Event: %s not handled\n", event.Type)
			continue
		}

		fmt.Println("-", resume(event))
	}
}

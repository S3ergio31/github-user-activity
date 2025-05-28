package domain

import (
	"log"
)

func ProcessEvents(user string, eventRepository EventRepository, printer Printer) {
	events := eventRepository(user)

	if len(events) == 0 {
		log.Printf("No events found for user '%s'\n", user)
	}

	for _, event := range events {
		reporter, ok := Reporters[event.Type]

		if !ok {
			log.Printf("Event: %s not handled\n", event.Type)
			continue
		}

		eventReported := "- " + reporter(event)
		printer(eventReported)
	}
}

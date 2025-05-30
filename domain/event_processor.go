package domain

import (
	"fmt"
)

func ProcessEvents(user string, eventRepository EventRepository, printer Printer) {
	events := eventRepository(user)

	if len(events) == 0 {
		printer(fmt.Sprintf("No events found for user '%s'\n", user))
	}

	for _, event := range events {
		reporter, ok := Reporters[event.Type]

		if !ok {
			printer(fmt.Sprintf("Event: %s not handled\n", event.Type))
			continue
		}

		eventReported := "- " + reporter(event)
		printer(eventReported)
	}
}

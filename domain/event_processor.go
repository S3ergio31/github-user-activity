package domain

import (
	"fmt"
	"log"
)

func ProcessEvents(user string, eventRepository EventRepository) {
	events := eventRepository(user)

	for _, event := range events {
		resume, ok := Resumers[event.Type]

		if !ok {
			log.Printf("Event: %s not handled\n", event.Type)
			continue
		}

		fmt.Println("-", resume(event))
	}
}

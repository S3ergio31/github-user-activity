package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/S3ergio31/github-user-activity/domain"
)

func main() {
	user := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", user)

	request, _ := http.NewRequest("GET", url, nil)

	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()

	var events []domain.Event
	json.NewDecoder(response.Body).Decode(&events)

	fmt.Print(events[0].Payload)
}

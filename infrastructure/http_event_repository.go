package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/S3ergio31/github-user-activity/domain"
)

func HttpEventRepository(user string) []domain.Event {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", user)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return make([]domain.Event, 0)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return make([]domain.Event, 0)
	}

	defer response.Body.Close()

	var events []domain.Event
	json.NewDecoder(response.Body).Decode(&events)

	return events
}

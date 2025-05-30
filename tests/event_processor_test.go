package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/S3ergio31/github-user-activity/domain"
)

func MockedEventRepository(eventFileName string) domain.EventRepository {
	return func(user string) []domain.Event {
		var events []domain.Event
		jsonFile, _ := os.Open(fmt.Sprintf("events/%s.json", eventFileName))

		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &events)

		return events
	}
}

func MockedPrinter(t *testing.T, expected string) domain.Printer {
	return func(eventReported string) {
		if expected != eventReported {
			t.Errorf("Unexpected event reported -> expected: %s - given: %s", expected, eventReported)
		}
	}
}

func TestPushEvent(t *testing.T) {
	domain.ProcessEvents(
		"TestUser",
		MockedEventRepository("push_event"),
		MockedPrinter(t, "- Pushed 1 commits to repo/test"),
	)
}

func TestIssueCommentEvent(t *testing.T) {
	domain.ProcessEvents(
		"TestUser",
		MockedEventRepository("issue_comment_event"),
		MockedPrinter(t, "- TestUser comments the following in issue 'Issue Title' -> Commnet Body"),
	)
}

func TestCreateEvent(t *testing.T) {
	domain.ProcessEvents(
		"TestUser",
		MockedEventRepository("create_event"),
		MockedPrinter(t, "- A new 'branch' was created in repo/test"),
	)
}

func TestPullRequestEvent(t *testing.T) {
	domain.ProcessEvents(
		"TestUser",
		MockedEventRepository("pull_request_event"),
		MockedPrinter(t, "- TestUser closed a pull request with title 'Pull request title'"),
	)
}

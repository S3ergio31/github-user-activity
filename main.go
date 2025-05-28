package main

import (
	"os"

	"github.com/S3ergio31/github-user-activity/domain"
	"github.com/S3ergio31/github-user-activity/infrastructure"
)

func main() {
	domain.ProcessEvents(
		getUser(),
		infrastructure.HttpEventRepository,
		infrastructure.Print,
	)
}

func getUser() string {
	if len(os.Args) == 1 {
		return ""
	}

	return os.Args[1]
}

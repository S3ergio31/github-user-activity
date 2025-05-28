package main

import (
	"os"

	"github.com/S3ergio31/github-user-activity/domain"
	"github.com/S3ergio31/github-user-activity/infrastructure"
)

func main() {
	user := os.Args[1]
	domain.ProcessEvents(
		user,
		infrastructure.HttpEventRepository,
		infrastructure.Print,
	)
}

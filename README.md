# 📊 github-user-activity

A command-line tool written in Go that displays a GitHub user's recent activity directly in the terminal. Inspired by the roadmap.sh project: [GitHub User Activity](https://roadmap.sh/projects/github-user-activity).

## 🚀 What does it do?

This program consumes GitHub's public API to fetch a user's recent events and displays them in the terminal with customized messages based on the event type.

## 🛠️ Installation

```bash
git clone https://github.com/S3ergio31/github-user-activity.git
cd github-user-activity
go build -o github-activity
```

## ⚙️ Usage

```bash
./github-activity <username>
```

Example:

```bash
./github-activity S3ergio31
```

## 📋 Sample Output

```
- Pushed 1 commits to S3ergio31/github-user-activity
- A new 'branch' was created in S3ergio31/github-user-activity
- A new 'repository' was created in S3ergio31/url-shortener
- A new 'tag' was created in S3ergio31/url-shortener
```

## 🧪 Tests

```bash
go test ./...
```

## 🌱 Learnings

This project allowed me to:

- Practice consuming REST APIs in Go.
- Handle and parse JSON data.
- Apply principles of clean and modular design.

## 📚 Technologies

- [Go](https://golang.org/)
- GitHub API

## 📄 License

MIT License © [S3ergio31](https://github.com/S3ergio31)

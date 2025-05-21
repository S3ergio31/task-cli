# 🧩 task-cli

A **minimalist terminal task tracker**, written in Go.  Inspired by [roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker), this project helped me learn and practice the Go language from scratch.

## 🚀 What is task-cli?

`task-cli` is a small command-line tool to manage to-do tasks. It allows adding, updating descriptions, listing, changing status, and deleting tasks directly from the terminal, without needing a database or external dependencies.

## 📦 Versions

### ✅ v1.0.0
> First functional version.
- All code in a single file.
- Familiarization with Go syntax and basic structure.

### 📁 v2.0.0
> Organization and quality.
- Modularized code using packages.
- Unit tests with **100% coverage**.

### 🧠 v3.0.0
> Architectural improvement.
- Introduced a generic `Command` interface to represent each command.
- Cleaner, more maintainable, and extensible code.

## 🛠️ Installation

```bash
git clone https://github.com/tuusuario/task-cli.git
cd task-cli/src/task-cli
go build -o task
```

## ⚙️ Usage

```bash
./task add "Learn Go"
./task update <uuid> "learn go"
./task list
./task mark <uuid> done
./task remove <uuid>
```

## 🧾 Available Commands

- `add <description>` – Adds a new task.
- `update <uuid> <description>` – Updates the description of a task.
- `list <status>` – Lists all tasks by status: todo, in-progress, done. If not specified, all tasks are shown.
- `mark <uuid> <status>` – Marks a task as todo, in-progress, or done.
- `remove <uuid>` – Deletes a task.

## 🧪 Tests

```bash
go test .
```

## 🌱 Learnings

This project allowed me to:
- Understand how Go works from the ground up.
- Apply good testing practices and code organization.
- Design an extensible architecture using interfaces.

## 📚 Technologies

- [Go](https://golang.org/)
- Terminal / CLI

## 📄 License

MIT License © [Sergio Fidelis](https://github.com/S3ergio31)

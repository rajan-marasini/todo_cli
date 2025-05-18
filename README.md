# Go Todo CLI App

A simple command-line TODO application written in Go. This application allows users to manage their task list using various command-line flags. Tasks are stored in a local JSON file, making the application lightweight and persistent between sessions.

## Features

-   ✅ Add new todo items
-   📝 Edit existing todos by index
-   ❌ Delete todos
-   🔁 Toggle completion status of todos
-   📋 List all todos in a neat table format
-   💾 Persistent storage using a JSON file

## File Structure

```

.
├── main.go # Entry point of the application
├── storage.go # Handles saving and loading todos to/from JSON
├── command.go # Handles command-line argument parsing
├── todo.go # Defines todo data structures and logic

```

## Getting Started

### Prerequisites

-   Go 1.18 or higher

### Build

```bash
go build -o todoapp
```

### Usage

Run the compiled binary with one or more of the following flags:

```bash
./todoapp --add "Buy groceries"
./todoapp --edit 0 "Buy groceries and cook dinner"
./todoapp --delete 0
./todoapp --toggle 0
./todoapp --list
```

### Flags

| Flag       | Description                            |
| ---------- | -------------------------------------- |
| `--add`    | Add a new todo item                    |
| `--edit`   | Edit an existing todo item by index    |
| `--delete` | Delete a todo item by index            |
| `--toggle` | Toggle the completion status of a todo |
| `--list`   | List all todo items                    |

## Example

```bash
./todoapp --add "Write documentation"
./todoapp --list
```

Expected Output:

```
+------+-----------------------+-----------+---------------------+
|  ID  |        TITLE          | COMPLETED |     CREATED AT      |
+------+-----------------------+-----------+---------------------+
|  0   | Write documentation   |   false   | 2025-05-09 19:00:00 |
+------+-----------------------+-----------+---------------------+
```

## Dependencies

-   [aquasecurity/table](https://github.com/aquasecurity/table) – for rendering tabular output in CLI

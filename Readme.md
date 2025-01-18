# Task Tracker CLI Application

This is a command-line application to track and manage your tasks. You can add tasks, update them, mark them as completed, delete them, and list all of them.

## Features

- **Add** a new task.
- **Update** an existing task.
- **Mark** a task as completed.
- **Delete** a task.
- **List** all tasks with their completion status.

## Prerequisites

- Go 1.17 or higher
- A terminal or command prompt

## Installation

### 1. Clone the repository

Clone the repository to your local machine:

```bash
git clone https://github.com/sumitnegi7/golang-todo-cli-cobra
cd task-tracker-go-cli
```

### 2. Build the binary

To build the executable for your platform, use the `go build` command:

```bash
go build -o task-tracker
```

This will generate a binary executable called `task-tracker` in your project directory.

### 3. Install the binary globally (optional)

If you want to install the application globally to run it from anywhere on your system, use the `go install` command:

```bash
go install
```

This will install the executable into your `$GOPATH/bin` directory, or `$GOBIN` if it is set.

### 4. Run the application

Once the binary is built, you can run the CLI application from your terminal:

```bash
./task-tracker [flags]
```

On Windows, you can run:

```bash
task-tracker.exe [flags]
```

## Supported Commands

### 1. **Add a Task**

Adds a new task to your task list.

```bash
task-tracker add "Task description"
```

**Example:**

```bash
task-tracker add "Eat breafast"
```

**Output:**
```
Todo added successfully. 3
```

### 2. **Update a Task**

Updates the task description by its ID.

```bash
task-tracker update <task_id> "New task description"
```

**Example:**

```bash
task-tracker update 1 "Eat eggs in breakfast"
```

**Output:**
```
Task updated successfully
```

### 3. **Mark Task as Done**

Marks the task as completed (done) by its ID.

```bash
task-tracker mark-done <task_id>
```

**Example:**

```bash
task-tracker mark-done 1
```

**Output:**
```
Task ID 1 marked as done.
```

### 4. **List All Tasks**

Lists all tasks with their statuses, including whether they are completed or not.

```bash
task-tracker list
```

**Example:**

```bash
task-tracker list
```

**Output:**
```
ID  TODO                    COMPLETED   CREATEDAT           
1   Eat eggs in breakfast  ✅  True     2025-01-17 02:00:21 
2   abc         ✅             True     2025-01-18 00:30:53 
```

### 5. **Delete a Task**

Deletes a task by its ID.

```bash
task-tracker del <task_id>
```

**Example:**

```bash
task-tracker del 1
```

**Output:**
```
Task with ID 1 deleted successfully.
```

---

## Example Workflow

1. **Add a Task:**
   ```bash
   task-tracker add "Buy groceries"
   ```

2. **Update a Task:**
   ```bash
   task-tracker update 1 "Buy milk"
   ```

3. **Mark Task as Done:**
   ```bash
   task-tracker mark-done 1
   ```

4. **List Tasks:**
   ```bash
   task-tracker list
   ```

5. **Delete a Task:**
   ```bash
   task-tracker del 1
   ```

---



## License

This project is licensed under the MIT License


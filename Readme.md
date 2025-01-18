
# 📝 Task Tracker CLI Application

Welcome to the **Task Tracker CLI** — a simple command-line tool to track and manage your tasks. You can **add**, **update**, **mark as completed**, **delete**, and **list** all your tasks with ease.

## 🚀 Features

- ➕ **Add** a new task.
- ✏️ **Update** an existing task.
- ✅ **Mark** a task as completed.
- ❌ **Delete** a task.
- 📜 **List** all tasks with their completion status.

## 🛠️ Prerequisites

- Go 1.17 or higher

## 🏗️ Installation

### 1️⃣ Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/sumitnegi7/golang-todo-cli-cobra
cd golang-todo-cli-cobra
```

### 2️⃣ Build the Binary

To build the executable for your platform, run:

```bash
go build -o task-tracker
```

This will generate a binary executable named `task-tracker` in your project directory.

### 3️⃣ Install the Binary Globally (Optional)

To install the app globally and run it from anywhere on your system, use:

```bash
go install
```

This will place the executable into your `$GOPATH/bin` directory (or `$GOBIN` if set).

### 4️⃣ Run the Application

Once the binary is built, you can run the CLI application:

On macOS/Linux:

```bash
./task-tracker [flags]
```

On Windows:

```bash
task-tracker.exe [flags]
```

## 🧑‍💻 Supported Commands

### 1️⃣ **Add a Task**

Add a new task to your task list:

```bash
task-tracker add "Task description"
```

**Example:**

```bash
task-tracker add "Eat breakfast"
```

**Output:**
```
Todo added successfully. Task ID: 3
```

### 2️⃣ **Update a Task**

Update the task description by its ID:

```bash
task-tracker update <task_id> "New task description"
```

**Example:**

```bash
task-tracker update 1 "Eat eggs for breakfast"
```

**Output:**
```
Task updated successfully
```

### 3️⃣ **Mark Task as Done**

Mark the task as completed (done) by its ID:

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

### 4️⃣ **List All Tasks**

List all tasks with their completion status:

```bash
task-tracker list
```

**Example:**

```bash
task-tracker list
```

**Output:**
```
ID   TODO                        COMPLETED   CREATED AT
-----------------------------------------------------------
1    Eat eggs for breakfast      ✅ True      2025-01-17 02:00:21
2    abc                         ✅ True      2025-01-18 00:30:53
```

### 5️⃣ **Delete a Task**

Delete a task by its ID:

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

## 💡 Example Workflow

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

## 📜 License

This project is licensed under the **MIT License**.

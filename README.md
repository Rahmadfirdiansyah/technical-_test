# Stability Team Technical Test

This repository contains a simple Task Manager API built with Go and Fiber.
Your task is to improve the stability and correctness of this system.

## Setup

Install dependencies:
go mod tidy

Run the server:
go run main.go

Server will run at:
http://localhost:3000

## Available Endpoints

GET /tasks  
GET /tasks/:id  
POST /tasks  
DELETE /tasks/:id


---

## 🔍 Issues Found

1. Incorrect HTTP status code when task is not found (returned 200 instead of 404)

#### Before 

```golang
  if task == nil {
    return c.Status(200).JSON(fiber.Map{
        "error": "task not found",
    })
}
```

#### After

```golang
  if task == nil {
    return c.Status(404).JSON(fiber.Map{
        "error": "task not found",
    })
}
```

2. Missing input validation in CreateTask (allowed empty title)

#### Before 

```golang
  if err := c.BodyParser(&task); err != nil {
      return err
  }
```

#### After

```golang
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if task.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "title is required",
		})
	}
```

3. No error handling for invalid ID (strconv.Atoi error ignored)

#### Before 

```golang
  id, _ := strconv.Atoi(idParam)
```

#### After

```golang
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid task ID",
		})
	}
```

4. Delete endpoint always returned success even if task was not found

#### Before 

```golang
  store.DeleteTask(id)

  return c.JSON(fiber.Map{
      "message": "deleted",
  })
```

#### After

```golang
  deleted := store.DeleteTask(id)

	if !deleted {
		return c.Status(404).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "deleted",
	})
```

5. Task ID was not auto-generated, causing potential duplicate IDs

#### Before 

```golang
  func AddTask(task models.Task) {
    Tasks = append(Tasks, task)
  }
```

#### After

```golang
  func AddTask(task models.Task) {
    task.ID = len(Tasks) + 1
    Tasks = append(Tasks, task)
  }
```

---

## 🔧 Fixes Implemented

1. Updated HTTP status code to 404 when task is not found
2. Added validation to ensure title is not empty
3. Added error handling for invalid ID input (return 400 Bad Request)
4. Improved delete logic to return 404 if task does not exist
5. Implemented simple auto-increment ID for new tasks

---

## 🚀 Improvement

Improved error handling and validation:
- Added proper validation for request body
- Ensured consistent error responses
- Prevented invalid data from being processed

---

## 📡 API Endpoints

### GET /tasks
Retrieve all tasks

### GET /tasks/:id
Retrieve a task by ID

### POST /tasks
Create a new task

### Delete /tasks/:id
Delete a task by ID
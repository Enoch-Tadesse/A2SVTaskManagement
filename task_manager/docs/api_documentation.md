# Task Management API Documentation

Welcome to the **Task Management API** documentation! This API, built with **Go** and the **Gin Framework**, provides a simple yet powerful way to manage tasks using basic CRUD operations. It uses an in-memory database to store tasks, making it lightweight and ideal for development and testing purposes. Below, you'll find detailed information about each endpoint, including request formats, response formats, and example usage, as well as details about the routing configuration.

---

## Table of Contents
- [Overview](#overview)
- [Base URL](#base-url)
- [Router Configuration](#router-configuration)
- [Endpoints](#endpoints)
  - [1. Get All Tasks](#1-get-all-tasks)
  - [2. Get Task by ID](#2-get-task-by-id)
  - [3. Create a Task](#3-create-a-task)
  - [4. Update a Task](#4-update-a-task)
  - [5. Delete a Task](#5-delete-a-task)
- [Task Model](#task-model)
- [Error Handling](#error-handling)

---

## Overview
The Task Management API allows users to perform **Create, Read, Update, and Delete (CRUD)** operations on tasks. Tasks are stored in memory, ensuring fast access without the need for a persistent database. The API is designed to be intuitive, with clear request and response structures, proper error handling, and standard HTTP status codes.

---

## Base URL
All endpoints are accessible relative to the base URL:

```
http://localhost:8080
```

---

## Router Configuration
The API uses the **Gin Framework** to handle routing, defined in the `router/router.go` file. The router is initialized with `gin.Default()`, which sets up a default Gin router with middleware for logging and recovery. All task-related endpoints are grouped under the `/tasks` route for better organization.

- **Router Setup**: The `gin.Default()` function initializes the router with default middleware.
- **Route Grouping**: All endpoints are grouped under `/tasks` using `router.Group("/tasks")`.
- **Endpoints**:
  - `GET /tasks`: Maps to `controllers.GetAllTasks`
  - `GET /tasks/:id`: Maps to `controllers.GetTaskByID`
  - `PUT /tasks/:id`: Maps to `controllers.UpdateTask`
  - `DELETE /tasks/:id`: Maps to `controllers.DeleteTask`
  - `POST /tasks`: Maps to `controllers.AddTask`
- **Server Start**: The server runs on port `8080` using `router.Run(":8080")`.

This configuration ensures a clean and modular routing structure, making it easy to extend the API with additional endpoints in the future.

---

## Endpoints

### 1. Get All Tasks
Retrieve a list of all tasks stored in the system.

- **Method**: `GET`
- **Endpoint**: `/tasks`
- **Request Parameters**: None
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**: Array of task objects
- **Example Request**:
  ```bash
  curl -X GET http://localhost:8080/tasks
  ```
- **Example Response**:
  ```json
  [
    {
      "id": 1,
      "title": "Complete Project Proposal",
      "description": "Draft and submit the project proposal to the client.",
      "due_date": "2025-07-20",
      "status": "pending"
    },
    {
      "id": 2,
      "title": "Team Meeting",
      "description": "Discuss project milestones with the team.",
      "due_date": "2025-07-18",
      "status": "completed"
    }
  ]
  ```

### 2. Get Task by ID
Retrieve the details of a specific task by its ID.

- **Method**: `GET`
- **Endpoint**: `/tasks/:id`
- **Request Parameters**:
  - `id` (path parameter, integer): The ID of the task
- **Response**:
  - **Status Code**: `200 OK` (if found), `400 Bad Request` (invalid ID), `404 Not Found` (task not found)
  - **Body**: Task object or error message
- **Example Request**:
  ```bash
  curl -X GET http://localhost:8080/tasks/1
  ```
- **Example Response (Success)**:
  ```json
  {
    "task": {
      "id": 1,
      "title": "Complete Project Proposal",
      "description": "Draft and submit the project proposal to the client.",
      "due_date": "2025-07-20",
      "status": "pending"
    }
  }
  ```
- **Example Response (Error)**:
  ```json
  {
    "error": "task does not exist"
  }
  ```

### 3. Create a Task
Create a new task with the provided details.

- **Method**: `POST`
- **Endpoint**: `/tasks`
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "YYYY-MM-DD",
    "status": "string" // Must be "pending", "completed", or "missed"
  }
  ```
- **Response**:
  - **Status Code**: `201 Created` (if successful), `400 Bad Request` (invalid input)
  - **Body**: Created task object or error message
- **Example Request**:
  ```bash
  curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Task",
    "description": "This is a new task.",
    "due_date": "2025-07-25",
    "status": "pending"
  }'
  ```
- **Example Response (Success)**:
  ```json
  {
    "message": "task created successfully",
    "task": {
      "id": 3,
      "title": "New Task",
      "description": "This is a new task.",
      "due_date": "2025-07-25",
      "status": "pending"
    }
  }
  ```
- **Example Response (Error)**:
  ```json
  {
    "error": "title can not be empty"
  }
  ```

### 4. Update a Task
Update an existing task with new details.

- **Method**: `PUT`
- **Endpoint**: `/tasks/:id`
- **Request Parameters**:
  - `id` (path parameter, integer): The ID of the task to update
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "YYYY-MM-DD",
    "status": "string" // Must be "pending", "completed", or "missed"
  }
  ```
- **Response**:
  - **Status Code**: `200 OK` (if successful), `400 Bad Request` (invalid input), `404 Not Found` (task not found)
  - **Body**: Updated task object or error message
- **Example Request**:
  ```bash
  curl -X PUT Http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Task",
    "description": "Updated description.",
    "due_date": "2025-07-22",
    "status": "completed"
  }'
  ```
- **Example Response (Success)**:
  ```json
  {
    "message": "task updated successfully",
    "task": {
      "id": 1,
      "title": "Updated Task",
      "description": "Updated description.",
      "due_date": "2025-07-22",
      "status": "completed"
    }
  }
  ```
- **Example Response (Error)**:
  ```json
  {
    "error": "task not found"
  }
  ```

### 5. Delete a Task
Delete a task by its ID.

- **Method**: `DELETE`
- **Endpoint**: `/tasks/:id`
- **Request Parameters**:
  - `id` (path parameter, integer): The ID of the task to delete
- **Response**:
  - **Status Code**: `200 OK` (if successful), `400 Bad Request` (invalid ID), `404 Not Found` (task not found)
  - **Body**: Success message or error message
- **Example Request**:
  ```bash
  curl -X DELETE http://localhost:8080/tasks/1
  ```
- **Example Response (Success)**:
  ```json
  {
    "message": "task deleted successfully"
  }
  ```
- **Example Response (Error)**:
  ```json
  {
    "error": "task with id 1 is not found"
  }
  ```

---

## Task Model
The task object has the following structure:

| Field         | Type   | Description                                      | Constraints                              |
|---------------|--------|--------------------------------------------------|------------------------------------------|
| `id`          | Integer| Unique identifier for the task                   | Auto-generated, read-only                |
| `title`       | String | Title of the task                               | Required, cannot be empty                |
| `description` | String | Detailed description of the task                | Optional                                 |
| `due_date`    | String | Due date for the task (format: `YYYY-MM-DD`)    | Required, must be in the future          |
| `status`      | String | Status of the task                              | Must be `pending`, `completed`, or `missed` |

---

## Error Handling
The API uses standard HTTP status codes to indicate the success or failure of requests:

- **200 OK**: Request was successful.
- **201 Created**: Resource was successfully created.
- **400 Bad Request**: Invalid input (e.g., missing fields, invalid ID, or past due date).
- **404 Not Found**: Requested resource (task) does not exist.

Error responses include a JSON object with an `error` field describing the issue.

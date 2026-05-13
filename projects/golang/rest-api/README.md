# Project: Build a REST API with Gin + MongoDB + JWT Auth

## Description

Design and implement a complete REST API for a task management application using Go's Gin framework, MongoDB for data persistence, and JWT-based authentication. The API includes user registration/login, role-based access control (RBAC), CRUD endpoints for tasks, pagination, and comprehensive test coverage.

This project integrates concepts from Go fundamentals, web development, API design, and testing.

## Learning Objectives

By completing this project, you will:

- Structure a Go project with clean architecture (handler → service → repository layers)
- Implement JWT authentication with access and refresh tokens
- Design RESTful endpoints with proper HTTP methods, status codes, and error responses
- Use the MongoDB Go driver for CRUD operations and aggregation
- Build middleware for authentication, logging, and rate limiting
- Write table-driven tests with mocks and achieve >80% coverage
- Implement pagination, filtering, and sorting for list endpoints

## Prerequisites

- Go 1.22+ installed
- MongoDB 7+ running locally or via Docker
- Completed: Go Setup, Types & Variables, Functions & Methods lessons
- Familiarity with: HTTP methods, JSON, basic auth concepts

## Architecture Overview

```
┌──────────────────────────────────────────────┐
│                HTTP Client                    │
└────────────────────┬─────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────────┐
│              Gin Router                       │
│  ├─ /api/v1/auth/register    POST            │
│  ├─ /api/v1/auth/login       POST            │
│  ├─ /api/v1/auth/refresh     POST            │
│  ├─ /api/v1/tasks            GET, POST       │
│  ├─ /api/v1/tasks/:id        GET, PUT, DELETE│
│  └─ /api/v1/users/me         GET, PUT        │
└────────────────────┬─────────────────────────┘
                     │
          ┌──────────┼──────────┐
          ▼          ▼          ▼
┌──────────┐  ┌──────────┐  ┌──────────┐
│   Auth   │  │  Logger  │  │  Rate    │
│Middleware│  │Middleware │  │ Limiter  │
└────┬─────┘  └──────────┘  └──────────┘
     │
     ▼
┌──────────────────────────────────────────────┐
│              Handlers (HTTP layer)            │
│  Parse request → Call service → Format response
└────────────────────┬─────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────────┐
│              Services (Business logic)        │
│  Validation → Business rules → Call repository
└────────────────────┬─────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────────┐
│              Repository (Data access)         │
│  MongoDB CRUD operations, aggregations        │
└────────────────────┬─────────────────────────┘
                     │
                     ▼
              ┌──────────────┐
              │   MongoDB    │
              └──────────────┘
```

## Acceptance Criteria

### Authentication

- [ ] **Register** — POST /api/v1/auth/register with email, password, name. Hash password with bcrypt. Return user object (no password).
- [ ] **Login** — POST /api/v1/auth/login. Return JWT access token (15min) and refresh token (7 days).
- [ ] **Refresh** — POST /api/v1/auth/refresh. Accept refresh token, return new access token.
- [ ] **Password Requirements** — Minimum 8 characters, at least one uppercase, one lowercase, one digit.

### Task CRUD

- [ ] **Create Task** — POST /api/v1/tasks. Fields: title (required), description, priority (low/medium/high), due_date. Auto-assign to authenticated user.
- [ ] **List Tasks** — GET /api/v1/tasks. Support pagination (?page=1&limit=20), filtering (?priority=high&status=pending), and sorting (?sort=-created_at).
- [ ] **Get Task** — GET /api/v1/tasks/:id. Return 404 if not found. Users can only see their own tasks (admins see all).
- [ ] **Update Task** — PUT /api/v1/tasks/:id. Partial updates allowed. Users can only update their own tasks.
- [ ] **Delete Task** — DELETE /api/v1/tasks/:id. Soft delete (set deleted_at). Users can only delete their own tasks.

### RBAC

- [ ] **Roles** — user (default) and admin.
- [ ] **User role** — Can only access their own tasks.
- [ ] **Admin role** — Can access all tasks and manage users.

### Error Handling

- [ ] Consistent error response format: `{"error": {"code": "VALIDATION_ERROR", "message": "..."}}`
- [ ] Proper HTTP status codes (400, 401, 403, 404, 409, 500)
- [ ] Request validation with descriptive error messages

### Testing

- [ ] Unit tests for services with mocked repositories (>80% coverage)
- [ ] Integration tests for handlers using httptest
- [ ] Test all auth flows: register, login, access protected route, refresh token

## Getting Started

### Step 1: Initialize the Project

```bash
mkdir task-api && cd task-api
go mod init github.com/yourusername/task-api
```

### Step 2: Install Dependencies

```bash
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/stretchr/testify
```

### Step 3: Project Structure

```
task-api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Environment-based config
│   ├── handler/
│   │   ├── auth_handler.go
│   │   ├── task_handler.go
│   │   └── user_handler.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── logger.go
│   │   └── ratelimit.go
│   ├── model/
│   │   ├── user.go
│   │   └── task.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   └── task_repository.go
│   ├── service/
│   │   ├── auth_service.go
│   │   └── task_service.go
│   └── router/
│       └── router.go
├── pkg/
│   └── jwt/
│       └── jwt.go               # JWT helper functions
├── docker-compose.yml           # MongoDB setup
├── Makefile
└── go.mod
```

### Step 4: Start with Docker Compose for MongoDB

```yaml
services:
  mongo:
    image: mongo:7
    ports:
      - "27017:27017"
    volumes:
      - mongodata:/data/db

volumes:
  mongodata:
```

### Step 5: Build Bottom-Up

1. Models → Repository → Service → Handler → Router
2. Start with auth (register/login) before tasks
3. Add middleware after basic CRUD works
4. Write tests alongside each layer

## Hints and Tips

- **Define interfaces for repositories** — This makes testing easy. Your service depends on a repository interface, and tests provide a mock implementation.
- **Use Gin's binding tags** — `binding:"required"` on struct fields gives automatic validation.
- **Store JWT secret in environment variables** — Never hardcode secrets.
- **Use MongoDB's `_id` field** — Generate ObjectIDs server-side rather than using custom string IDs.
- **Pagination pattern** — Use skip/limit for simplicity or cursor-based for performance at scale.

## Bonus Challenges

1. **WebSocket Notifications** — Push real-time updates when a task is assigned or completed
2. **File Attachments** — Allow attaching files to tasks using GridFS or S3
3. **Task Comments** — Add a comments sub-resource on tasks
4. **API Documentation** — Generate OpenAPI/Swagger docs from code annotations
5. **Dockerize** — Create a multi-stage Dockerfile and add the API to docker-compose

## Resources

- [Gin Web Framework](https://gin-gonic.com/docs/)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [golang-jwt Library](https://github.com/golang-jwt/jwt)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Testify: Go Testing Toolkit](https://github.com/stretchr/testify)

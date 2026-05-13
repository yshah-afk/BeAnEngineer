# Project: Dockerize a Go + React App with Multi-Stage Builds

## Description

Containerize a full-stack application consisting of a Go backend API and a React frontend. Create optimized multi-stage Dockerfiles for both services, orchestrate them with Docker Compose alongside MongoDB and Redis, and implement a production-ready local development workflow with hot reload.

This project brings together Docker fundamentals, multi-stage builds, and Docker Compose concepts.

## Learning Objectives

By completing this project, you will:

- Write optimized multi-stage Dockerfiles for Go (scratch/distroless) and React (Nginx)
- Orchestrate multiple services with Docker Compose (API, frontend, database, cache)
- Configure networks for proper service isolation
- Set up volumes for both data persistence and development hot-reload
- Implement health checks and proper service dependency ordering
- Manage environment variables and secrets across environments
- Reduce Docker image sizes by 90%+ compared to naive builds

## Prerequisites

- Docker and Docker Compose installed
- Completed: Docker Fundamentals, Multi-Stage Builds, Docker Compose lessons
- Basic Go and React application code (provided or from previous projects)

## Architecture Overview

```
┌─────────────────────────────────────────────────────┐
│                  Docker Compose                      │
│                                                      │
│  ┌─────────┐     ┌─────────┐     ┌───────────────┐  │
│  │ Frontend │     │   API   │     │    Nginx      │  │
│  │ (build)  │────▶│  (Go)   │◀───▶│  (reverse     │  │
│  │ React    │     │  :8080  │     │   proxy)      │  │
│  │  → Nginx │     └────┬────┘     │   :80/:443    │  │
│  │  :3000   │          │          └───────────────┘  │
│  └─────────┘          │                              │
│                  ┌────┴────┐    ┌────────────┐       │
│                  │ MongoDB │    │   Redis    │       │
│                  │  :27017 │    │   :6379    │       │
│                  └─────────┘    └────────────┘       │
│                       │              │               │
│                  [mongodata]    [redisdata]           │
│                  (volume)      (volume)               │
└─────────────────────────────────────────────────────┘
```

## Acceptance Criteria

### Go Backend Dockerfile

- [ ] Multi-stage build with `golang:1.22-alpine` builder and `scratch` or `distroless` runner
- [ ] Dependencies cached in a separate layer (COPY go.mod, go.sum first)
- [ ] Binary compiled with `CGO_ENABLED=0` and `-ldflags="-s -w"`
- [ ] Final image size under 20 MB
- [ ] Runs as non-root user
- [ ] Health check endpoint at `/health`

### React Frontend Dockerfile

- [ ] Multi-stage build with `node:20-alpine` builder and `nginx:alpine` runner
- [ ] Node modules cached in a separate layer
- [ ] Build output served by Nginx with custom configuration
- [ ] Final image size under 50 MB
- [ ] Nginx configured with gzip, caching headers, and SPA fallback
- [ ] Runs as non-root (nginx user)

### Docker Compose

- [ ] All services defined: api, web, mongo, redis
- [ ] Custom network with proper isolation (frontend network, backend network)
- [ ] Named volumes for MongoDB and Redis data persistence
- [ ] Health checks on all services with proper `depends_on` conditions
- [ ] Environment variables via `.env` file (not hardcoded)
- [ ] Single `docker compose up` starts the entire stack

### Development Workflow

- [ ] `docker-compose.override.yml` with bind mounts for hot reload
- [ ] Go backend rebuilds on file change (using air or similar)
- [ ] Frontend dev server with HMR proxied through the stack
- [ ] Database data persists across `docker compose down` and `up`
- [ ] `docker compose down -v` cleanly removes all data

### Production Configuration

- [ ] `docker-compose.prod.yml` with optimized settings
- [ ] Resource limits (memory, CPU) on all services
- [ ] Restart policies (unless-stopped)
- [ ] No development volumes or debug ports exposed

## Getting Started

### Step 1: Create the Application Skeleton

If you don't have a Go API and React frontend from previous projects, create minimal versions:

**Go API** (`backend/main.go`):

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/api/status", statusHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("API listening on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

### Step 2: Write the Dockerfiles

Start with the Go backend Dockerfile (multi-stage), then the React frontend.

### Step 3: Create docker-compose.yml

Define all four services with proper networking, volumes, and health checks.

### Step 4: Test the Full Stack

```bash
docker compose up --build
curl http://localhost:8080/health
curl http://localhost:3000
```

### Step 5: Optimize and Measure

```bash
docker images | grep myapp    # Compare image sizes
docker stats                   # Monitor resource usage
```

## Hints and Tips

- **Layer caching is key** — Copy dependency files before source code in every Dockerfile
- **Use .dockerignore** — Exclude `node_modules`, `.git`, `vendor`, and build artifacts
- **Nginx SPA config** — Add `try_files $uri $uri/ /index.html;` for client-side routing
- **Air for Go hot-reload** — Use `cosmtrek/air` in the dev Dockerfile stage
- **Network isolation** — Put frontend and API on one network, API and databases on another. Frontend should never directly access the database.

## Bonus Challenges

1. **Add Nginx Reverse Proxy** — Route `/api/*` to the Go backend and `/` to the React frontend through a single Nginx entry point
2. **Add SSL with mkcert** — Generate local SSL certificates for HTTPS development
3. **Implement CI Build** — Create a GitHub Actions workflow that builds and tests the Docker images
4. **Image Scanning** — Add Trivy or Snyk scanning to check for vulnerabilities
5. **Docker BuildKit** — Enable BuildKit and use cache mounts for faster builds

## Resources

- [Docker Multi-Stage Builds](https://docs.docker.com/build/building/multi-stage/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Nginx Docker Best Practices](https://www.docker.com/blog/how-to-use-the-official-nginx-docker-image/)
- [Air: Live Reload for Go](https://github.com/cosmtrek/air)
- [Docker Security Best Practices](https://docs.docker.com/develop/security-best-practices/)

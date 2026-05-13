# BeAnEngineer

A comprehensive learning platform to help engineers ramp up faster across backend, frontend, AI, systems, and infrastructure.

`BeAnEngineer` ships as a full-stack web app with structured tracks, deep lesson content, interactive learning tools, and an AI tutor experience.

## What’s Included

- 6 learning tracks: `DSA`, `Golang`, `React`, `System Design`, `DevOps`, and `LLM/AI Engineering`
- 44 modules and 135 lessons written in Markdown/MDX
- interactive lesson reader with syntax highlighting and Mermaid diagrams
- quizzes, flashcards, bookmarks, notes, progress tracking, and streaks
- AI tutor support through Ollama or OpenAI-compatible APIs
- code playground support for Go, Python, and JavaScript
- admin tooling for content and lesson management

## Stack

| Layer | Tech |
| --- | --- |
| Frontend | React 19, TypeScript, Vite 8, Tailwind CSS 4, Zustand, TanStack Query |
| Backend | Go, Gin, JWT auth, MongoDB driver |
| Database | MongoDB |
| Content | Markdown/MDX lessons under `content/` |
| AI | Ollama or OpenAI-compatible API |
| Infra | Docker Compose, Docker, Kubernetes manifests, GitHub Actions |

## Repository Structure

```text
.
├── backend/              # Go API server
├── frontend/             # React application
├── content/              # All lesson content grouped by track/module
├── scripts/              # Seed and utility scripts
├── infra/                # Docker and Kubernetes assets
├── .github/workflows/    # CI/CD workflows
├── ARCHITECTURE.md       # System design and technical architecture
├── CURRICULUM.md         # Learning track and module outline
├── RESOURCES.md          # Curated external references
└── docker-compose.yml    # Local full-stack development stack
```

## Quick Start

### 1. Clone and configure

```bash
git clone https://github.com/yshah-afk/BeAnEngineer.git
cd BeAnEngineer
cp .env.example .env
```

### 2. Run the full stack with Docker

```bash
docker compose up --build
```

Services:

- frontend: [http://localhost:5173](http://localhost:5173)
- backend: [http://localhost:8080](http://localhost:8080)
- mongodb: `mongodb://localhost:27017`

### 3. Seed the database

From the backend module directory:

```bash
cd backend
CONTENT_DIR=../content go run ../scripts/seed.go
```

The seed script loads the lesson catalog from `content/` and creates the default admin user:

- email: `admin@mastery-hub.dev`
- password: `Admin123!`

## Local Development

### Backend

```bash
cd backend
go run ./cmd/server/main.go
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Optional AI Tutor

Start Ollama if you want local tutor support:

```bash
docker compose --profile ai up -d ollama
```

Then pull a model, for example:

```bash
docker exec -it $(docker ps -qf name=ollama) ollama pull llama3
```

## Documentation

- `ARCHITECTURE.md` for backend/frontend architecture, APIs, and deployment details
- `CURRICULUM.md` for the learning roadmap
- `RESOURCES.md` for recommended references and external study material

## Current Scope

The project currently includes:

- in-depth DSA coverage for arrays, linked lists, trees, graphs, sorting, dynamic programming, and backtracking
- production-oriented Golang content including concurrency, APIs, testing, profiling, and microservices
- modern React content including routing, state management, performance, testing, and full-stack patterns
- practical DevOps content including Docker, Kubernetes, Terraform, CI/CD, monitoring, security, and GitOps
- system design lessons covering distributed systems, storage, API design, observability, and real-world case studies
- AI engineering lessons covering prompting, RAG, fine-tuning, agents, deployment, and applications

## Contributing

Issues and pull requests are welcome. For larger changes, open an issue first so the direction can be discussed.

## License

This repository is distributed under the `GPL-3.0` license. See `LICENSE`.

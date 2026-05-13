# AI & Full-Stack Mastery Hub — Complete Curriculum

> **Source of truth** for all learning content across the platform.
> Every lesson includes a learning objective, difficulty level, and estimated completion time.

---

## Track 1: LLM & AI Engineering

### Module 1: Transformer Foundations

1. **Tokenization Deep Dive** — Understand how raw text is converted into numerical tokens using BPE, WordPiece, SentencePiece, and tiktoken, and learn how vocabulary size and subword strategy affect model performance. `[Beginner]` `[30 min]`
2. **Embeddings & Vector Representations** — Explore how words and sentences are represented as dense vectors through Word2Vec, GloVe, and contextual embeddings, and apply similarity metrics to measure semantic closeness. `[Beginner]` `[35 min]`
3. **Attention Mechanism** — Dissect self-attention, multi-head attention, and scaled dot-product attention to understand how transformers weigh token relationships, including attention masks and positional encoding. `[Intermediate]` `[45 min]`
4. **Transformer Architecture** — Compare encoder-decoder, encoder-only (BERT), and decoder-only (GPT) architectures, and understand cross-attention, layer normalization, and feed-forward sublayers. `[Intermediate]` `[45 min]`
5. **Model Zoo** — Survey the landscape of major LLMs — BERT, GPT-2/3/4, T5, LLaMA, Mistral, Phi, and Gemma — comparing architectures, parameter counts, training data, and ideal use cases. `[Beginner]` `[30 min]`

### Module 2: Training Paradigms

1. **Pretraining** — Understand causal language modeling (CLM) and masked language modeling (MLM) objectives, along with the data pipeline engineering and compute requirements for pretraining from scratch. `[Intermediate]` `[40 min]`
2. **Supervised Fine-Tuning (SFT)** — Learn to prepare instruction-following datasets, implement training loops, and tune hyperparameters to adapt a pretrained model to a specific task. `[Intermediate]` `[45 min]`
3. **Parameter-Efficient Fine-Tuning** — Master LoRA, QLoRA, prefix tuning, and adapter methods to fine-tune large models with minimal GPU memory, and learn when to choose each approach. `[Intermediate]` `[40 min]`
4. **PEFT with Hugging Face** — Apply the PEFT library hands-on to fine-tune models on custom datasets, merge adapter weights, and push models to the Hub. `[Intermediate]` `[45 min]`
5. **RLHF** — Understand Reinforcement Learning from Human Feedback end-to-end: collecting human preference data, training reward models, and optimizing with PPO following the InstructGPT approach. `[Advanced]` `[50 min]`
6. **DPO & Alternatives** — Compare Direct Preference Optimization, ORPO, and SimPO with classical RLHF, understanding when each method is preferred and how to implement them. `[Advanced]` `[45 min]`

### Module 3: Inference & Optimization

1. **Quantization Fundamentals** — Learn INT8, INT4, GGUF, AWQ, GPTQ, and bitsandbytes quantization techniques, and evaluate the trade-offs between model quality, speed, and memory footprint. `[Intermediate]` `[40 min]`
2. **KV Cache & Memory Management** — Understand how the key-value cache accelerates autoregressive generation, and explore paged attention and memory optimization strategies for long-context inference. `[Advanced]` `[45 min]`
3. **Speculative Decoding** — Implement speculative decoding with draft models, understand acceptance criteria, and measure throughput gains for latency-sensitive applications. `[Advanced]` `[40 min]`
4. **Batching Strategies** — Compare continuous batching, dynamic batching, and static batching to maximize throughput while managing latency in production inference servers. `[Intermediate]` `[35 min]`
5. **Benchmarking Inference** — Design and run inference benchmarks measuring tokens/sec, time-to-first-token, latency percentiles (p50/p95/p99), and build cost models for deployment decisions. `[Intermediate]` `[35 min]`

### Module 4: Frameworks & Libraries

1. **Hugging Face Ecosystem** — Navigate the Transformers, Datasets, Tokenizers, and Hub libraries, and learn to use Spaces for demos and model cards for documentation. `[Beginner]` `[30 min]`
2. **PyTorch Essentials for LLMs** — Build a working mental model of tensors, autograd, DataLoader, training loops, and mixed-precision training as applied to language model development. `[Intermediate]` `[45 min]`
3. **LangChain Core** — Construct LLM applications using chains, prompt templates, output parsers, memory modules, and callbacks for observability. `[Intermediate]` `[40 min]`
4. **LangGraph** — Design graph-based agent workflows with explicit state management, checkpointing for long-running tasks, and human-in-the-loop approval steps. `[Advanced]` `[50 min]`
5. **LlamaIndex** — Build data-aware LLM applications using data connectors, index structures, query engines, and response synthesizers for document Q&A. `[Intermediate]` `[40 min]`
6. **Other Frameworks** — Survey Haystack, CrewAI, AutoGen, and Semantic Kernel, comparing their agent orchestration models, strengths, and community adoption. `[Beginner]` `[30 min]`

### Module 5: Running Local LLMs

1. **Ollama** — Install Ollama, pull models from the library, use the REST API, create custom Modelfiles with system prompts, and configure GPU acceleration. `[Beginner]` `[25 min]`
2. **llama.cpp** — Build llama.cpp from source, understand the GGUF model format, run inference via CLI, and set up the built-in HTTP server for local serving. `[Intermediate]` `[35 min]`
3. **vLLM** — Deploy models with vLLM's PagedAttention, configure an OpenAI-compatible API server, enable tensor parallelism, and tune for production throughput. `[Intermediate]` `[40 min]`
4. **LM Studio** — Use the LM Studio GUI to download, manage, and run models locally, and expose a local API server for integration with other tools. `[Beginner]` `[20 min]`

### Module 6: RAG Systems

1. **RAG Architecture Overview** — Understand the retrieval-augmented generation pipeline end-to-end, and learn decision criteria for choosing RAG vs fine-tuning for a given use case. `[Beginner]` `[30 min]`
2. **Document Processing & Chunking** — Implement document loaders and text splitters (recursive, semantic), experiment with chunk size strategies, and extract metadata for filtering. `[Intermediate]` `[40 min]`
3. **Vector Databases** — Set up and compare Pinecone, Weaviate, Chroma, Qdrant, and pgvector — covering indexing, querying, filtering, and cost considerations. `[Intermediate]` `[45 min]`
4. **Embedding Models** — Evaluate embedding models from OpenAI, Cohere, sentence-transformers, and BGE, choosing appropriate dimensions and implementing batch embedding pipelines. `[Intermediate]` `[35 min]`
5. **Hybrid Search & Re-ranking** — Combine BM25 keyword search with vector similarity, apply cross-encoder re-rankers (Cohere Rerank), and implement reciprocal rank fusion for optimal retrieval. `[Advanced]` `[45 min]`
6. **Advanced RAG Patterns** — Implement parent-child retrieval, HyDE, self-querying retrievers, multi-index strategies, corrective RAG, and graph-based RAG for complex knowledge bases. `[Advanced]` `[50 min]`

### Module 7: AI Agents

1. **Agent Fundamentals** — Understand the ReAct reasoning pattern, observation-action loops, and compare single-agent architectures for task automation. `[Intermediate]` `[35 min]`
2. **Tool Use & Function Calling** — Implement OpenAI-style function calling with tool schemas, handle structured outputs, and build reliable tool-augmented LLM applications. `[Intermediate]` `[40 min]`
3. **Model Context Protocol (MCP)** — Learn the MCP architecture including server/client roles, tool discovery, resource management, and how MCP standardizes LLM-tool integration. `[Intermediate]` `[40 min]`
4. **Multi-Agent Orchestration** — Design multi-agent systems using CrewAI roles, AutoGen conversations, and LangGraph supervisor patterns for complex collaborative tasks. `[Advanced]` `[50 min]`
5. **Building Production Agents** — Implement error handling, fallback strategies, safety guardrails, cost controls, and debugging workflows for agents deployed in production. `[Advanced]` `[45 min]`

### Module 8: Prompt Engineering

1. **Prompt Fundamentals** — Master system/user/assistant message roles, temperature and top-p sampling parameters, and formatting conventions that improve LLM output quality. `[Beginner]` `[25 min]`
2. **Few-Shot & In-Context Learning** — Design effective few-shot prompts with strategic example selection, dynamic few-shot templating, and retrieval-augmented prompting for consistent results. `[Intermediate]` `[35 min]`
3. **Chain-of-Thought & Reasoning** — Apply Chain-of-Thought, zero-shot CoT, Tree-of-Thought, and self-consistency techniques to dramatically improve LLM reasoning on complex tasks. `[Intermediate]` `[35 min]`
4. **Structured Output** — Force LLMs to produce valid JSON using JSON mode, Pydantic model integration, and grammar-constrained generation for reliable downstream parsing. `[Intermediate]` `[30 min]`
5. **Guardrails & Safety** — Implement input/output validation pipelines, content filtering, NeMo Guardrails, and LLM-as-judge patterns to ensure safe and policy-compliant outputs. `[Intermediate]` `[35 min]`

### Module 9: Evaluation & Observability

1. **LLM Evaluation Metrics** — Apply BLEU, ROUGE, BERTScore, and human evaluation rubrics, and design LLM-as-judge scoring systems for open-ended generation tasks. `[Intermediate]` `[40 min]`
2. **RAG Evaluation with RAGAS** — Measure RAG pipeline quality using faithfulness, answer relevancy, context precision, and context recall metrics from the RAGAS framework. `[Intermediate]` `[35 min]`
3. **DeepEval & Testing** — Write unit tests for LLM applications using DeepEval assertions, create benchmark datasets, and integrate evaluation into CI pipelines. `[Intermediate]` `[35 min]`
4. **Tracing & Debugging** — Set up LangSmith, Phoenix, or Langfuse to trace LLM calls end-to-end, visualize chain execution, and identify performance bottlenecks. `[Intermediate]` `[40 min]`
5. **Hallucination Detection** — Implement factual grounding checks, citation verification, and self-reflection techniques to detect and mitigate LLM hallucinations in production. `[Advanced]` `[40 min]`

### Module 10: Safety & Alignment

1. **Red Teaming LLMs** — Classify attack taxonomies, craft adversarial prompts, and use automated red-teaming tools to systematically discover model vulnerabilities. `[Advanced]` `[45 min]`
2. **Jailbreak Defense** — Analyze common jailbreak techniques (DAN, prompt leaking, role-play exploits) and implement defense strategies including system prompt hardening. `[Advanced]` `[40 min]`
3. **PII & Data Privacy** — Detect and anonymize personally identifiable information in LLM inputs/outputs using regex, NER models, and presidio, following data handling best practices. `[Intermediate]` `[30 min]`
4. **OWASP LLM Top 10** — Study all ten risks — prompt injection, insecure output handling, training data poisoning, model theft, and more — with mitigation strategies for each. `[Intermediate]` `[40 min]`

### Module 11: Model Deployment

1. **Serving with vLLM & TGI** — Deploy LLMs using vLLM and Text Generation Inference, configuring GPU allocation, scaling replicas, and optimizing for throughput. `[Advanced]` `[45 min]`
2. **Triton Inference Server** — Set up a NVIDIA Triton model repository, configure ensemble models, enable dynamic batching, and serve multiple model versions. `[Advanced]` `[50 min]`
3. **Cost & Latency Optimization** — Implement model selection routing, response caching, request batching, and auto-scaling policies to minimize serving costs while meeting latency SLAs. `[Advanced]` `[40 min]`
4. **Production Observability** — Monitor deployed LLMs with dashboards tracking throughput, latency distributions, error rates, token usage, and cost per request. `[Intermediate]` `[35 min]`

### Track 1 Hands-on Projects

- **Project 1: Build a RAG Chatbot with Ollama + Chroma** — Design and implement a fully local RAG chatbot that ingests documents, stores embeddings in Chroma, and generates answers using Ollama, with a Streamlit UI. `[Intermediate]` `[3–4 hours]`
- **Project 2: Fine-tune a Small LLM with LoRA on Custom Data** — Prepare a custom instruction dataset, fine-tune a 7B model using LoRA/QLoRA, evaluate with RAGAS and human review, and deploy with vLLM. `[Advanced]` `[4–5 hours]`

---

## Track 2: Golang for Scalable Backends

### Module 1: Language Fundamentals

1. **Go Setup & Tooling** — Install Go, configure GOPATH and modules, set up VS Code / GoLand with debugging, and master essential go tools (fmt, vet, lint, build, test). `[Beginner]` `[25 min]`
2. **Types, Variables & Control Flow** — Work with Go's type system including basic types, structs, interfaces, type assertions, and control flow constructs (if, switch, for). `[Beginner]` `[30 min]`
3. **Functions & Methods** — Write functions with multiple returns, variadic parameters, and closures; attach methods to types with value and pointer receivers; use defer, panic, and recover. `[Beginner]` `[35 min]`
4. **Pointers & Memory** — Understand stack vs heap allocation, pointer semantics, escape analysis, and how the Go garbage collector manages memory. `[Intermediate]` `[30 min]`
5. **Error Handling Patterns** — Implement idiomatic Go error handling using the error interface, error wrapping with %w, sentinel errors, custom error types, and errors.Is/As. `[Intermediate]` `[30 min]`
6. **Packages & Modules** — Manage dependencies with go.mod, understand semantic versioning, organize code with internal packages, and resolve dependency conflicts. `[Beginner]` `[25 min]`

### Module 2: Concurrency

1. **Goroutines** — Launch goroutines, understand the Go scheduler and runtime.GOMAXPROCS, detect and prevent goroutine leaks using patterns and tooling. `[Intermediate]` `[35 min]`
2. **Channels** — Use buffered and unbuffered channels, directional channel types, the select statement, and common channel patterns for signaling and data passing. `[Intermediate]` `[40 min]`
3. **sync Package** — Apply WaitGroup, Mutex, RWMutex, Once, Pool, and sync.Map to safely share state between goroutines in concurrent programs. `[Intermediate]` `[35 min]`
4. **Context Package** — Propagate cancellation signals and deadlines through call chains using context.WithCancel, WithTimeout, and WithValue, following best practices. `[Intermediate]` `[30 min]`
5. **Concurrency Patterns** — Implement fan-in/fan-out, pipeline, worker pool, rate limiting, and errgroup patterns to build robust concurrent Go applications. `[Advanced]` `[45 min]`

### Module 3: Web Development

1. **net/http Fundamentals** — Build HTTP servers using the standard library's Handler interface, ServeMux router, middleware chaining pattern, and http.Client for outbound requests. `[Beginner]` `[35 min]`
2. **Gin Framework** — Create RESTful APIs with Gin's routing, middleware, request binding, validation tags, and structured error handling. `[Intermediate]` `[40 min]`
3. **Fiber & Echo** — Explore Fiber's Express-like API and Echo's minimalist approach, compare performance characteristics and developer ergonomics across frameworks. `[Intermediate]` `[30 min]`
4. **Request/Response Handling** — Parse and produce JSON, XML, form data, and file uploads; implement streaming responses and Server-Sent Events in Go. `[Intermediate]` `[35 min]`
5. **Database Integration** — Connect to MongoDB from Go using the official driver, manage connection pools, perform CRUD operations, and build aggregation pipelines. `[Intermediate]` `[40 min]`

### Module 4: API Design

1. **RESTful API Design** — Apply REST conventions for resource naming, API versioning, pagination, filtering, sorting, and HATEOAS links in Go services. `[Intermediate]` `[35 min]`
2. **gRPC** — Define Protocol Buffer schemas, implement gRPC services with unary and streaming RPCs, and add interceptors for logging and auth. `[Advanced]` `[45 min]`
3. **GraphQL with Go** — Build a GraphQL API using gqlgen with schema-first design, write resolvers, implement dataloaders for N+1 prevention, and handle subscriptions. `[Advanced]` `[45 min]`
4. **API Security** — Implement JWT authentication, OAuth2 flows, rate limiting middleware, input validation, and CORS configuration to secure Go APIs. `[Intermediate]` `[40 min]`

### Module 5: Testing & Architecture

1. **Testing in Go** — Write table-driven tests and subtests, use testify for assertions, create mocks and test fixtures, and measure code coverage. `[Intermediate]` `[35 min]`
2. **Benchmarks & Profiling** — Run Go benchmarks, use pprof for CPU/memory profiling, analyze traces, and identify performance bottlenecks in production code. `[Advanced]` `[40 min]`
3. **Clean Architecture** — Structure Go applications with clear layers, dependency injection via interfaces, the repository pattern, and domain-driven boundaries. `[Intermediate]` `[35 min]`
4. **Project Structure** — Organize Go projects using the standard layout conventions (internal/, cmd/, pkg/) and understand when to use each directory. `[Beginner]` `[25 min]`
5. **Building Production Services** — Implement graceful shutdown, health check endpoints, structured logging with slog, and externalized configuration for production-ready Go services. `[Advanced]` `[40 min]`

### Track 2 Hands-on Projects

- **Project 3: Build a REST API with Gin + MongoDB + JWT Auth** — Design and implement a complete REST API with user authentication, role-based access control, CRUD endpoints, pagination, and comprehensive test coverage. `[Intermediate]` `[4–5 hours]`
- **Project 4: Build a Concurrent Web Scraper with Worker Pools** — Create a concurrent web scraper using goroutines, channels, worker pools, and rate limiting that respects robots.txt and handles errors gracefully. `[Intermediate]` `[3–4 hours]`

---

## Track 3: Frontend — React + Vite + TypeScript

### Module 1: React 19 Fundamentals

1. **React 19 & Vite Setup** — Scaffold a new React 19 project with Vite, configure TypeScript in strict mode, set up ESLint/Prettier, and understand the dev/build pipeline. `[Beginner]` `[25 min]`
2. **Components & JSX** — Build functional components with typed props, use children and composition patterns, and understand the JSX compilation process. `[Beginner]` `[30 min]`
3. **Hooks Deep Dive** — Master useState, useEffect, useRef, useMemo, and useCallback with correct dependency arrays, and extract reusable logic into custom hooks. `[Intermediate]` `[45 min]`
4. **React 19 New Features** — Use the new use() hook for promises and context, implement Actions and useFormStatus for form handling, apply useOptimistic for instant UI updates, and understand Server Components. `[Intermediate]` `[40 min]`
5. **Event Handling & Forms** — Handle synthetic events, build controlled and uncontrolled forms, and implement robust form validation with React Hook Form and Zod schemas. `[Intermediate]` `[35 min]`
6. **Lists, Keys & Rendering** — Understand React's reconciliation algorithm, use keys correctly for list rendering, implement conditional rendering patterns, and apply lazy loading with Suspense. `[Beginner]` `[30 min]`

### Module 2: State Management

1. **State Patterns** — Analyze local vs global state needs, identify prop drilling problems, and choose the right state management strategy for different application scales. `[Beginner]` `[25 min]`
2. **Zustand** — Create lightweight stores with Zustand, use selectors for performance, add middleware (devtools, persist, immer), and organize stores for large apps. `[Intermediate]` `[35 min]`
3. **Redux Toolkit** — Build Redux stores with slices, handle async logic with createAsyncThunk, use RTK Query for data fetching, and understand when Redux is the right choice. `[Intermediate]` `[40 min]`
4. **TanStack Query** — Manage server state with queries, mutations, caching, stale-while-revalidate, optimistic updates, infinite scroll, and prefetching strategies. `[Intermediate]` `[40 min]`
5. **Combining Strategies** — Implement the recommended pattern of Zustand for client/UI state combined with TanStack Query for server state to minimize complexity and maximize performance. `[Intermediate]` `[30 min]`

### Module 3: Styling & UI Components

1. **Tailwind CSS v4** — Use utility-first classes, configure the @theme directive, build responsive layouts, implement dark mode, and customize the design system. `[Beginner]` `[30 min]`
2. **shadcn/ui** — Install and customize shadcn/ui components, understand the copy-paste philosophy, theme with CSS variables, and extend components for your design system. `[Beginner]` `[30 min]`
3. **Layout Patterns** — Build production layouts using Flexbox and CSS Grid, implement responsive designs with a mobile-first approach, and handle common layout challenges. `[Intermediate]` `[35 min]`
4. **Animation & Transitions** — Add polished animations with Framer Motion, use CSS transitions for micro-interactions, and implement skeleton loaders for loading states. `[Intermediate]` `[30 min]`

### Module 4: Routing & Patterns

1. **React Router v7** — Configure routes with the new data router, implement nested routes with outlets, use loaders and actions for data flow, and handle errors with error boundaries. `[Intermediate]` `[35 min]`
2. **SPA vs SSR vs SSG** — Compare Single-Page Applications, Server-Side Rendering, and Static Site Generation trade-offs, understand hydration, and decide which approach fits your use case. `[Intermediate]` `[30 min]`
3. **Code Splitting** — Implement React.lazy with dynamic imports, split code by route, analyze bundle size with visualizers, and reduce initial load time. `[Intermediate]` `[30 min]`
4. **Advanced Patterns** — Apply compound components, render props, higher-order components, and headless component patterns to build flexible, reusable UI abstractions. `[Advanced]` `[40 min]`

### Module 5: Testing & Accessibility

1. **Unit Testing with Vitest** — Set up Vitest for React projects, write component unit tests, mock modules and API calls, and track code coverage. `[Intermediate]` `[35 min]`
2. **Integration Testing** — Test user workflows with Testing Library, simulate user events, handle async operations, and write tests that reflect real usage. `[Intermediate]` `[35 min]`
3. **E2E Testing with Playwright** — Configure Playwright for end-to-end testing, create page objects, run visual regression tests, and integrate into CI pipelines. `[Intermediate]` `[40 min]`
4. **Accessibility (a11y)** — Implement ARIA attributes, use semantic HTML, ensure keyboard navigation, test with screen readers, and automate a11y audits with axe-core. `[Intermediate]` `[35 min]`
5. **Performance Optimization** — Use React DevTools Profiler to identify re-renders, apply memoization strategically, virtualize long lists, and optimize bundle size. `[Advanced]` `[40 min]`

### Track 3 Hands-on Projects

- **Project 5: Build a Dashboard with Zustand + TanStack Query + shadcn/ui** — Create a data-rich analytics dashboard with real-time updates, charts, filters, responsive layout, and dark mode using the recommended state management stack. `[Intermediate]` `[4–5 hours]`
- **Project 6: Build a Markdown Editor with Live Preview and Dark Mode** — Implement a split-pane Markdown editor with syntax highlighting, live preview, file save/load, keyboard shortcuts, and theme toggling. `[Intermediate]` `[3–4 hours]`

---

## Track 4: DevOps & Cloud Native

### Module 1: Docker

1. **Docker Fundamentals** — Understand images, containers, Dockerfiles, build context, and layer caching to build and run containerized applications. `[Beginner]` `[30 min]`
2. **Multi-Stage Builds** — Optimize Docker images with multi-stage builds to reduce image size, leverage build cache, and separate build-time from runtime dependencies. `[Intermediate]` `[30 min]`
3. **Docker Compose** — Define and run multi-service applications with Compose, configure networks and volumes, manage environment variables, and orchestrate local development stacks. `[Intermediate]` `[35 min]`
4. **Docker Networking & Storage** — Configure bridge, host, and overlay networks, use bind mounts and named volumes, and understand container-to-container communication. `[Intermediate]` `[30 min]`
5. **Docker Security** — Run containers as non-root users, scan images for vulnerabilities, manage secrets securely, and apply Docker security best practices. `[Intermediate]` `[30 min]`

### Module 2: Kubernetes

1. **Kubernetes Architecture** — Understand the control plane (API server, scheduler, controller manager, etcd), worker nodes, kubelet, and how components communicate. `[Intermediate]` `[40 min]`
2. **Core Objects** — Create and manage Pods, ReplicaSets, Deployments, Services, and Namespaces to deploy and expose applications in a cluster. `[Intermediate]` `[40 min]`
3. **Configuration** — Manage application configuration with ConfigMaps and Secrets, set environment variables, define resource requests/limits, and use volumes for persistent data. `[Intermediate]` `[35 min]`
4. **Networking** — Configure Services (ClusterIP, NodePort, LoadBalancer), set up Ingress controllers and rules, and understand Kubernetes DNS for service discovery. `[Intermediate]` `[40 min]`
5. **Helm** — Package Kubernetes applications with Helm charts, use templates and values files, manage chart repositories, and perform versioned releases and rollbacks. `[Intermediate]` `[35 min]`
6. **Operators & CRDs** — Extend Kubernetes with Custom Resource Definitions, understand the operator pattern for managing complex stateful applications, and explore the Operator SDK. `[Advanced]` `[45 min]`

### Module 3: CI/CD

1. **GitHub Actions Fundamentals** — Write workflows with jobs, steps, triggers, and secrets to automate build, test, and deployment pipelines in GitHub repositories. `[Beginner]` `[30 min]`
2. **CI Pipeline Design** — Design CI pipelines that build, test, lint, run security scans, and manage artifacts with caching for fast feedback loops. `[Intermediate]` `[35 min]`
3. **CD & GitOps** — Implement deployment strategies (blue-green, canary, rolling) and understand GitOps principles with ArgoCD for declarative continuous delivery. `[Intermediate]` `[40 min]`
4. **Container Registry** — Push and pull images from Docker Hub, GHCR, and ECR, implement image tagging strategies (semver, SHA, latest), and automate image builds. `[Beginner]` `[25 min]`

### Module 4: Observability

1. **Structured Logging** — Implement structured logging in Go with slog, define log levels and correlation IDs, and understand the ELK stack for log aggregation. `[Intermediate]` `[30 min]`
2. **Prometheus & Metrics** — Instrument Go applications with Prometheus metrics (counters, gauges, histograms), write PromQL queries, and configure alerting rules. `[Intermediate]` `[40 min]`
3. **Grafana Dashboards** — Design effective Grafana dashboards with panels, variables, and annotations, and integrate alerting with notification channels. `[Intermediate]` `[35 min]`
4. **Distributed Tracing** — Instrument applications with OpenTelemetry, understand spans and trace context propagation, and visualize traces in Jaeger for debugging distributed systems. `[Advanced]` `[40 min]`

### Track 4 Hands-on Projects

- **Project 7: Dockerize a Go + React App with Multi-Stage Builds** — Containerize a full-stack application with optimized multi-stage Dockerfiles for both Go backend and React frontend, orchestrated with Docker Compose. `[Intermediate]` `[3–4 hours]`
- **Project 8: Deploy to Kubernetes with Helm + GitHub Actions CI/CD** — Create Helm charts for the full-stack app, build a GitHub Actions pipeline for automated testing and deployment, and configure Ingress and TLS. `[Advanced]` `[4–5 hours]`

---

## Track 5: DSA & Problem Solving

### Module 1: Arrays & Strings

1. **Array Fundamentals** — Master core array operations, the two-pointer technique for sorted arrays, and the sliding window approach for subarray problems. `[Beginner]` `[35 min]`
2. **String Manipulation** — Solve pattern matching, anagram detection, and palindrome problems using efficient string processing and Go's strings.Builder. `[Beginner]` `[30 min]`
3. **Prefix Sum & Hashing** — Apply prefix sums for range queries and hash maps for O(1) lookups, frequency counting, and two-sum style problems. `[Intermediate]` `[35 min]`
4. **Sorting & Searching** — Implement binary search variants (lower/upper bound), merge sort, quick sort, and custom comparators for complex sorting needs. `[Intermediate]` `[40 min]`
5. **Practice Set** — Solve 15 curated LeetCode problems covering arrays and strings with detailed Go and Python solutions and complexity analysis. `[Intermediate]` `[60 min]`

### Module 2: Linked Lists, Stacks & Queues

1. **Linked Lists** — Implement singly and doubly linked lists, detect cycles with Floyd's algorithm, reverse lists iteratively and recursively, and merge sorted lists. `[Intermediate]` `[35 min]`
2. **Stacks** — Build stack implementations, solve problems using monotonic stacks, evaluate expressions, and find next greater elements. `[Intermediate]` `[35 min]`
3. **Queues & Deques** — Implement queues and double-ended queues, apply BFS traversal patterns, and solve sliding window maximum with deques. `[Intermediate]` `[30 min]`
4. **Practice Set** — Solve 10 curated problems covering linked lists, stacks, and queues with step-by-step solutions and visual explanations. `[Intermediate]` `[45 min]`
5. **Implementation in Go** — Write idiomatic Go implementations of linked lists, stacks, queues, and deques using generics, interfaces, and standard library containers. `[Intermediate]` `[35 min]`

### Module 3: Trees & Graphs

1. **Binary Trees** — Implement all tree traversals (inorder, preorder, postorder, level-order), construct trees from traversal pairs, and solve depth/path problems. `[Intermediate]` `[40 min]`
2. **Binary Search Trees** — Perform insert, delete, and search operations on BSTs, understand balancing concepts, and explore AVL tree rotations. `[Intermediate]` `[35 min]`
3. **Heaps & Priority Queues** — Build min/max heaps, implement heapify, solve top-K and merge-K problems, and use Go's container/heap interface. `[Intermediate]` `[40 min]`
4. **Graph Representations** — Model graphs with adjacency lists and matrices, handle weighted/unweighted and directed/undirected graphs, and convert between representations. `[Intermediate]` `[30 min]`
5. **Graph Algorithms** — Implement BFS, DFS, Dijkstra's shortest path, topological sort, union-find with path compression, and minimum spanning trees (Kruskal/Prim). `[Advanced]` `[50 min]`
6. **Practice Set** — Solve 15 curated tree and graph problems with solutions demonstrating multiple approaches and trade-offs. `[Advanced]` `[60 min]`

### Module 4: Dynamic Programming

1. **DP Fundamentals** — Identify overlapping subproblems and optimal substructure, implement memoization (top-down) and tabulation (bottom-up), and understand state transitions. `[Intermediate]` `[40 min]`
2. **1D DP** — Solve Fibonacci variants, climbing stairs, house robber, and coin change problems, building intuition for one-dimensional state transitions. `[Intermediate]` `[40 min]`
3. **2D DP** — Tackle grid path counting, longest common subsequence, edit distance, and 0/1 knapsack with two-dimensional state tables. `[Advanced]` `[45 min]`
4. **Interval & String DP** — Solve palindrome partitioning, matrix chain multiplication, and burst balloons using interval DP and string-based state definitions. `[Advanced]` `[45 min]`
5. **Practice Set** — Solve 15 curated DP problems spanning 1D, 2D, and interval categories with detailed state transition explanations. `[Advanced]` `[60 min]`

### Module 5: Greedy, Sliding Window & Two Pointers

1. **Greedy Algorithms** — Apply greedy strategies to activity selection, interval scheduling, and Huffman coding, and prove greedy correctness with exchange arguments. `[Intermediate]` `[35 min]`
2. **Sliding Window** — Master fixed-size and variable-size window techniques, and solve minimum window substring and longest substring without repeating characters. `[Intermediate]` `[35 min]`
3. **Two Pointers** — Use two-pointer patterns on sorted arrays for pair sums, container with most water, and 3Sum, understanding when to apply this technique. `[Intermediate]` `[30 min]`
4. **Practice Set** — Solve 10 curated problems covering greedy, sliding window, and two-pointer techniques with comparative solutions. `[Intermediate]` `[45 min]`

### Module 6: Complexity & Problem-Solving Frameworks

1. **Big-O Analysis** — Analyze time and space complexity rigorously, understand amortized analysis, and recognize common complexity classes across algorithms. `[Beginner]` `[30 min]`
2. **Problem-Solving Framework** — Apply a systematic approach to algorithmic problems: understand, plan, pattern-match, implement, test, and handle edge cases. `[Beginner]` `[25 min]`
3. **Mock Interview Practice** — Complete timed problem sets simulating real coding interviews, practice clear communication, and apply structured whiteboard problem-solving. `[Intermediate]` `[45 min]`

### Track 5 Hands-on Projects

- **Project 9: Implement a Custom Data Structure Library in Go** — Build production-quality implementations of an LRU Cache, Trie, and Graph with generics, full test coverage, benchmarks, and documentation. `[Intermediate]` `[4–5 hours]`
- **Project 10: Solve 50 LeetCode Problems with Documented Solutions** — Complete a curated set of 50 problems across all difficulty levels with solutions in both Go and Python, including complexity analysis and approach explanations. `[Intermediate]` `[15–20 hours]`

---

## Track 6: System Design

### Module 1: Fundamentals

1. **Scalability Basics** — Compare vertical and horizontal scaling strategies, understand load balancing algorithms (round-robin, least connections, consistent hashing), and design for growth. `[Beginner]` `[30 min]`
2. **Caching Strategies** — Implement cache-aside, write-through, and write-behind patterns, understand eviction policies (LRU, LFU, TTL), and compare Redis vs Memcached. `[Intermediate]` `[40 min]`
3. **Database Concepts** — Compare SQL and NoSQL databases, design effective indexes, implement sharding strategies, and configure replication for high availability. `[Intermediate]` `[45 min]`
4. **CAP Theorem & Consistency** — Apply the CAP theorem and PACELC framework to real system decisions, and understand the spectrum from eventual to strong consistency. `[Intermediate]` `[35 min]`
5. **Message Queues** — Design event-driven architectures with Kafka and RabbitMQ, implement pub/sub and point-to-point patterns, and handle message ordering and exactly-once delivery. `[Intermediate]` `[40 min]`
6. **Networking & CDNs** — Understand DNS resolution, CDN edge caching, reverse proxy patterns, API gateways, and real-time protocols (WebSocket vs SSE vs long polling). `[Beginner]` `[35 min]`

### Module 2: HLD Case Studies

1. **Design Twitter/X** — Design a social media platform handling feed generation with fan-out strategies, real-time updates, trending topics, and search at scale. `[Advanced]` `[50 min]`
2. **Design Netflix** — Architect a video streaming service covering CDN design, adaptive bitrate streaming, recommendation engine, and microservice decomposition. `[Advanced]` `[50 min]`
3. **Design Uber** — Design a ride-sharing platform with location services, driver-rider matching, surge pricing algorithms, and real-time GPS tracking. `[Advanced]` `[50 min]`
4. **Design ChatGPT** — Architect an LLM serving platform covering model serving infrastructure, conversation management, streaming responses, cost optimization, and multi-tenancy. `[Advanced]` `[50 min]`
5. **Design a RAG System** — Design a production RAG pipeline covering document ingestion, vector search at scale, retrieval strategies, and quality monitoring. `[Advanced]` `[45 min]`
6. **Design a Vector Database** — Architect a vector database with storage engines, approximate nearest neighbor indexing (HNSW, IVF), distributed search, and consistency guarantees. `[Advanced]` `[50 min]`

### Module 3: Low-Level Design

1. **OOP Principles** — Apply encapsulation, inheritance, polymorphism, and composition-over-inheritance to design maintainable, extensible object-oriented systems. `[Intermediate]` `[30 min]`
2. **SOLID Principles** — Implement all five SOLID principles — single responsibility, open-closed, Liskov substitution, interface segregation, and dependency inversion — with practical Go examples. `[Intermediate]` `[40 min]`
3. **Design Patterns** — Apply factory, singleton, observer, strategy, adapter, and decorator patterns with idiomatic Go implementations and real-world use cases. `[Intermediate]` `[45 min]`
4. **LLD Practice** — Design the class structure and interfaces for a parking lot, elevator system, chess game, and in-memory file system with full UML and Go code. `[Advanced]` `[50 min]`

### Module 4: Interview Frameworks

1. **RESHADED Framework** — Apply the Requirements, Estimation, Storage, High-level, API, Detailed, Evaluation, Deployment framework to structure system design interviews systematically. `[Intermediate]` `[35 min]`
2. **The 4S Method** — Use the Scope, Service, Storage, Scale method for quick system design structuring, and compare it with RESHADED for different interview formats. `[Intermediate]` `[30 min]`
3. **Mock System Design Interview** — Walk through a full 45-minute system design interview end-to-end, practicing timing, communication, trade-off analysis, and diagram drawing. `[Advanced]` `[50 min]`

### Track 6 Hands-on Projects

- **Project 11: Design a Complete URL Shortener Architecture** — Produce a full system design document for a URL shortener covering high-level design, low-level design, API specifications, database schema, caching strategy, and scaling plan. `[Intermediate]` `[3–4 hours]`
- **Project 12: Build a Simplified Message Queue in Go** — Implement a basic message queue supporting pub/sub, persistence, consumer groups, and at-least-once delivery in Go with full tests. `[Advanced]` `[4–5 hours]`

---

## Summary

| Track | Modules | Lessons | Projects | Estimated Hours |
|-------|---------|---------|----------|-----------------|
| 1 — LLM & AI Engineering | 11 | 55 | 2 | ~36 |
| 2 — Golang for Scalable Backends | 5 | 25 | 2 | ~15 |
| 3 — Frontend (React + Vite + TS) | 5 | 24 | 2 | ~15 |
| 4 — DevOps & Cloud Native | 4 | 19 | 2 | ~12 |
| 5 — DSA & Problem Solving | 6 | 28 | 2 | ~32 |
| 6 — System Design | 4 | 19 | 2 | ~13 |
| **Totals** | **35** | **170** | **12** | **~123** |

> **Note:** Estimated hours include lesson time only. Hands-on project time (approximately 45–55 additional hours) is not included in per-track lesson estimates but is accounted for in planning. Total platform learning time with projects: **~170 hours**.

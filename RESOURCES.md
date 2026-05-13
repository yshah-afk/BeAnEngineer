# RESOURCES.md — Curated Learning Resources

> **AI & Full-Stack Mastery Hub** — A hand-picked collection of high-quality external resources to accelerate your learning across all six tracks.

## How to Use This List

- **Organized by track & topic** — Resources follow the same structure as the platform's six learning tracks so you can find what you need quickly.
- **Tagged by type** — Every entry carries a `[Type]` tag (`Article`, `Course`, `Video`, `Docs`, `Paper`, `Repo`, `Tool`, `Book`) so you can filter for your preferred learning style.
- **Tagged by year** — A `[Year]` tag helps you gauge currency. We prioritize resources from 2024–2026 while keeping foundational classics.
- **Quality over quantity** — Each link was selected for depth, accuracy, and community reputation. When free and paid alternatives exist, the free option is listed first.

---

## Track 1: LLM & AI Engineering

### Transformer Foundations

- [Attention Is All You Need (Vaswani et al.)](https://arxiv.org/abs/1706.03762) — The original transformer paper that started it all. `[Paper]` `[2017]`
- [The Illustrated Transformer (Jay Alammar)](https://jalammar.github.io/illustrated-transformer/) — Visual, intuitive guide to the transformer architecture. `[Article]` `[2018]`
- [Andrej Karpathy — Let's build GPT from scratch](https://www.youtube.com/watch?v=kCc8FmEb1nY) — Building a GPT model step by step in code. `[Video]` `[2023]`
- [3Blue1Brown — Transformers Explained Visually](https://www.youtube.com/watch?v=wjZofJX0v4M) — The math behind attention with beautiful animations. `[Video]` `[2024]`
- [Hugging Face NLP Course](https://huggingface.co/learn/nlp-course) — Free, comprehensive NLP course covering tokenizers through generation. `[Course]` `[2024]`
- [Stanford CS224N: NLP with Deep Learning](https://web.stanford.edu/class/cs224n/) — Full Stanford NLP course with lectures and assignments. `[Course]` `[2024]`
- [The Illustrated GPT-2 (Jay Alammar)](https://jalammar.github.io/illustrated-gpt2/) — Visual walkthrough of GPT-2 architecture and generation. `[Article]` `[2019]`

### Training & Fine-Tuning

- [Hugging Face PEFT Library](https://github.com/huggingface/peft) — Parameter-efficient fine-tuning methods (LoRA, Prefix Tuning, etc.). `[Repo]` `[2025]`
- [LoRA Paper (Hu et al.)](https://arxiv.org/abs/2106.09685) — Low-Rank Adaptation of large language models. `[Paper]` `[2021]`
- [QLoRA Paper](https://arxiv.org/abs/2305.14314) — Efficient finetuning of quantized LLMs. `[Paper]` `[2023]`
- [RLHF Blog (Hugging Face)](https://huggingface.co/blog/rlhf) — Practical guide to Reinforcement Learning from Human Feedback. `[Article]` `[2023]`
- [DPO Paper — Direct Preference Optimization](https://arxiv.org/abs/2305.18290) — Simpler alternative to RLHF for alignment. `[Paper]` `[2023]`
- [Unsloth](https://github.com/unslothai/unsloth) — 2x faster LLM fine-tuning with 80% less memory. `[Repo]` `[2025]`
- [Axolotl](https://github.com/axolotl-ai-cloud/axolotl) — Streamlined fine-tuning tool supporting multiple techniques. `[Repo]` `[2025]`
- [TheAiSingularity Curriculum](https://theaisingularity.org/curriculum/) — Free, hands-on program with 64 notebooks covering pretraining through alignment. `[Course]` `[2025]`

### Inference & Optimization

- [vLLM Documentation](https://docs.vllm.ai/) — High-throughput, memory-efficient LLM serving engine. `[Docs]` `[2025]`
- [GGUF Format Spec](https://github.com/ggerganov/ggml/blob/master/docs/gguf.md) — GGUF quantization format specification. `[Docs]` `[2024]`
- [AWQ Paper — Activation-aware Weight Quantization](https://arxiv.org/abs/2306.00978) — Hardware-efficient quantization for LLMs. `[Paper]` `[2023]`
- [GPTQ Paper](https://arxiv.org/abs/2210.17323) — Accurate post-training quantization for generative models. `[Paper]` `[2023]`
- [TensorRT-LLM (NVIDIA)](https://github.com/NVIDIA/TensorRT-LLM) — Optimized LLM inference on NVIDIA GPUs. `[Repo]` `[2025]`
- [SGLang](https://github.com/sgl-project/sglang) — Fast serving framework for LLMs and vision-language models. `[Repo]` `[2025]`

### Frameworks & Libraries

- [LangChain Documentation](https://python.langchain.com/docs/) — LLM application framework for chains, agents, and tools. `[Docs]` `[2025]`
- [LangGraph Documentation](https://langchain-ai.github.io/langgraph/) — Graph-based stateful agent framework built on LangChain. `[Docs]` `[2026]`
- [LlamaIndex Documentation](https://docs.llamaindex.ai/) — Data framework for connecting LLMs with external data. `[Docs]` `[2025]`
- [CrewAI Documentation](https://docs.crewai.com/) — Multi-agent role-based orchestration framework. `[Docs]` `[2026]`
- [AutoGen (Microsoft)](https://github.com/microsoft/autogen) — Framework for multi-agent conversations and workflows. `[Repo]` `[2025]`
- [Haystack Documentation](https://docs.haystack.deepset.ai/) — Composable LLM orchestration for search and RAG. `[Docs]` `[2025]`
- [PyTorch Tutorials](https://pytorch.org/tutorials/) — Official PyTorch learning resources and examples. `[Docs]` `[2025]`
- [Hugging Face Transformers](https://huggingface.co/docs/transformers/) — The de-facto transformer models library. `[Docs]` `[2025]`
- [OpenAI Agents SDK](https://openai.github.io/openai-agents-python/) — OpenAI's official SDK for building agents with tool use. `[Docs]` `[2025]`
- [Google ADK (Agent Development Kit)](https://google.github.io/adk-docs/) — Google's framework for building AI agents. `[Docs]` `[2025]`

### Local LLMs

- [Ollama](https://ollama.com/) — Run open-source LLMs locally with a single command. `[Tool]` `[2025]`
- [llama.cpp](https://github.com/ggerganov/llama.cpp) — Lightweight C/C++ LLM inference engine. `[Repo]` `[2025]`
- [LM Studio](https://lmstudio.ai/) — Desktop app for discovering and running local LLMs. `[Tool]` `[2025]`
- [Open WebUI](https://github.com/open-webui/open-webui) — Self-hosted ChatGPT-like interface for local LLMs. `[Repo]` `[2025]`
- [Jan](https://jan.ai/) — Open-source, offline-first alternative to ChatGPT. `[Tool]` `[2025]`

### RAG (Retrieval-Augmented Generation)

- [RAG Survey Paper](https://arxiv.org/abs/2312.10997) — Comprehensive survey of RAG techniques and architectures. `[Paper]` `[2024]`
- [Complete Guide to RAG: Naive, Advanced, and Graph RAG](https://www.mrlatte.net/en/research/2026/04/27/rag-complete-guide/) — End-to-end guide with theory and runnable code. `[Article]` `[2026]`
- [Advanced RAG — Hybrid Search, Reranking & Knowledge Graphs](https://myengineeringpath.dev/genai-engineer/advanced-rag/) — Practical advanced RAG implementation guide. `[Article]` `[2026]`
- [Pinecone Learning Center](https://www.pinecone.io/learn/) — Vector database concepts and RAG tutorials. `[Docs]` `[2025]`
- [Weaviate Documentation](https://weaviate.io/developers/weaviate) — AI-native vector search engine docs. `[Docs]` `[2025]`
- [Chroma Documentation](https://docs.trychroma.com/) — Open-source AI-native embedding database. `[Docs]` `[2025]`
- [Qdrant Documentation](https://qdrant.tech/documentation/) — High-performance vector similarity search engine. `[Docs]` `[2025]`
- [pgvector](https://github.com/pgvector/pgvector) — Open-source vector similarity extension for PostgreSQL. `[Repo]` `[2025]`
- [Milvus Documentation](https://milvus.io/docs) — Scalable open-source vector database for AI applications. `[Docs]` `[2025]`

### Agents & Tool Use

- [ReAct Paper — Reasoning + Acting](https://arxiv.org/abs/2210.03629) — Synergizing reasoning and acting in language models. `[Paper]` `[2023]`
- [Model Context Protocol (MCP) Specification](https://modelcontextprotocol.io/) — The open standard for connecting AI to tools and data (by Anthropic). `[Docs]` `[2025]`
- [Building Production-Grade AI Agents with MCP](https://dev.to/nebulagg/building-production-grade-ai-agents-with-mcp-a-complete-guide-for-2026-3bo2) — Practical guide to MCP-powered agents. `[Article]` `[2026]`
- [OpenAI Function Calling Guide](https://platform.openai.com/docs/guides/function-calling) — Tool use and function calling with OpenAI models. `[Docs]` `[2025]`
- [mcp-agent](https://github.com/lastmile-ai/mcp-agent) — Composable agent framework built on MCP. `[Repo]` `[2025]`
- [Anthropic Prompt Engineering for Agents](https://docs.anthropic.com/en/docs/build-with-claude/prompt-engineering) — Official guide for building Claude-based agents. `[Docs]` `[2025]`
- [Toolformer Paper](https://arxiv.org/abs/2302.04761) — Teaching language models to use tools autonomously. `[Paper]` `[2023]`

### Prompt Engineering

- [OpenAI Prompt Engineering Guide](https://platform.openai.com/docs/guides/prompt-engineering) — Official strategies for getting better results from LLMs. `[Docs]` `[2025]`
- [Prompt Engineering Guide (DAIR.AI)](https://www.promptingguide.ai/) — Comprehensive community-maintained prompting resource. `[Article]` `[2025]`
- [Chain-of-Thought Prompting Paper](https://arxiv.org/abs/2201.11903) — Eliciting reasoning in large language models. `[Paper]` `[2022]`
- [Tree of Thoughts Paper](https://arxiv.org/abs/2305.10601) — Deliberate problem solving with LLMs via tree search. `[Paper]` `[2023]`
- [Anthropic's Prompt Engineering Guide](https://docs.anthropic.com/en/docs/build-with-claude/prompt-engineering) — Best practices for Claude models. `[Docs]` `[2025]`
- [Google's Prompt Engineering Guide](https://ai.google.dev/gemini-api/docs/prompting-intro) — Prompting strategies for Gemini. `[Docs]` `[2025]`

### Evaluation & Safety

- [RAGAS Documentation](https://docs.ragas.io/) — Evaluation framework for RAG pipelines. `[Docs]` `[2025]`
- [DeepEval](https://github.com/confident-ai/deepeval) — Open-source LLM evaluation and testing framework. `[Repo]` `[2025]`
- [LangSmith Documentation](https://docs.smith.langchain.com/) — LLM observability, tracing, and evaluation platform. `[Docs]` `[2025]`
- [OWASP Top 10 for LLM Applications](https://owasp.org/www-project-top-10-for-large-language-model-applications/) — Security risks specific to LLM-powered apps. `[Docs]` `[2024]`
- [Weights & Biases Prompts](https://docs.wandb.ai/guides/prompts/) — LLM experiment tracking and evaluation. `[Docs]` `[2025]`
- [Braintrust](https://www.braintrust.dev/) — LLM evaluation, logging, and experimentation platform. `[Tool]` `[2025]`

---

## Track 2: Golang

### Language Fundamentals

- [Go Official Tutorial](https://go.dev/doc/tutorial/) — Getting started with Go from the official team. `[Docs]` `[2025]`
- [Go by Example](https://gobyexample.com/) — Annotated example programs covering every Go feature. `[Article]` `[2025]`
- [Effective Go](https://go.dev/doc/effective_go) — Idiomatic Go conventions and best practices. `[Docs]` `[2025]`
- [The Go Programming Language Specification](https://go.dev/ref/spec) — Complete language specification. `[Docs]` `[2025]`
- [100 Go Mistakes and How to Avoid Them](https://100go.co/) — Common pitfalls with fixes and explanations. `[Book]` `[2024]`
- [Let's Go! (Alex Edwards)](https://lets-go.alexedwards.net/) — Building professional web applications in Go. `[Book]` `[2024]`
- [Let's Go Further! (Alex Edwards)](https://lets-go-further.alexedwards.net/) — Advanced API and web application patterns. `[Book]` `[2024]`
- [Go Wiki](https://go.dev/wiki/) — Community-maintained Go knowledge base. `[Docs]` `[2025]`

### Web Development & APIs

- [Gin Documentation](https://gin-gonic.com/docs/) — High-performance HTTP framework for Go. `[Docs]` `[2025]`
- [Fiber Documentation](https://docs.gofiber.io/) — Express.js-inspired Go web framework. `[Docs]` `[2025]`
- [Echo Documentation](https://echo.labstack.com/docs) — Minimalist, high-performance Go web framework. `[Docs]` `[2025]`
- [Chi Router](https://github.com/go-chi/chi) — Lightweight, composable HTTP router for Go. `[Repo]` `[2025]`
- [Go MongoDB Driver](https://www.mongodb.com/docs/drivers/go/current/) — Official MongoDB driver for Go. `[Docs]` `[2025]`
- [GORM Documentation](https://gorm.io/docs/) — Full-featured ORM for Go. `[Docs]` `[2025]`
- [sqlx](https://github.com/jmoiron/sqlx) — Extensions to Go's standard `database/sql`. `[Repo]` `[2025]`

### Concurrency

- [Go Concurrency Patterns (Rob Pike)](https://www.youtube.com/watch?v=f6kdp27TYZs) — Classic talk on Go concurrency patterns. `[Video]` `[2012]`
- [Advanced Go Concurrency Patterns (Sameer Ajmani)](https://www.youtube.com/watch?v=QDDwwePbDtw) — Advanced patterns from the Go team. `[Video]` `[2013]`
- [Concurrency in Go (Katherine Cox-Buday)](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) — Definitive book on Go concurrency. `[Book]` `[2017]`
- [Go Concurrency Guide](https://github.com/luk4z7/go-concurrency-guide) — Practical concurrency guide with examples. `[Repo]` `[2024]`

### gRPC & API Design

- [gRPC-Go Documentation](https://grpc.io/docs/languages/go/) — Building gRPC services in Go. `[Docs]` `[2025]`
- [gqlgen](https://gqlgen.com/) — Type-safe GraphQL server generation for Go. `[Docs]` `[2025]`
- [buf](https://buf.build/docs/) — Modern Protobuf tooling for API design. `[Docs]` `[2025]`
- [Connect-Go](https://connectrpc.com/docs/go/getting-started/) — Simple, reliable RPC framework compatible with gRPC. `[Docs]` `[2025]`

### Microservices & Architecture

- [Go Microservices Patterns (AsyncSquadLabs)](https://asyncsquadlabs.com/blog/microservices-go-best-practices/) — Production architecture patterns and best practices. `[Article]` `[2025]`
- [Go Kit](https://gokit.io/) — Toolkit for building microservices in Go. `[Docs]` `[2025]`
- [Go Micro](https://github.com/go-micro/go-micro) — Framework for distributed systems and microservices. `[Repo]` `[2025]`
- [KodeKloud Advanced Golang](https://kodekloud.com/courses/advanced-golang/) — Advanced Go with concurrency, reflection, and APIs. `[Course]` `[2025]`

### Testing & Tooling

- [Go Testing Documentation](https://pkg.go.dev/testing) — Standard library testing package reference. `[Docs]` `[2025]`
- [testify](https://github.com/stretchr/testify) — Assertions, mocks, and test suites for Go. `[Repo]` `[2025]`
- [GoMock](https://github.com/uber-go/mock) — Mocking framework for Go interfaces. `[Repo]` `[2025]`
- [golangci-lint](https://golangci-lint.run/) — Fast, configurable Go linter aggregator. `[Tool]` `[2025]`
- [Go Vulnerability Database](https://pkg.go.dev/vuln/) — Official Go vulnerability checking. `[Tool]` `[2025]`

---

## Track 3: React Frontend

### React Core

- [React Official Documentation](https://react.dev/) — Complete React 19 documentation with interactive examples. `[Docs]` `[2025]`
- [React 19 Release Blog Post](https://react.dev/blog/2024/12/05/react-19) — Overview of React 19 features and changes. `[Article]` `[2024]`
- [Vite Documentation](https://vite.dev/) — Next-generation frontend build tool. `[Docs]` `[2025]`
- [The Joy of React (Josh Comeau)](https://www.joyofreact.com/) — Modern React course updated for React Compiler and RSC. `[Course]` `[2025]`
- [Epic Web (Kent C. Dodds)](https://www.epicweb.dev/) — Full-stack web development curriculum. `[Course]` `[2025]`
- [Scrimba — Learn React](https://scrimba.com/learn-react-c0e) — Free interactive React course aligned with official docs. `[Course]` `[2025]`

### Next.js & Server Components

- [Next.js Documentation](https://nextjs.org/docs) — Official Next.js 15 App Router documentation. `[Docs]` `[2025]`
- [Complete Next.js 15 Guide (Wojciechowski)](https://wojciechowski.app/en/articles/nextjs-complete-guide) — Modern React framework tutorial covering App Router and RSC. `[Article]` `[2025]`
- [React Server Components: Client vs Server Guide](https://theartcode.dev/en/articles/react-server-components-when-to-use-client-vs-server) — When to use server vs client components. `[Article]` `[2025]`
- [Vercel AI SDK](https://sdk.vercel.ai/docs) — Build AI-powered apps with React Server Components. `[Docs]` `[2025]`

### State Management

- [Zustand](https://github.com/pmndrs/zustand) — Small, fast, scalable state management. `[Repo]` `[2025]`
- [Redux Toolkit](https://redux-toolkit.js.org/) — Official, opinionated Redux toolset. `[Docs]` `[2025]`
- [TanStack Query (React Query)](https://tanstack.com/query/latest) — Powerful async state management for server state. `[Docs]` `[2025]`
- [Jotai](https://jotai.org/) — Primitive and flexible state management for React. `[Docs]` `[2025]`

### Styling & UI

- [Tailwind CSS v4 Documentation](https://tailwindcss.com/docs) — Utility-first CSS framework. `[Docs]` `[2025]`
- [shadcn/ui](https://ui.shadcn.com/) — Re-usable component collection built on Radix + Tailwind. `[Docs]` `[2025]`
- [Radix UI Primitives](https://www.radix-ui.com/primitives) — Unstyled, accessible UI component primitives. `[Docs]` `[2025]`
- [Framer Motion](https://www.framer.com/motion/) — Production-ready animation library for React. `[Docs]` `[2025]`

### Testing

- [Vitest](https://vitest.dev/) — Blazing fast unit testing powered by Vite. `[Docs]` `[2025]`
- [Playwright](https://playwright.dev/) — Cross-browser end-to-end testing framework. `[Docs]` `[2025]`
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/) — Testing React components the way users interact with them. `[Docs]` `[2025]`
- [Storybook](https://storybook.js.org/) — UI component workshop for development and testing. `[Docs]` `[2025]`

### TypeScript

- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/) — Official TypeScript handbook and reference. `[Docs]` `[2025]`
- [Total TypeScript (Matt Pocock)](https://www.totaltypescript.com/) — Advanced TypeScript for professionals. `[Course]` `[2025]`
- [Type Challenges](https://github.com/type-challenges/type-challenges) — TypeScript type system exercises. `[Repo]` `[2025]`
- [TypeScript Deep Dive (Basarat)](https://basarat.gitbook.io/typescript/) — Free, comprehensive TypeScript book. `[Book]` `[2024]`

### Performance & Tooling

- [React DevTools](https://react.dev/learn/react-developer-tools) — Browser extension for debugging React apps. `[Tool]` `[2025]`
- [Million.js](https://million.dev/) — Automatic React performance optimization. `[Repo]` `[2025]`
- [Biome](https://biomejs.dev/) — Fast formatter and linter for JS/TS (Prettier + ESLint replacement). `[Tool]` `[2025]`

---

## Track 4: DevOps & Cloud Native

### Docker & Containers

- [Docker Official Documentation](https://docs.docker.com/) — Complete Docker reference and guides. `[Docs]` `[2025]`
- [Docker Best Practices](https://docs.docker.com/build/building/best-practices/) — Dockerfile authoring best practices. `[Docs]` `[2025]`
- [Podman Documentation](https://podman.io/docs) — Daemonless container engine (Docker alternative). `[Docs]` `[2025]`
- [Dive](https://github.com/wagoodman/dive) — Tool for exploring Docker image layers and reducing size. `[Repo]` `[2025]`

### Kubernetes

- [Kubernetes Official Documentation](https://kubernetes.io/docs/) — Comprehensive K8s reference. `[Docs]` `[2025]`
- [Helm Documentation](https://helm.sh/docs/) — The package manager for Kubernetes. `[Docs]` `[2025]`
- [Kustomize Documentation](https://kustomize.io/) — Template-free Kubernetes configuration. `[Docs]` `[2025]`
- [KodeKloud Kubernetes Tutorial](https://kodekloud.com/blog/kubernetes-tutorial-for-beginners-2025/) — Hands-on beginner guide. `[Article]` `[2025]`
- [Kubernetes the Hard Way (Kelsey Hightower)](https://github.com/kelseyhightower/kubernetes-the-hard-way) — Bootstrap K8s from scratch for deep understanding. `[Repo]` `[2025]`
- [K9s](https://k9scli.io/) — Terminal-based Kubernetes management UI. `[Tool]` `[2025]`
- [Lens](https://k8slens.dev/) — Kubernetes IDE for managing clusters. `[Tool]` `[2025]`

### CI/CD & GitOps

- [GitHub Actions Documentation](https://docs.github.com/en/actions) — CI/CD platform integrated with GitHub. `[Docs]` `[2025]`
- [ArgoCD Documentation](https://argo-cd.readthedocs.io/) — Declarative GitOps continuous delivery for Kubernetes. `[Docs]` `[2025]`
- [Flux Documentation](https://fluxcd.io/docs/) — CNCF GitOps toolkit for Kubernetes. `[Docs]` `[2025]`
- [Tekton Documentation](https://tekton.dev/docs/) — Cloud-native CI/CD pipelines. `[Docs]` `[2025]`

### Infrastructure as Code

- [Terraform Documentation](https://developer.hashicorp.com/terraform/docs) — Industry-standard IaC tool. `[Docs]` `[2025]`
- [Terraform Tutorials (HashiCorp)](https://developer.hashicorp.com/terraform/tutorials) — Official hands-on learning path. `[Docs]` `[2025]`
- [Pulumi Documentation](https://www.pulumi.com/docs/) — IaC using real programming languages. `[Docs]` `[2025]`
- [Crossplane Documentation](https://docs.crossplane.io/) — Kubernetes-native infrastructure management. `[Docs]` `[2026]`

### Observability

- [Prometheus Documentation](https://prometheus.io/docs/) — Open-source monitoring and alerting. `[Docs]` `[2025]`
- [Grafana Documentation](https://grafana.com/docs/) — Visualization and dashboarding platform. `[Docs]` `[2025]`
- [OpenTelemetry Documentation](https://opentelemetry.io/docs/) — Vendor-neutral observability framework for traces, metrics, and logs. `[Docs]` `[2025]`
- [Jaeger Documentation](https://www.jaegertracing.io/docs/) — Open-source distributed tracing. `[Docs]` `[2025]`
- [Loki Documentation](https://grafana.com/docs/loki/latest/) — Log aggregation system by Grafana Labs. `[Docs]` `[2025]`

### Service Mesh & Networking

- [Istio Documentation](https://istio.io/latest/docs/) — Leading service mesh for Kubernetes. `[Docs]` `[2025]`
- [Envoy Proxy Documentation](https://www.envoyproxy.io/docs/) — Cloud-native edge and service proxy. `[Docs]` `[2025]`
- [Cilium Documentation](https://docs.cilium.io/) — eBPF-powered Kubernetes networking and security. `[Docs]` `[2025]`

### Security

- [Trivy](https://github.com/aquasecurity/trivy) — Comprehensive vulnerability scanner for containers and IaC. `[Repo]` `[2025]`
- [Falco](https://falco.org/docs/) — Cloud-native runtime security (CNCF). `[Docs]` `[2025]`
- [CNCF Technology Radar](https://www.cncf.io/reports/) — Latest technology adoption insights from the CNCF community. `[Article]` `[2026]`

---

## Track 5: DSA & Problem Solving

### Practice Platforms

- [LeetCode](https://leetcode.com/) — Industry-standard coding problem platform. `[Tool]` `[2025]`
- [NeetCode Roadmap](https://neetcode.io/roadmap) — Structured problem roadmap with video explanations. `[Tool]` `[2025]`
- [HackerRank](https://www.hackerrank.com/) — Coding challenges and interview preparation. `[Tool]` `[2025]`
- [CodeForces](https://codeforces.com/) — Competitive programming contests and problems. `[Tool]` `[2025]`

### Curated Problem Lists

- [Blind 75](https://leetcode.com/discuss/general-discussion/460599/blind-75-leetcode-questions) — Essential 75 interview problems. `[Article]` `[2020]`
- [Grind 75](https://www.techinterviewhandbook.org/grind75) — Updated, customizable version of Blind 75. `[Tool]` `[2024]`
- [NeetCode 150](https://neetcode.io/practice) — Extended curated problem set with video solutions. `[Tool]` `[2025]`

### Go-Specific DSA

- [LeetCode-in-Go](https://github.com/LeetCode-in-Go/LeetCode-in-Go) — Comprehensive Go LeetCode solutions with 100% test coverage. `[Repo]` `[2025]`
- [go-dsa (spring1843)](https://github.com/spring1843/go-dsa) — 100+ DSA problems in Go across 15 topics. `[Repo]` `[2025]`
- [dsa-golang (AshBuk)](https://github.com/AshBuk/dsa-golang) — Idiomatic Go data structure and algorithm implementations. `[Repo]` `[2025]`

### Courses & Books

- [Algorithms, Part I (Sedgewick — Princeton/Coursera)](https://www.coursera.org/learn/algorithms-part1) — Classic algorithms course with rigorous treatment. `[Course]` `[2024]`
- [Algorithms, Part II (Sedgewick — Princeton/Coursera)](https://www.coursera.org/learn/algorithms-part2) — Graph algorithms, strings, and advanced topics. `[Course]` `[2024]`
- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/9780262046305/introduction-to-algorithms/) — The definitive algorithms textbook (4th edition). `[Book]` `[2022]`
- [The Algorithm Design Manual (Skiena)](https://www.algorist.com/) — Practical algorithm design with real-world focus. `[Book]` `[2020]`
- [Tech Interview Handbook](https://www.techinterviewhandbook.org/) — Free curated interview prep guide by an ex-Meta engineer. `[Article]` `[2025]`

### Visualization

- [VisuAlgo](https://visualgo.net/) — Interactive algorithm and data structure visualizations. `[Tool]` `[2025]`
- [Algorithm Visualizer](https://algorithm-visualizer.org/) — Interactive online platform for visualizing algorithms. `[Tool]` `[2025]`
- [Data Structure Visualizations (USFCA)](https://www.cs.usfca.edu/~galles/visualization/Algorithms.html) — University of San Francisco algorithm visualizations. `[Tool]` `[2024]`

---

## Track 6: System Design

### Comprehensive Guides

- [System Design Primer (Donne Martin)](https://github.com/donnemartin/system-design-primer) — The most popular open-source system design resource. `[Repo]` `[2025]`
- [ByteByteGo (Alex Xu)](https://bytebytego.com/) — Visual system design explanations and newsletter. `[Course]` `[2025]`
- [Grokking System Design Interview (Design Gurus)](https://www.designgurus.io/course/grokking-the-system-design-interview) — Leading interactive system design course. `[Course]` `[2026]`
- [Grokking the Advanced System Design Interview](https://www.designgurus.io/course/grokking-the-advanced-system-design-interview) — Deep dives into real-world distributed systems. `[Course]` `[2026]`
- [RESHADED Framework (Hello Interview)](https://www.hellointerview.com/learn/system-design/in-a-hurry/delivery) — Structured system design interview framework. `[Article]` `[2025]`
- [System Design Handbook](https://www.systemdesignhandbook.com/guides/system-design-interview/) — 2026-edition interview preparation guide. `[Article]` `[2026]`

### Books

- [Designing Data-Intensive Applications (Martin Kleppmann)](https://dataintensive.net/) — The "bible" of distributed systems and data architecture. `[Book]` `[2017]`
- [System Design Interview Vol. 1 (Alex Xu)](https://www.amazon.com/System-Design-Interview-insiders-Second/dp/B08CMF2CQF) — Step-by-step system design interview framework. `[Book]` `[2022]`
- [System Design Interview Vol. 2 (Alex Xu)](https://www.amazon.com/System-Design-Interview-Insiders-Guide/dp/1736049119) — Advanced system design case studies. `[Book]` `[2022]`
- [Understanding Distributed Systems (Roberto Vitillo)](https://understandingdistributed.systems/) — Accessible introduction to distributed systems. `[Book]` `[2022]`
- [Building Microservices (Sam Newman)](https://www.oreilly.com/library/view/building-microservices-2nd/9781492034018/) — Definitive guide to microservice architecture. `[Book]` `[2021]`

### Design Patterns

- [Refactoring Guru — Design Patterns](https://refactoring.guru/design-patterns) — Visual guide to all GoF design patterns. `[Article]` `[2025]`
- [Head First Design Patterns (2nd Ed.)](https://www.oreilly.com/library/view/head-first-design/9781492077992/) — Accessible, visual approach to design patterns. `[Book]` `[2021]`
- [Patterns of Distributed Systems (Martin Fowler)](https://martinfowler.com/articles/patterns-of-distributed-systems/) — Enterprise-grade distributed system patterns. `[Article]` `[2024]`

### Architecture & Distributed Systems

- [The Twelve-Factor App](https://12factor.net/) — Methodology for building modern SaaS applications. `[Article]` `[2017]`
- [Martin Fowler's Architecture Guide](https://martinfowler.com/architecture/) — Collected writings on software architecture. `[Article]` `[2025]`
- [CAP Theorem (Brewer)](https://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed/) — CAP revisited — understanding distributed trade-offs. `[Article]` `[2012]`
- [Google SRE Book](https://sre.google/sre-book/table-of-contents/) — Free online book on Site Reliability Engineering. `[Book]` `[2016]`
- [DDIA.fun](https://ddia.fun/) — Interactive companion to Designing Data-Intensive Applications. `[Tool]` `[2025]`

---

## General Developer Resources

### Developer Roadmaps & Curricula

- [roadmap.sh](https://roadmap.sh/) — Community-driven developer roadmaps for every role. `[Tool]` `[2025]`
- [freeCodeCamp](https://www.freecodecamp.org/) — Free, comprehensive coding curriculum with certifications. `[Course]` `[2025]`
- [The Missing Semester of Your CS Education (MIT)](https://missing.csail.mit.edu/) — Essential developer tools most schools don't teach. `[Course]` `[2024]`
- [EngineersOfAI](https://engineersofai.com/) — Production-grade AI/ML engineering curriculum. `[Course]` `[2025]`

### Developer Productivity

- [Oh My Zsh](https://ohmyz.sh/) — Framework for managing Zsh configuration. `[Tool]` `[2025]`
- [Cursor IDE](https://www.cursor.com/) — AI-powered code editor for pair programming. `[Tool]` `[2025]`
- [Warp Terminal](https://www.warp.dev/) — Modern, AI-powered terminal. `[Tool]` `[2025]`
- [lazygit](https://github.com/jesseduffield/lazygit) — Simple terminal UI for git commands. `[Repo]` `[2025]`
- [tmux Cheat Sheet](https://tmuxcheatsheet.com/) — Quick reference for terminal multiplexer. `[Article]` `[2025]`
- [GitHub Copilot](https://github.com/features/copilot) — AI pair programmer integrated into your editor. `[Tool]` `[2025]`

### Career & Interview Prep

- [Tech Interview Handbook](https://www.techinterviewhandbook.org/) — Free, curated technical interview preparation by an ex-Meta engineer. `[Article]` `[2025]`
- [Interviewing.io](https://interviewing.io/) — Anonymous mock interviews with engineers from top companies. `[Tool]` `[2025]`
- [Levels.fyi](https://www.levels.fyi/) — Compensation data and career level comparisons. `[Tool]` `[2025]`

---

## Legend

| Tag | Meaning |
|-----|---------|
| `[Article]` | Blog post, tutorial, or written guide |
| `[Book]` | Published book (print or digital) |
| `[Course]` | Structured learning program (free or paid) |
| `[Docs]` | Official documentation |
| `[Paper]` | Academic or research paper |
| `[Repo]` | Open-source repository |
| `[Tool]` | Software tool or platform |
| `[Video]` | Video content (talk, tutorial, series) |

---

> **Last Updated:** May 13, 2026
>
> Links are periodically validated for availability. If you find a broken link, please open an issue or submit a PR. Resources are reviewed quarterly to ensure they remain current and high-quality.

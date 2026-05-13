# Project: Build a RAG Chatbot with Ollama + Chroma

## Description

Build a fully local Retrieval-Augmented Generation (RAG) chatbot that ingests documents, stores embeddings in ChromaDB, retrieves relevant context for each user query, and generates answers using a local LLM via Ollama. The chatbot includes a Streamlit UI for interactive conversations.

This project ties together core concepts from the LLM & AI Engineering track: tokenization, embeddings, vector databases, retrieval strategies, and prompt engineering.

## Learning Objectives

By completing this project, you will:

- Set up and run a local LLM using Ollama
- Implement a document ingestion pipeline with chunking strategies
- Store and query embeddings using ChromaDB
- Build a retrieval pipeline with similarity search and re-ranking
- Design effective prompts that incorporate retrieved context
- Create an interactive chat UI with Streamlit
- Handle conversation history and context window management

## Prerequisites

- Python 3.10+
- Completed: Tokenization, Embeddings, and Attention lessons
- Familiarity with: RAG concepts, vector databases, prompt engineering
- Hardware: 8GB+ RAM (16GB recommended for 7B models)
- Disk space: ~5GB for model weights

## Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│                    Streamlit UI                          │
│  ┌─────────────────┐  ┌──────────────────────────────┐  │
│  │ Document Upload  │  │    Chat Interface            │  │
│  │ (PDF, TXT, MD)  │  │    [User Message]            │  │
│  └────────┬────────┘  │    [Bot Response + Sources]   │  │
│           │           └──────────┬───────────────────┘  │
└───────────┼──────────────────────┼──────────────────────┘
            │                      │
            ▼                      ▼
┌───────────────────┐   ┌──────────────────────┐
│  Ingestion Pipeline│   │   Query Pipeline      │
│  ├─ Load documents │   │  ├─ Embed query       │
│  ├─ Split into     │   │  ├─ Search ChromaDB   │
│  │   chunks        │   │  ├─ Re-rank results   │
│  ├─ Generate       │   │  ├─ Build prompt      │
│  │   embeddings    │   │  └─ Call Ollama LLM   │
│  └─ Store in       │   └──────────────────────┘
│     ChromaDB       │              │
└───────────────────┘              ▼
         │               ┌──────────────────┐
         ▼               │   Ollama (Local)  │
┌──────────────────┐     │   ├─ llama3.1     │
│  ChromaDB        │     │   ├─ mistral      │
│  (Vector Store)  │     │   └─ phi-3        │
└──────────────────┘     └──────────────────┘
```

## Acceptance Criteria

### Core Requirements

- [ ] **Document Ingestion** — Upload PDF, TXT, and Markdown files through the UI
- [ ] **Text Chunking** — Implement recursive character text splitting with configurable chunk size (default 500) and overlap (default 50)
- [ ] **Embedding Generation** — Use a sentence-transformer model (e.g., all-MiniLM-L6-v2) to embed chunks
- [ ] **Vector Storage** — Store embeddings in ChromaDB with metadata (source file, page number, chunk index)
- [ ] **Similarity Search** — Retrieve top-k relevant chunks (default k=5) using cosine similarity
- [ ] **LLM Generation** — Send retrieved context + user query to Ollama and stream the response
- [ ] **Source Attribution** — Display which document chunks were used to generate each answer
- [ ] **Conversation History** — Maintain chat history within a session (last 10 exchanges)

### UI Requirements

- [ ] Sidebar with document upload and collection management
- [ ] Chat interface with message history
- [ ] Source documents expandable below each bot response
- [ ] Settings panel for chunk size, overlap, top-k, and model selection
- [ ] Loading indicators during embedding and generation

### Quality Requirements

- [ ] Handle documents up to 100 pages without crashing
- [ ] Response latency under 10 seconds for a 7B model
- [ ] Graceful error handling when Ollama is not running
- [ ] Clear error messages for unsupported file types

## Getting Started

### Step 1: Install Ollama and Pull a Model

```bash
# Install Ollama (macOS)
brew install ollama

# Start the Ollama server
ollama serve

# Pull a model (in another terminal)
ollama pull llama3.1:8b
```

### Step 2: Set Up the Python Environment

```bash
mkdir rag-chatbot && cd rag-chatbot
python -m venv venv
source venv/bin/activate

pip install streamlit chromadb sentence-transformers langchain langchain-community
pip install pypdf python-docx tiktoken
```

### Step 3: Project Structure

```
rag-chatbot/
├── app.py                 # Streamlit main app
├── ingestion.py           # Document loading and chunking
├── embeddings.py          # Embedding generation and storage
├── retrieval.py           # Query and retrieval logic
├── llm.py                 # Ollama integration
├── config.py              # Configuration constants
├── requirements.txt
└── data/                  # Sample documents for testing
    ├── sample.pdf
    └── sample.md
```

### Step 4: Start with the Ingestion Pipeline

Build the document loader and chunker first. Test it independently before connecting to ChromaDB.

### Step 5: Build Incrementally

1. Document ingestion → ChromaDB storage
2. Query embedding → similarity search (verify retrieval quality)
3. Prompt construction → Ollama generation
4. Streamlit UI → tie everything together
5. Conversation history → context management

## Hints and Tips

- **Chunk size matters** — Start with 500 characters, 50 overlap. Too small = fragments lose context. Too large = retrieval returns irrelevant content alongside relevant content.
- **Test retrieval before generation** — Make sure the right chunks are being retrieved before adding the LLM. Bad retrieval = bad answers regardless of model quality.
- **Use metadata filters** — Store source file names as metadata in ChromaDB so users can filter retrieval to specific documents.
- **Context window management** — Calculate how many tokens your context + history + system prompt consume. Leave room for the model's response.
- **Streaming** — Use Ollama's streaming API for a better user experience. Streamlit supports `st.write_stream()`.

## Bonus Challenges

1. **Hybrid Search** — Implement BM25 keyword search alongside vector search and combine results with reciprocal rank fusion
2. **Multi-Model Support** — Let users switch between different Ollama models mid-conversation
3. **Document Comparison** — Add a mode where users can ask questions that compare information across multiple documents
4. **Export Conversations** — Export chat history as Markdown with source citations
5. **Evaluation** — Implement basic RAG evaluation: create a test set of question-answer pairs and measure retrieval recall and answer quality

## Resources

- [Ollama API Documentation](https://github.com/ollama/ollama/blob/main/docs/api.md)
- [ChromaDB Getting Started](https://docs.trychroma.com/getting-started)
- [LangChain Text Splitters](https://python.langchain.com/docs/modules/data_connection/document_transformers/)
- [Streamlit Chat Elements](https://docs.streamlit.io/develop/api-reference/chat)
- [Sentence Transformers](https://www.sbert.net/docs/usage/semantic_textual_similarity.html)

# Project: Implement a Custom Data Structure Library in Go

## Description

Build production-quality implementations of three fundamental data structures in Go: an LRU Cache, a Trie (prefix tree), and a Graph. Each implementation uses Go generics, includes comprehensive tests and benchmarks, and is documented with godoc-compatible comments.

This project reinforces DSA concepts through implementation and Go best practices through real engineering.

## Learning Objectives

By completing this project, you will:

- Implement classic data structures from scratch without relying on standard library containers
- Use Go generics to create type-safe, reusable data structures
- Write table-driven tests with edge cases and achieve >90% code coverage
- Create benchmarks to measure and optimize performance
- Document code with godoc-compatible comments
- Apply the interface pattern for polymorphic data structure usage

## Prerequisites

- Go 1.22+ installed
- Completed: Go fundamentals lessons (setup, types, functions)
- Completed: Array fundamentals, string manipulation, prefix sum/hashing lessons
- Understanding of: linked lists, hash maps, trees, graphs

## Architecture Overview

```
datastructures/
├── lru/
│   ├── lru.go           # LRU Cache implementation
│   ├── lru_test.go      # Tests
│   └── lru_bench_test.go # Benchmarks
├── trie/
│   ├── trie.go          # Trie implementation
│   ├── trie_test.go     # Tests
│   └── trie_bench_test.go
├── graph/
│   ├── graph.go         # Graph implementation
│   ├── algorithms.go    # BFS, DFS, Dijkstra, topological sort
│   ├── graph_test.go    # Tests
│   └── algorithms_test.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Acceptance Criteria

### LRU Cache

- [ ] **Generic type** — `LRUCache[K comparable, V any]`
- [ ] **O(1) Get** — Returns value and moves to most recently used
- [ ] **O(1) Put** — Inserts or updates; evicts least recently used if at capacity
- [ ] **O(1) Delete** — Removes a specific key
- [ ] **Len/Cap** — Returns current size and maximum capacity
- [ ] **Keys** — Returns all keys in order from most to least recently used
- [ ] **Thread safety** — All operations protected by sync.RWMutex
- [ ] **Implementation** — Doubly linked list + hash map (no container/list)

**Interface:**
```go
type Cache[K comparable, V any] interface {
    Get(key K) (V, bool)
    Put(key K, value V)
    Delete(key K) bool
    Len() int
    Cap() int
    Keys() []K
    Clear()
}
```

### Trie (Prefix Tree)

- [ ] **Insert** — Add a word to the trie
- [ ] **Search** — Check if an exact word exists
- [ ] **StartsWith** — Check if any word starts with a given prefix
- [ ] **Delete** — Remove a word (clean up nodes with no children)
- [ ] **AutoComplete** — Return all words matching a prefix (with limit)
- [ ] **CountWordsWithPrefix** — Count words starting with prefix
- [ ] **Size** — Total number of words stored

**Interface:**
```go
type Trie interface {
    Insert(word string)
    Search(word string) bool
    StartsWith(prefix string) bool
    Delete(word string) bool
    AutoComplete(prefix string, limit int) []string
    CountWordsWithPrefix(prefix string) int
    Size() int
}
```

### Graph

- [ ] **Generic vertices** — `Graph[T comparable]`
- [ ] **Directed and undirected** — Configurable at creation
- [ ] **Weighted edges** — Support for weighted graphs
- [ ] **AddVertex / AddEdge** — Build the graph
- [ ] **RemoveVertex / RemoveEdge** — Modify the graph
- [ ] **Neighbors** — Get adjacent vertices
- [ ] **BFS** — Breadth-first traversal returning visited order
- [ ] **DFS** — Depth-first traversal (iterative, not recursive)
- [ ] **ShortestPath** — Dijkstra's algorithm for weighted graphs
- [ ] **TopologicalSort** — For directed acyclic graphs
- [ ] **HasCycle** — Detect cycles in directed graphs

**Interface:**
```go
type Graph[T comparable] interface {
    AddVertex(v T)
    AddEdge(from, to T, weight float64)
    RemoveVertex(v T) bool
    RemoveEdge(from, to T) bool
    Neighbors(v T) []T
    HasVertex(v T) bool
    HasEdge(from, to T) bool
    Vertices() []T
    Edges() []Edge[T]
    BFS(start T) []T
    DFS(start T) []T
    ShortestPath(from, to T) ([]T, float64)
    TopologicalSort() ([]T, error)
    HasCycle() bool
}
```

### Testing Requirements

- [ ] Table-driven tests for all operations
- [ ] Edge cases: empty structures, single element, duplicate operations
- [ ] Concurrency tests for LRU cache (run with `-race` flag)
- [ ] Benchmark tests for all core operations
- [ ] Code coverage >90% (`go test -cover`)

### Documentation

- [ ] Godoc comments on all exported types and methods
- [ ] Complexity analysis (Big-O) in comments for each method
- [ ] Usage examples in test files or `Example*` functions

## Getting Started

### Step 1: Initialize the Project

```bash
mkdir datastructures && cd datastructures
go mod init github.com/yourusername/datastructures
mkdir lru trie graph
```

### Step 2: Start with the LRU Cache

The LRU cache combines two data structures:
- A **doubly linked list** for maintaining access order (O(1) move to front, remove from back)
- A **hash map** for O(1) key lookup

Build the doubly linked list first, then the cache on top of it.

### Step 3: Build the Trie

Start with Insert and Search, then add StartsWith and AutoComplete. Delete is the trickiest — you need to clean up parent nodes that have no remaining children.

### Step 4: Build the Graph

Start with the adjacency list representation, then add algorithms one at a time. Test each algorithm independently.

### Step 5: Create the Makefile

```makefile
.PHONY: test bench cover lint

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...
```

## Hints and Tips

- **Build the linked list first** — The LRU cache doubly linked list is a great warm-up and is needed for the cache.
- **Use generics** — `[K comparable, V any]` lets your cache work with any key-value types.
- **Test with `-race`** — Go's race detector catches concurrency bugs that tests alone won't find.
- **Benchmark before optimizing** — Use `go test -bench=. -benchmem` to identify actual bottlenecks.
- **Graph adjacency list** — Use `map[T][]Edge[T]` for the adjacency list. It's simpler than a 2D slice for sparse graphs.

## Bonus Challenges

1. **LFU Cache** — Implement a Least Frequently Used cache (harder than LRU)
2. **Persistent Trie** — Add serialization/deserialization to save the trie to disk
3. **Minimum Spanning Tree** — Add Kruskal's or Prim's algorithm to the graph
4. **Visualization** — Generate DOT format output for graphs that can be rendered with Graphviz
5. **Generics Constraint** — Add an `Ordered` constraint for the graph that enables sorted outputs

## Resources

- [Go Generics Tutorial](https://go.dev/doc/tutorial/generics)
- [Go Testing Documentation](https://pkg.go.dev/testing)
- [Go Benchmarking](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [LRU Cache — LeetCode 146](https://leetcode.com/problems/lru-cache/)
- [Implement Trie — LeetCode 208](https://leetcode.com/problems/implement-trie-prefix-tree/)

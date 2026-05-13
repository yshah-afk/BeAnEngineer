# Project: Design a Complete URL Shortener Architecture

## Description

Produce a comprehensive system design document for a URL shortener service (like bit.ly or tinyurl). This project covers the full system design process: requirements gathering, back-of-the-envelope estimation, high-level design, low-level design, API specification, database schema, caching strategy, and scaling plan.

This is a **design-only project** — you will produce architecture documents and diagrams, not code. The focus is on the thinking process and trade-off analysis that system design interviews demand.

## Learning Objectives

By completing this project, you will:

- Apply the RESHADED framework to structure a system design
- Perform capacity estimation (storage, bandwidth, QPS)
- Design APIs with proper REST conventions
- Choose databases and justify the selection with requirements
- Design a caching layer with appropriate strategies and eviction policies
- Plan for horizontal scaling, high availability, and fault tolerance
- Analyze trade-offs and justify design decisions

## Prerequisites

- Completed: Scalability Basics, Caching Strategies, Database Concepts lessons
- Understanding of: HTTP, DNS, load balancing, distributed systems basics

## Architecture Overview

This is what your final design should roughly look like (your version may differ based on your trade-off decisions):

```
┌─────────┐     ┌──────────────┐     ┌──────────────────────┐
│  Client  │────▶│  CDN / Edge  │────▶│    Load Balancer      │
└─────────┘     └──────────────┘     └──────────┬───────────┘
                                                 │
                                    ┌────────────┼────────────┐
                                    ▼            ▼            ▼
                              ┌──────────┐ ┌──────────┐ ┌──────────┐
                              │  API     │ │  API     │ │  API     │
                              │ Server 1 │ │ Server 2 │ │ Server 3 │
                              └────┬─────┘ └────┬─────┘ └────┬─────┘
                                   │             │            │
                              ┌────┴─────────────┴────────────┴────┐
                              │          Redis Cache Cluster        │
                              └────────────────┬───────────────────┘
                                               │
                              ┌────────────────┼───────────────────┐
                              │                │                    │
                         ┌────▼─────┐   ┌─────▼────┐   ┌─────────▼──┐
                         │ DB Shard │   │ DB Shard │   │ Analytics  │
                         │    1     │   │    2     │   │   (Kafka   │
                         └──────────┘   └──────────┘   │ + ClickHouse)
                                                       └────────────┘
```

## Acceptance Criteria

### Deliverable 1: Requirements Document

- [ ] **Functional Requirements**
  - Shorten a URL and return a unique short code
  - Redirect short URL to original URL
  - Optional: custom aliases, expiration dates
  - Optional: analytics (click count, referrers, geo)
- [ ] **Non-Functional Requirements**
  - Availability: 99.99% (< 52 minutes downtime/year)
  - Latency: Redirect in < 50ms (p99)
  - Scale: 100M URLs created/month, 10B redirects/month
  - Durability: URLs should never be lost once created
- [ ] **Constraints and assumptions** documented

### Deliverable 2: Capacity Estimation

- [ ] **Storage estimation** — URLs/month × average URL size × retention period
- [ ] **Bandwidth** — Read and write throughput in MB/s
- [ ] **QPS calculation** — Create QPS, Read QPS (with peak multiplier)
- [ ] **Cache sizing** — 80/20 rule for hot URLs
- [ ] Show all calculations with clear assumptions

### Deliverable 3: API Design

- [ ] REST API specification with endpoints, methods, request/response bodies
- [ ] Proper HTTP status codes for all cases
- [ ] Rate limiting strategy documented
- [ ] Authentication approach for API users

**Expected endpoints:**
```
POST   /api/v1/urls          — Create short URL
GET    /api/v1/urls/:code    — Get URL metadata
DELETE /api/v1/urls/:code    — Delete/deactivate URL
GET    /:code                — Redirect (301 or 302)
GET    /api/v1/urls/:code/stats — Get analytics
```

### Deliverable 4: Short Code Generation Strategy

- [ ] Analyze at least 3 approaches:
  1. Hash-based (MD5/SHA256 + base62 encoding)
  2. Counter-based (distributed counter + base62)
  3. Pre-generated key service (KGS)
- [ ] Collision handling strategy
- [ ] Code length analysis (how many unique codes?)
- [ ] Justify your chosen approach

### Deliverable 5: Database Design

- [ ] Schema diagram with all tables/collections
- [ ] Index strategy with justification
- [ ] Database choice with trade-off analysis
- [ ] Sharding strategy (shard key selection, rebalancing)
- [ ] Replication topology (read replicas, consistency)

### Deliverable 6: Caching Strategy

- [ ] Cache layer design (what to cache, where)
- [ ] Eviction policy (LRU with TTL)
- [ ] Cache warming strategy for popular URLs
- [ ] Invalidation approach for deleted/expired URLs
- [ ] Cache sizing calculation

### Deliverable 7: Scaling Plan

- [ ] How to scale from 1K to 1M to 1B redirects/day
- [ ] Horizontal scaling strategy for API servers
- [ ] Database scaling roadmap (single → replicas → sharding)
- [ ] CDN strategy for global redirect performance
- [ ] Auto-scaling triggers and policies

### Deliverable 8: System Diagram

- [ ] High-level architecture diagram (Mermaid or draw.io)
- [ ] Data flow diagram for create and redirect operations
- [ ] Failure scenarios and recovery strategies

## Getting Started

### Step 1: Apply the RESHADED Framework

1. **R**equirements — Gather functional and non-functional requirements
2. **E**stimation — Back-of-the-envelope calculations
3. **S**torage — Database design and data modeling
4. **H**igh-level Design — Component diagram and data flow
5. **A**PI Design — Endpoint specifications
6. **D**etailed Design — Deep dives into key components
7. **E**valuation — Trade-offs and alternatives considered
8. **D**eployment — Scaling, monitoring, and operational concerns

### Step 2: Start with Estimation

```
Write QPS: 100M URLs/month ÷ (30 × 24 × 3600) ≈ 40 URLs/second
Read QPS: 10B redirects/month ÷ (30 × 24 × 3600) ≈ 3,850 redirects/second
Read:Write ratio = ~100:1

Storage per URL: ~500 bytes (short_code + original_url + metadata)
Storage per month: 100M × 500 bytes = 50 GB/month
Storage per year: 600 GB/year
5-year storage: 3 TB
```

### Step 3: Write the Design Document

Create a Markdown document with all deliverables. Use Mermaid diagrams for architecture and data flow.

### Step 4: Review with Trade-off Analysis

For every major decision, document:
- What alternatives were considered
- Why this option was chosen
- What are the downsides of this choice
- Under what conditions would you change this decision

## Hints and Tips

- **301 vs 302 redirect** — 301 (permanent) is cached by browsers but you lose analytics. 302 (temporary) hits your servers every time, enabling click tracking. Most URL shorteners use 302.
- **Base62 encoding** — `[a-zA-Z0-9]` gives 62 characters. A 7-character code gives 62^7 ≈ 3.5 trillion unique URLs.
- **Read-heavy optimization** — With a 100:1 read-to-write ratio, optimize aggressively for reads: cache everything, use read replicas, consider CDN.
- **The 80/20 rule** — 20% of URLs probably account for 80% of redirects. Cache those hot URLs in Redis.
- **Custom aliases** — These need a separate check for uniqueness and complicate the generation strategy.

## Bonus Challenges

1. **Analytics Deep Dive** — Design the analytics pipeline: event ingestion, real-time counters, and historical aggregation
2. **Abuse Prevention** — Design systems to prevent spam, phishing, and malicious URL shortening
3. **Global Distribution** — Design for a globally distributed service with edge PoPs on every continent
4. **A/B Testing** — Design a feature that lets users create A/B test URLs that distribute traffic between two destinations
5. **Implementation** — Actually build the URL shortener in Go with Redis + PostgreSQL

## Resources

- [System Design Primer: URL Shortener](https://github.com/donnemartin/system-design-primer#design-pastebincom)
- [ByteByteGo: Design a URL Shortener](https://bytebytego.com/courses/system-design-interview/design-a-url-shortener)
- [Designing Data-Intensive Applications](https://dataintensive.net/)
- [RESHADED Framework](https://github.com/ashishps1/awesome-system-design-resources)
- [Base62 Encoding Explained](https://en.wikipedia.org/wiki/Base62)

# bqsql Implementation Roadmap

This document outlines the strategic phases for evolving `bqsql` from a basic CLI tool into a robust, API-first BigQuery management utility and interactive terminal.

---

## Phase 1: Performance & Core Stability
*Focus on optimizing data handling and reducing latency for frequent operations.*

| Task | Description | Issue Link | Complexity | Length | Priority | Dependencies |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Conventional Commits & Changelog** | Implement automated CHANGELOG.md generation based on Conventional Commits. | [#37](https://github.com/rantolin/bqsql/issues/37) | Easy | Short | **P0** | None |
| **Result Pagination & Streaming** | Transition from in-memory result handling to a paginated/streaming approach. | [#4](https://github.com/rantolin/bqsql/issues/4) | Hard | Large | **P0** | None |
| **Schema Caching** | Implement local caching of table schemas for `describe` and `query` operations. | [#23](https://github.com/rantolin/bqsql/issues/23) | Medium | Medium | **P1** | None |

---

## Phase 2: CLI Usability & UX
*Refining how users interact with the command line and its outputs.*

| Task | Description | Issue Link | Complexity | Length | Priority | Dependencies |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Different Output Formats** | Add `--format` flag to support `csv`, `json`, and `jsonl` outputs. | [#3](https://github.com/rantolin/bqsql/issues/3) | Medium | Medium | **P0** | None |
| **Revise Usage Messages** | Audit and update all `Short` and `Long` help messages for consistency. | [#7](https://github.com/rantolin/bqsql/issues/7) | Easy | Short | **P0** | None |
| **Shell Auto-completion** | Implement Bash, Zsh, and Fish completion for BigQuery resources. | [#24](https://github.com/rantolin/bqsql/issues/24) | Medium | Short | **P1** | None |

---

## Phase 3: The Management Suite (CRUD & Security)
*Transforming bqsql into a fully operational BigQuery management utility.*

| Task | Description | Issue Link | Complexity | Length | Priority | Dependencies |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Resource CRUD Operations** | Implement full lifecycle management (Create, Update, Delete) for Datasets and Tables. | [#27](https://github.com/rantolin/bqsql/issues/27) | Medium | Large | **P1** | None |
| **IAM & Security Management** | Add commands to manage dataset/table permissions (IAM Policies). | [#31](https://github.com/rantolin/bqsql/issues/31) | Medium | Medium | **P2** | None |
| **Dry Run Support** | Add a `--dry-run` flag to queries to show estimated bytes processed. | [#22](https://github.com/rantolin/bqsql/issues/22) | Easy | Short | **P1** | None |
| **Job & Task Management** | Implement `bqsql jobs` to list, inspect, and cancel active BigQuery jobs. | [#28](https://github.com/rantolin/bqsql/issues/28) | Medium | Medium | **P1** | None |

---

## Phase 4: Architecture & Internal SDK Refactor
*Preparing the codebase for API-first integration and external consumption.*

| Task | Description | Issue Link | Complexity | Length | Priority | Dependencies |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Library Decoupling** | Refactor core logic into an internal library (`pkg/bqcore`) for reuse. | [#25](https://github.com/rantolin/bqsql/issues/25) | Hard | Large | **P0** | Phase 1 & 3 |
| **Structured Logging & Errors** | Implement structured logging (`slog`) and custom error types for the SDK. | [#1](https://github.com/rantolin/bqsql/issues/1) | Medium | Medium | **P1** | None |
| **Testing Suite** | Establish a comprehensive testing suite (Unit, Integration) for the new SDK. | [#26](https://github.com/rantolin/bqsql/issues/26) | Medium | Large | **P0** | Library Decoupling |

---

## Phase 5: The bqsql API Layer
*Exposing the bqsql functionality as a service.*

| Task | Description | Issue Link | Complexity | Length | Priority | Dependencies |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **REST/gRPC API Server** | Implement a server that wraps the internal `bqcore` SDK for remote execution. | [#32](https://github.com/rantolin/bqsql/issues/32) | Hard | XLarge | **P2** | Phase 4 |
| **Authentication Middleware** | Add OAuth2/Service Account credential handling for the API layer. | [#33](https://github.com/rantolin/bqsql/issues/33) | Medium | Medium | **P2** | API Server |

---

## Phase 6: Advanced Interactive Terminal (The "SnowSQL/psql" Experience)
*Creating a world-class interactive shell for BigQuery.*

| Task | Description | Issue Link | Complexity | Length | Priority | Dependencies |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Interactive REPL** | Create a dedicated shell mode (`bqsql shell`) with persistent connections. | [#29](https://github.com/rantolin/bqsql/issues/29) | Hard | Large | **P1** | Phase 4 |
| **Rich Shell Features** | Implement command history, multiline input, and syntax highlighting. | [#34](https://github.com/rantolin/bqsql/issues/34) | Hard | XLarge | **P2** | Interactive REPL |
| **Live Autocomplete** | Add context-aware autocomplete for SQL keywords and resource names. | [#35](https://github.com/rantolin/bqsql/issues/35) | Hard | XLarge | **P2** | Rich Shell Features |

---

## Phase 7: Developer Experience (DX) & CI/CD
*Standardizing the release and maintenance lifecycle.*

| Task | Description | Issue Link | Complexity | Length | Priority | Dependencies |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **CI/CD Pipeline** | Set up GitHub Actions for automated linting, testing, and releases. | [#9](https://github.com/rantolin/bqsql/issues/9) | Medium | Medium | **P0** | Phase 4 |
| **Release Management** | Implement automated versioning and release generation (GoReleaser). | [#30](https://github.com/rantolin/bqsql/issues/30) | Easy | Short | **P1** | CI/CD Pipeline |

---

## Recommended Implementation Sequence

This sequence prioritizes foundational stability and quick wins first, followed by structural refactoring to unblock the API and interactive shell.

### Step 1: Quick Wins & Stability (Current Focus)
1. **Conventional Commits & Changelog** (P0) - [#37](https://github.com/rantolin/bqsql/issues/37)
2. **Revise Usage Messages** (P0) - [#7](https://github.com/rantolin/bqsql/issues/7)
3. **Different Output Formats** (P0) - [#3](https://github.com/rantolin/bqsql/issues/3)
4. **Result Pagination & Streaming** (P0) - [#4](https://github.com/rantolin/bqsql/issues/4)

### Step 2: Strategic CLI Enhancements
5. **Dry Run Support** (P1) - [#22](https://github.com/rantolin/bqsql/issues/22)
6. **Schema Caching** (P1) - [#23](https://github.com/rantolin/bqsql/issues/23)
7. **Shell Auto-completion** (P1) - [#24](https://github.com/rantolin/bqsql/issues/24)

### Step 3: Structural Refactoring (The SDK Pivot)
8. **Library Decoupling** (P0) - [#25](https://github.com/rantolin/bqsql/issues/25)
9. **Testing Suite** (P0) - [#26](https://github.com/rantolin/bqsql/issues/26)
10. **CI/CD Pipeline** (P0) - [#9](https://github.com/rantolin/bqsql/issues/9)
11. **Structured Logging & Errors** (P1) - [#1](https://github.com/rantolin/bqsql/issues/1)

### Step 4: Full Management & REPL
12. **Resource CRUD Operations** (P1) - [#27](https://github.com/rantolin/bqsql/issues/27)
13. **Job & Task Management** (P1) - [#28](https://github.com/rantolin/bqsql/issues/28)
14. **Interactive REPL** (P1) - [#29](https://github.com/rantolin/bqsql/issues/29)
15. **Release Management** (P1) - [#30](https://github.com/rantolin/bqsql/issues/30)

### Step 5: Advanced Ecosystem (Visionary)
16. **IAM & Security Management** (P2) - [#31](https://github.com/rantolin/bqsql/issues/31)
17. **REST/gRPC API Server** (P2) - [#32](https://github.com/rantolin/bqsql/issues/32)
18. **Authentication Middleware** (P2) - [#33](https://github.com/rantolin/bqsql/issues/33)
19. **Rich Shell Features** (P2) - [#34](https://github.com/rantolin/bqsql/issues/34)
20. **Live Autocomplete** (P2) - [#35](https://github.com/rantolin/bqsql/issues/35)

---

## Legend
- **Priority**:
    - **P0**: **Critical / Foundational.** Essential for core stability, unblocking other tasks, or fixing major UX issues.
    - **P1**: **High / Strategic.** Core features that expand the tool's utility and align with the primary roadmap.
    - **P2**: **Medium / Visionary.** Advanced features, polish, or expansions for niche use cases.
- **Complexity**: Easy, Medium, Hard.
- **Length**: Short (1-2 days), Medium (3-5 days), Large (1-2 weeks), XLarge (2+ weeks).

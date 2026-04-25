# PRP: Testing Suite (Phase 1: Utilities)

## 🎯 Goal
- **Objective:** Establish a robust testing foundation starting with the `utils/` package.
- **Why:** To enable safe refactoring and decoupling of the core logic from the CLI commands, ensuring high code quality and reliability.
- **Success Criteria:**
  - [ ] 100% test coverage for "pure" utility functions in `utils/formats.go`.
  - [ ] Refactored `CalculateRowWidths` to be testable without a live BigQuery connection.
  - [ ] All tests follow the project's "Table-Driven Tests" standard.
  - [ ] Zero regressions in `cmd/query.go` and `cmd/head.go` after refactoring.

## 📚 Context
```yaml
docs:
  - https://pkg.go.dev/testing
  - https://go.dev/wiki/TableDrivenTests
  - https://cloud.google.com/go/bigquery
codebase:
  examples:
    - file: N/A (Initial test implementation)
  targets:
    - file: utils/formats.go
    - file: utils/formats_test.go
  context:
    - path: cmd/
    - file: cmd/query.go
gotchas:
  - bigquery.RowIterator is a concrete struct, not an interface. To mock it, we need to abstract the result iteration.
  - Calculation of widths currently iterates the result set once, which might be expensive for large sets (P0 task #4 addresses this, but we must maintain compatibility).
```

## 🗺️ Implementation Blueprint
### Data Models / Schema
- Introduce an interface `RowProvider` or similar to abstract `bigquery.RowIterator`.

### Integration Points
- `utils.CalculateRowWidths`: Refactor to accept an interface or a slice of rows.
- `utils.PrintFormatedRow`: Ensure it remains pure and easy to test.

### Pseudocode / Logic
### Pseudocode / Logic
```go
// Option B: Comprehensive Interface
type RowProvider interface {
    Next(dst interface{}) error
    Schema() bigquery.Schema
}

// Wrapper for the real bigquery.RowIterator
type BigQueryRowProvider struct {
    it *bigquery.RowIterator
}
func (b *BigQueryRowProvider) Next(dst interface{}) error { return b.it.Next(dst) }
func (b *BigQueryRowProvider) Schema() bigquery.Schema    { return b.it.Schema }

// Refactored function signature
func CalculateRowWidths(provider RowProvider) ([]int, error) {
    schema := provider.Schema()
    widths := make([]int, len(schema))
    // ... iteration logic using provider.Next() ...
}
```

### Sequential Tasks
- [ ] **Task 1: Pure Function Tests.** Implement tests for `max` and `PrintFormatedRow` in `utils/formats_test.go`.
- [ ] **Task 2: TDD for RowProvider Contract.**
    - Define `RowProvider` interface in `utils/formats.go`.
    - Create a "Contract Test" in `utils/formats_test.go` that verifies any `RowProvider` (Mock or Real) returns the expected schema and rows.
    - Implement `MockRowProvider` and verify it passes the contract test.
    - Implement `BigQueryRowProvider` (wrapper) and verify it satisfies the interface.
- [ ] **Task 3: TDD Step 1 - Failing Test for CalculateRowWidths (Red).**
    - Implement a table-driven test for `CalculateRowWidths` using the now-verified `MockRowProvider`.
    - Verify that the test fails to compile because `CalculateRowWidths` still expects `*bigquery.RowIterator`.
- [ ] **Task 4: TDD Step 2 - Refactor CalculateRowWidths (Green).**
    - Change `CalculateRowWidths` signature to accept `RowProvider`.
    - Update logic to get the schema directly from the provider.
    - Verify all tests in `utils/formats_test.go` pass.
- [ ] **Task 5: Integration & Cleanup (Refactor).**
    - Update `cmd/query.go` and `cmd/head.go` to wrap their iterators in `BigQueryRowProvider`.
    - Perform a final validation of the CLI tool.


## ✅ Validation Loop
### Level 1: Structure & Syntax
```bash
go vet ./utils/...
```

### Level 2: Functional Correctness
```bash
go test -v ./utils/...
```

### Level 3: Integration/E2E
```bash
go build -o bqsql main.go && ./bqsql --help
```

## 🚫 Anti-Patterns
- **No Global State:** Do not use global variables for mocks.
- **No Over-Engineering:** Keep the interface minimal (only what's needed for `CalculateRowWidths`).
- **No Test-Only Exports:** Use the `_test` package or keep internal logic testable without exporting everything.

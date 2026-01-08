# AGENTS.md

Guidelines and best practices for agentic coding agents (and humans!) working on the vid-dl Go project.

## Build, Lint & Test Commands

### Build Commands
```bash
# Build the CLI
go build ./cmd/vid-dl

# Run CLI directly
go run ./cmd/vid-dl [URL]
```

### Test Commands
```bash
# Run all tests in all packages
go test ./...

# Run all tests for a single package
go test ./internal/input
# or
go test ./internal/downloader
# or
go test ./internal/tui

# Run a single test function in a package
# (replace TestFunctionName with your test function's name)
go test -run TestFunctionName ./internal/downloader

# Verbose output for all tests
go test -v ./...

# Run with coverage reporting
go test -cover ./...
```

### Lint & Static Analysis Commands
```bash
# Go formatting
go fmt ./...

# Go vet (static analysis)
go vet ./...

# Install and run golangci-lint (requires network access only once)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run
```

### Dependency Management
```bash
# Add or tidy dependencies
go mod tidy
# Add a specific dependency
go get github.com/example/package
```

---

## Project Architecture

- `cmd/vid-dl/` — Main CLI entrypoint
- `internal/downloader/` — Video downloading logic (yt-dlp integration)
- `internal/input/` — Parsing/validation for user/URL input
- `internal/tui/` — Terminal UI powered by Bubble Tea/huh

**External dependencies:**
- `yt-dlp`, `ffmpeg` — Must be on `$PATH` at runtime
- Go modules: `github.com/charmbracelet/huh`, `github.com/charmbracelet/bubbletea`

---

## Code Style Guidelines

### Import Grouping
Use three groups, in this order:
1. Standard library (`fmt`, `os`, ...)
2. Internal packages (`github.com/cristhianflo/vid-dl/internal/…`)
3. Third-party (`github.com/charmbracelet/huh`)

```go
import (
    "fmt"
    "os"

    "github.com/cristhianflo/vid-dl/internal/input"
    "github.com/cristhianflo/vid-dl/internal/tui"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/huh"
)
```

### Naming Conventions
- Packages: lowercase, one word (`input`, `tui`, `downloader`)
- Types: `MixedCase` (e.g. `VideoSource`, `Downloader`)
- Functions: `ExportedFunction()` and `unexportedFunction()`
- Variables: use `mixedCase` (e.g. `videoURL`, `err`)
- Constants: grouped in `const ()`, named in `CamelCase`

### Error Handling
- Check and handle all errors immediately; never ignore them!
- Always wrap errors with context unless returning them directly
  ```go
  value, err := someFunc()
  if err != nil {
      return fmt.Errorf("problem doing X: %w", err)
  }
  ```
- Return early on error
- Never use panics for normal error handling

### Function & Struct Layout
- Group required fields at the top of structs, optional fields last
- Methods grouped on their type
- Functions limited in size, single responsibility preferred

### Interface & Factory Patterns
- Prefer small, focused interfaces:
  ```go
  type Downloader interface {
      GetFormats() (*Video, error)
      DownloadVideo(format *Format) error
  }
  ```
- Implement via factory, e.g.:
  ```go
  func NewDownloader(source *input.VideoSource) (Downloader, error) {...}
  ```

### Formatting & Readability
- Use `go fmt` to maintain formatting everywhere
- Use linting/analysis before commit
- Avoid long lines; break up logic for readability

---

## Test Writing Patterns

Use table-driven tests for coverage:
```go
func TestIsValidURL(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantErr bool
    }{
        {"valid URL", "https://example.com", false},
        {"invalid", "notaurl", true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := IsValidURL(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("got error %v, wantErr=%v", err, tt.wantErr)
            }
        })
    }
}
```

---

## File Structure and Naming
- Use `snake_case.go` for source
- Tests in `*_test.go`
- Names must be descriptive but concise

---

## Workflow for Agentic Coding
1. Use provided build/test/lint commands on each change
2. Handle all errors as described above
3. Make one clear, self-contained change at a time
4. Never commit secrets or configuration files
5. Update imports if you restructure packages
6. Always keep code idiomatic Go; run `go fmt` and lint
7. Ensure all modified code is covered by tests or manual build+run!

---

## Patterns to Avoid
- Ignoring errors
- Long or global-variable-based code
- "God" functions: break up large logic
- Hardcoded filesystem paths or URLs
- Skipping nil pointer checks or unsafe casts
- Mixing UI, parsing, network, or download logic in the same function

---

## Bubble Tea TUI Guidelines
- Isolate business logic from UI
- Use `Init`, `Update`, `View` per Bubble Tea conventions
- Reuse simple styles & components
- Use explicit messages for async events
- Keep TUI logic in `internal/tui/`

---

## External Tool Integration
- Always wrap `exec.Command` with clear error handling
- Prefer context on exec failures:
   ```go
   cmd := exec.Command("yt-dlp", args...)
   out, err := cmd.CombinedOutput()
   if err != nil {
       return fmt.Errorf("yt-dlp failed: %w\noutput: %s", err, string(out))
   }
   ```

---

# End of instructions for AGENTS.md

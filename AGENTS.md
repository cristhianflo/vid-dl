# AGENTS.md

Guidelines for agentic coding agents working on vid-dl project.

## Build/Lint/Test Commands

### Build Commands
```bash
go build ./cmd/vid-dl
go run ./cmd/vid-dl [URL]
```

### Test Commands
```bash
# All tests
go test ./...

# Package tests
go test ./internal/input
go test ./internal/downloader  
go test ./internal/tui

# Single test
go test -run TestFunctionName ./internal/input
go test -v ./...
go test -cover ./...
```

### Lint Commands
```bash
go fmt ./...
go vet ./...
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run
```

### Dependencies
```bash
go mod tidy
go get github.com/example/package
```

## Architecture

- `cmd/vid-dl/` - Main entry point
- `internal/downloader/` - Video downloaders  
- `internal/input/` - URL parsing/validation
- `internal/tui/` - Terminal UI components

External tools: `yt-dlp`, `ffmpeg` (must be in $PATH)
Go deps: `github.com/charmbracelet/huh`, `github.com/charmbracelet/bubbletea`

## Code Style

### Import Organization
```go
import (
    "fmt"
    "os"                     // Standard library first
    
    "cristhianflo/vid-dl/internal/input"
    "cristhianflo/vid-dl/internal/tui"     // Internal next
    
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/huh"           // Third-party last
)
```

### Naming Conventions
- Packages: `downloader`, `input`, `tui` (single lowercase)
- Functions: `NewDownloader()` (exported), `makeFormat()` (unexported)
- Types: `VideoSource`, `Downloader` (MixedCase)
- Variables: `videoURL` (local), `YoutubeVideo` (exported)
- Constants: grouped with `const ()`

### Error Handling
```go
result, err := someFunction()
if err != nil {
    return nil, fmt.Errorf("context: %w", err)
}
```
Handle errors immediately, wrap with context, return early.

### Struct Organization
```go
type VideoSource struct {
    OriginalURL *url.URL    // Required fields first
    VideoID     string
    Type        VideoType
    
    // Optional fields last
    Metadata    map[string]string
}
```

### Interface Design
```go
type Downloader interface {
    GetFormats(source *input.VideoSource) (*Video, error)
}

type YtdlpDownloader struct{}
func (d *YtdlpDownloader) GetFormats(...) { ... }
```

Small interfaces, clear implementations, factory pattern for creation.

### External Tool Integration
```go
func runYtdlp(args ...string) ([]byte, error) {
    cmd := exec.Command("yt-dlp", args...)
    output, err := cmd.Output()
    if err != nil {
        return nil, fmt.Errorf("failed to run yt-dlp: %w", err)
    }
    return output, nil
}
```

## Development Workflow

### Before Committing
1. `go fmt ./...`
2. `go vet ./...` 
3. `go test ./...` (if tests exist)
4. `go mod tidy`

### Adding Video Sources
1. Add `VideoType` constant in `input/input.go`
2. Create parser function in `input/input.go`
3. Create downloader implementation
4. Register in factory function

### Adding TUI Components
1. Create in `tui/*.go` file
2. Follow Bubble Tea: `Init`, `Update`, `View`
3. Keep focused and reusable

## Testing Patterns
```go
func TestIsValidURL(t *testing.T) {
    tests := []struct {
        name     string
        input    string  
        wantErr  bool
    }{
        {"valid URL", "https://example.com", false},
        {"invalid URL", "not-a-url", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := IsValidURL(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("IsValidURL() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

## File Naming
- `snake_case.go` for all files
- `*_test.go` for tests
- Descriptive but concise names

## Patterns to Avoid
- Don't ignore errors
- Don't use global variables
- Don't create god functions
- Don't hardcode URLs/paths
- Don't ignore nil pointer checks
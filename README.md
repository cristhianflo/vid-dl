# vid-dl

A powerful and user-friendly CLI tool to download videos from your terminal with modern Go and TUI technologies. Download from YouTube (and more in the future!) using a simple, interactive interface backed by ffmpeg and yt-dlp. Designed for extensibility, effortless usage, and robust video/audio processing.

## Installation

### Build (requires Go 1.22+)
```
git clone https://github.com/cristhianflo/vid-dl.git
cd vid-dl
go build -o vid-dl ./cmd/vid-dl/
```

## Technologies

| Category         | Technologies                              |
| ---------------- | ----------------------------------------- |
| CLI/Backend      | Go 1.22+                                  |
| Terminal UI      | Bubble Tea, Huh (interactive CLI)         |
| Video Processing | ffmpeg (external), yt-dlp (external)      |
| Tests/Linting    | `go test`,`go vet`                           |
| Infrastructure   | Local, binary build, cross-platform ready |

## Features

- **Interactive CLI:** Paste a video url and pick a format option directly in the terminal.
- **Multiple Sources:** Currently supports YouTube; extensible to more video platforms.
- **Robust Downloading:** Integrates yt-dlp for video fetching and ffmpeg for converting audio/video.
- **Reliable Validation:** Validates URLs, video IDs, and file paths before download.
- **Cross-platform:** Runs on Linux, Mac, Windows (x86_64/ARM64) with minimal setup.
- **Testing & Linting:** Standard Go tests and vetting for code reliability.

## What I Learned

Developing vid-dl provided an opportunity to tackle several new challenges and integrate a rich set of CLI and video processing tools.

### CLI/TUI Apps with Go

Building a smooth interactive terminal interface using Bubble Tea & Huh was both challenging and rewarding. Learned how to structure stateful, interactive logic in a Go CLI ecosystem.

### External Tool Integration

It’s surprisingly seamless to call and coordinate powerful tools like `yt-dlp` and `ffmpeg` from Go, making advanced media processing available in simple workflows.

### Validation and Error Handling

Good CLIs must give strong feedback for invalid inputs. I’ve reaffirmed the importance of robust validation, careful error messaging, and early failure handling.

## Directory Structure

- `cmd/vid-dl/` — Main CLI entry point.
- `internal/downloader/` — Download logic, yt-dlp wrappers.
- `internal/input/` — Validation/parsing of URLs and user input.
- `internal/tui/` — Interactive terminal UI with Bubble Tea.
- `README.md` — You are here!

## Testing

To run all tests and build checks:

- Run all tests: `go test ./...`
- Format and vet: `go fmt ./...`, `go vet ./...`

See AGENTS.md for build/test/lint details.

## Demo

<!-- Add CLI demo gifs, screenshots, or video links here -->

## Roadmap/Future Enhancements

- Support for additional video sources: Vimeo, TikTok, etc.
- Playlists & channel downloads.
- Download queueing, resume, more output formats.
- Enhanced error feedback and logging.
- Configurable output templates.

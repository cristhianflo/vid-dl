# vid-dl

A powerful and user-friendly CLI tool for downloading videos.

## Project Scope & Features

- **Purpose:** Download videos from YouTube via command-line.
- **Functionalities:**
  - Initial focus on downloading videos from YouTube.
  - Uses [ffmpeg](https://www.ffmpeg.org/) and [yt-dlp](https://github.com/yt-dlp/yt-dlp) as external executables.
  - Will provide an easy-to-use interactive UI using [gum](https://github.com/charmbracelet/gum).
  - Designed for extensibility (easy to add more sources or processing steps later).

## Dependencies

- **ffmpeg** — Required for video/audio processing.
- **yt-dlp** — Powerful, flexible YouTube downloader backend.
- **gum** — CLI toolkit for TUI/interactive CLI prompts.

Ensure these dependencies are installed (outside of Go modules) and available in your `$PATH` prior to using the tool.

## Structure

- **cmd/vid-dl/main.go** — CLI entrypoint
- **pkg/** — reusable Go packages (business logic)
- **internal/** — private/internal project code

----

### Future Enhancements
- Support for more video sources/platforms
- Customizable output (format, quality)
- Download playlists or entire channels

---

*This README will be updated as features are implemented.*

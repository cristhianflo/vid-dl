package tui

import "fmt"

func humanFileSize(size int64) string {
	if size < 0 {
		return fmt.Sprintf("-%s", humanFileSize(-size))
	}
	if size == 0 {
		return "0 B"
	}

	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}

	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	// KiB, MiB, GiB, TiB, PiB, EiB
	suffix := "KMGTPE"[exp]

	return fmt.Sprintf("%.1f %ciB", float64(size)/float64(div), suffix)
}

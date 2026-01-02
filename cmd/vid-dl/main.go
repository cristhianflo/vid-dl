package main

import (
	"fmt"
	"os"
	"os/exec"

	"cristhianflo/vid-dl/internal/input"
	"cristhianflo/vid-dl/internal/ytdlp"
)

func main() {
	var videoURL string

	if len(os.Args) > 1 {
		videoURL = os.Args[1]
	} else {
		cmd := exec.Command("gum", "input", "--placeholder", "Enter the video URL...")
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error running gum: %v\n", err)
			os.Exit(1)
		}
		videoURL = string(output)
	}

	parsedURL, err := input.ParseYouTubeURL(videoURL)
	if err != nil {
		fmt.Printf("Invalid YouTube URL: %v\n", err)
		os.Exit(1)
	}

	ytdlpResult, err := ytdlp.GetVideoInfo(parsedURL)
	if err != nil {
		fmt.Printf("Error getting ytdlp result: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Valid YouTube URL\nResult: %v\n", ytdlpResult)
}

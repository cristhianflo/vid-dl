package main

import (
	"fmt"
	"os"

	"cristhianflo/vid-dl/internal/downloader"
	"cristhianflo/vid-dl/internal/input"
	"cristhianflo/vid-dl/internal/tui"
)

func main() {
	videoURL, err := tui.GetVideoURL(os.Args)
	if err != nil {
		fmt.Printf("Error getting video URL: %v\n", err)
		os.Exit(1)
	}

	videoSource, err := input.GetVideoSource(videoURL)
	if err != nil {
		fmt.Printf("Error getting video source: %v\n", err)
		os.Exit(1)
	}

	videoDownloader, err := downloader.NewDownloader(videoSource)
	if err != nil {
		fmt.Printf("Error creating video downloader: %v\n", err)
		os.Exit(1)
	}

	model, err := tui.NewModel(videoDownloader)
	if err != nil {
		fmt.Printf("Error getting video formats: %v\n", err)
		os.Exit(1)
	}

	p := tui.NewTui(*model)
	_, err = p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

package tui

import (
	"github.com/cristhianflo/vid-dl/internal/input"

	"github.com/charmbracelet/huh"
)

func GetVideoURL(args []string) (*input.VideoURL, error) {
	if len(args) > 1 {
		err := input.IsEmpty(args[1])
		if err != nil {
			return nil, err
		}
		u, err := input.IsValidURL(args[1])
		if err != nil {
			return nil, err
		}

		return &input.VideoURL{
			OriginalURL: u,
		}, nil
	}

	var videoURL string
	err := huh.NewInput().
		Title("Video URL").
		Value(&videoURL).
		Placeholder("e.g: https://youtube.com/watch?v=123...").
		Validate(input.IsEmpty).
		Run()

	u, err := input.IsValidURL(videoURL)
	if err != nil {
		return nil, err
	}

	return &input.VideoURL{
		OriginalURL: u,
	}, nil
}

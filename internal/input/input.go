package input

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type VideoType int

const (
	YoutubeVideo VideoType = iota
)

type VideoSource struct {
	OriginalURL *url.URL
	VideoID     string
	Type        VideoType
}

func ParseYouTubeURL(input string) (*VideoSource, error) {
	u, err := IsValidURL(input)
	if err != nil {
		return nil, errors.New("Input is not a valid URL.")
	}

	var videoSource VideoSource
	videoSource.OriginalURL = u
	videoSource.Type = YoutubeVideo
	host := strings.TrimPrefix(u.Host, "www.")

	switch host {
	case "youtube.com", "m.youtube.com":
		if u.Path == "/watch" {
			videoSource.VideoID = u.Query().Get("v")
		} else if videoID, ok := strings.CutPrefix(u.Path, "/shorts/"); ok {
			videoSource.VideoID = videoID
		}
	case "youtu.be":
		videoSource.VideoID = strings.Trim(u.Path, "/")
	}

	var validID = regexp.MustCompile(`^[\w-]{11}$`)
	if !validID.MatchString(videoSource.VideoID) {
		return nil, errors.New("invalid or missing YouTube video ID in URL")
	}

	return &videoSource, nil
}

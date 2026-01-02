package input

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type ParsedURL struct {
	OriginalURL *url.URL
	VideoID     string
}

func ParseYouTubeURL(userInput string) (*ParsedURL, error) {
	trimmed := strings.TrimSpace(userInput)
	if trimmed == "" {
		return nil, errors.New("input is empty")
	}
	u, err := url.Parse(trimmed)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, errors.New("invalid URL format: failed to parse")
	}

	var parsedURL ParsedURL
	parsedURL.OriginalURL = u
	host := strings.TrimPrefix(u.Host, "www.")

	switch host {
	case "youtube.com", "m.youtube.com":
		if u.Path == "/watch" {
			parsedURL.VideoID = u.Query().Get("v")
		} else if videoID, ok := strings.CutPrefix(u.Path, "/shorts/"); ok {
			parsedURL.VideoID = videoID
		}
	case "youtu.be":
		parsedURL.VideoID = strings.Trim(u.Path, "/")
	}

	var validID = regexp.MustCompile(`^[\w-]{11}$`)
	if !validID.MatchString(parsedURL.VideoID) {
		return nil, errors.New("invalid or missing YouTube video ID in URL")
	}

	return &parsedURL, nil
}

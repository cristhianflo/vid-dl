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

type VideoURL struct {
	OriginalURL *url.URL
}

type VideoSource struct {
	*VideoURL
	OriginalURL *url.URL
	VideoID     string
	Type        VideoType
}

func GetVideoSource(videoURL *VideoURL) (*VideoSource, error) {
	var videoSource VideoSource
	u := videoURL.OriginalURL
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

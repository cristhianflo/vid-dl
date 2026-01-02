package ytdlp

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os/exec"

	"cristhianflo/vid-dl/internal/input"
)

type FileExtension string

const (
	Mhtml FileExtension = "mhtml"
	M4a   FileExtension = "m4a"
	Webm  FileExtension = "webm"
	Mp4   FileExtension = "mp4"
)

type YtdlpFormat struct {
	ID          string
	Ext         FileExtension
	Resolution  string
	Fps         float32
	Filesize    int64
	DownloadURL *url.URL
}

type YtdlpResult struct {
	VideoID    string
	VideoTitle string
	Formats    []YtdlpFormat
}

type ytDLPJson struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Formats []struct {
		FormatID   string   `json:"format_id"`
		Ext        string   `json:"ext"`
		Resolution *string  `json:"resolution"`
		Fps        *float32 `json:"fps"`
		Filesize   *int64   `json:"filesize_approx"`
		URL        string   `json:"url"`
	} `json:"formats"`
}

func GetVideoInfo(parsed *input.ParsedURL) (*YtdlpResult, error) {
	cmd := exec.Command("yt-dlp", "-j", parsed.OriginalURL.String())
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run yt-dlp: %w", err)
	}

	var yd ytDLPJson
	err = json.Unmarshal(output, &yd)
	if err != nil {
		return nil, fmt.Errorf("failed to decode yt-dlp output: %w", err)
	}

	result := &YtdlpResult{
		VideoID:    yd.ID,
		VideoTitle: yd.Title,
	}

	for _, f := range yd.Formats {
		format := YtdlpFormat{
			ID:         f.FormatID,
			Ext:        FileExtension(f.Ext),
			Resolution: "",
			Fps:        0,
			Filesize:   0,
		}
		if f.Resolution != nil {
			format.Resolution = *f.Resolution
		}
		if f.Fps != nil {
			format.Fps = *f.Fps
		}
		if f.Filesize != nil {
			format.Filesize = *f.Filesize
		}
		if parsedURL, err := url.Parse(f.URL); err != nil {
			format.DownloadURL = parsedURL
		}
		result.Formats = append(result.Formats, format)
	}
	return result, nil
}

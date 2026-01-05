package downloader

import (
	"cristhianflo/vid-dl/internal/input"
	"encoding/json"
	"fmt"
	"os/exec"
)

type YtdlpDownloader struct {
	source *input.VideoSource
}

type ytdlpFormat struct {
	FormatID   string   `json:"format_id"`
	Ext        string   `json:"ext"`
	Resolution *string  `json:"resolution"`
	Fps        *float32 `json:"fps"`
	Filesize   *int64   `json:"filesize_approx"`
	ACodec     string   `json:"acodec"`
	VCodec     string   `json:"vcodec"`
}

type ytdlpResponse struct {
	ID      string        `json:"id"`
	Title   string        `json:"title"`
	Formats []ytdlpFormat `json:"formats"`
}

func makeVideoFormat(f ytdlpFormat) Format {
	format := Format{
		ID:         f.FormatID + "+bestaudio",
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

	return format
}

func makeAudioFormat(f ytdlpFormat) Format {
	filesize := int64(0)
	if f.Filesize != nil {
		filesize = *f.Filesize
	}

	return Format{
		ID:         "bestaudio",
		Ext:        Mp3,
		Resolution: "audio only",
		Fps:        0,
		Filesize:   filesize,
	}
}

func (d *YtdlpDownloader) GetFormats() (*Video, error) {
	cmd := exec.Command("yt-dlp", "--no-playlist", "-j", "--format-sort=resolution,ext,tbr", d.source.OriginalURL.String())
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run yt-dlp: %w", err)
	}

	var resp ytdlpResponse
	err = json.Unmarshal(output, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode yt-dlp output: %w", err)
	}

	result := &Video{
		ID:    resp.ID,
		Title: resp.Title,
	}

	bestAudioIdx := -1
	videoFormats := make(map[string]int)

	for i, f := range resp.Formats {

		if f.Ext == string(Mhtml) || f.Resolution == nil {
			continue
		}

		if f.VCodec == "none" {
			if bestAudioIdx == -1 {
				bestAudioIdx = i
			} else if f.Filesize != nil && resp.Formats[bestAudioIdx].Filesize != nil && *f.Filesize > *resp.Formats[bestAudioIdx].Filesize {
				bestAudioIdx = i
			}
			continue
		}

		res := *f.Resolution

		if existingIdx, exists := videoFormats[res]; !exists {
			videoFormats[res] = i
		} else {
			if f.Filesize != nil && resp.Formats[existingIdx].Filesize != nil && *f.Filesize > *resp.Formats[existingIdx].Filesize {
				videoFormats[res] = i
			}
		}
	}

	var finalFormats []Format

	if bestAudioIdx != -1 {
		bestAudioFormat := makeAudioFormat(resp.Formats[bestAudioIdx])
		finalFormats = append(finalFormats, bestAudioFormat)
	}

	for _, idx := range videoFormats {
		f := resp.Formats[idx]

		finalFormats = append(finalFormats, makeVideoFormat(f))
	}

	result.Formats = finalFormats
	return result, nil
}

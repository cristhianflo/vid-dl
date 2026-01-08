package downloader

import (
	"errors"
	"github.com/cristhianflo/vid-dl/internal/input"
)

type FileExtension string

const (
	Mhtml FileExtension = "mhtml"
	M4a   FileExtension = "m4a"
	Webm  FileExtension = "webm"
	Mp3   FileExtension = "mp3"
	Mp4   FileExtension = "mp4"
)

type Video struct {
	ID      string
	Title   string
	Formats []Format
}

type Format struct {
	ID         string
	Ext        FileExtension
	Resolution string
	Fps        float32
	Filesize   int64
}

type Downloader interface {
	GetFormats() (*Video, error)
	DownloadVideo(format *Format) error
}

func NewDownloader(source *input.VideoSource) (Downloader, error) {
	switch source.Type {
	case input.YoutubeVideo:
		return &YtdlpDownloader{
			source: source,
		}, nil
	default:
		return nil, errors.New("unsupported video source")
	}
}

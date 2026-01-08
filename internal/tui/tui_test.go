package tui

import (
	"errors"
	"testing"

	"github.com/cristhianflo/vid-dl/internal/downloader"
)

type fakeDownloader struct {
	Formats   []downloader.Format
	ShouldErr bool
}

func (f *fakeDownloader) GetFormats() (*downloader.Video, error) {
	if f.ShouldErr {
		return nil, errors.New("fake error")
	}
	return &downloader.Video{ID: "ID123", Title: "Title123", Formats: f.Formats}, nil
}

func (f *fakeDownloader) DownloadVideo(format *downloader.Format) error {
	if format.ID == "failme" {
		return errors.New("download error")
	}
	return nil
}

func TestNewModel_ErrorOnGetFormats(t *testing.T) {
	_, err := NewModel(&fakeDownloader{ShouldErr: true})
	if err == nil {
		t.Error("expected error on formats failure")
	}
}

func TestNewModel_NoFormats(t *testing.T) {
	_, err := NewModel(&fakeDownloader{Formats: nil})
	if err == nil || err.Error() != "no video formats available" {
		t.Errorf("should error on empty formats, got: %v", err)
	}
}

func TestNewModel_SingleFormat(t *testing.T) {
	f := downloader.Format{ID: "foo", Ext: downloader.Mp4, Resolution: "720p", Fps: 24, Filesize: 2000}
	model, err := NewModel(&fakeDownloader{Formats: []downloader.Format{f}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(model.formats) != 1 || model.formats[0].ID != "foo" {
		t.Errorf("wrong formats: %+v", model.formats)
	}
}

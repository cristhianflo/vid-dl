package downloader

import (
	"testing"
)

func TestMakeVideoFormat(t *testing.T) {
	res := "720p"
	fps := float32(30)
	size := int64(1000000)
	in := ytdlpFormat{
		FormatID:   "22",
		Ext:        "mp4",
		Resolution: &res,
		Fps:        &fps,
		Filesize:   &size,
		ACodec:     "aac",
		VCodec:     "h264",
	}
	f := makeVideoFormat(in)
	if f.ID != "22+bestaudio" {
		t.Errorf("want ID 22+bestaudio, got %s", f.ID)
	}
	if f.Ext != Mp4 {
		t.Errorf("want Mp4 extension, got %v", f.Ext)
	}
	if f.Resolution != "720p" {
		t.Errorf("want resolution 720p, got %v", f.Resolution)
	}
	if f.Fps != 30 {
		t.Errorf("want Fps 30, got %v", f.Fps)
	}
	if f.Filesize != 1000000 {
		t.Errorf("want Filesize 1000000, got %v", f.Filesize)
	}
}

func TestMakeAudioFormat(t *testing.T) {
	size := int64(54321)
	in := ytdlpFormat{
		Filesize: &size,
	}
	f := makeAudioFormat(in)
	if f.ID != "bestaudio" {
		t.Errorf("want ID bestaudio, got %s", f.ID)
	}
	if f.Ext != Mp3 {
		t.Errorf("want Mp3 extension, got %v", f.Ext)
	}
	if f.Resolution != "audio only" {
		t.Errorf("want 'audio only', got %v", f.Resolution)
	}
	if f.Filesize != size {
		t.Errorf("want Filesize %v, got %v", size, f.Filesize)
	}
}

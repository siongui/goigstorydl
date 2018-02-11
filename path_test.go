package igstorydl

import (
	"testing"
)

func TestStripQueryString(t *testing.T) {
	u := StripQueryString("https://example.com/myvideo.mp4?abc=d")
	if u != "https://example.com/myvideo.mp4" {
		t.Error(u)
	}
}

func TestFileExt(t *testing.T) {
	ext := FileExt("https://example.com/myvideo.mp4?abc=d")
	if ext != ".mp4" {
		t.Error(ext)
	}
}

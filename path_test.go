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

func TestAddTitleInPath(t *testing.T) {
	np := AddTitleInPath("stories/username/username-2018-02-10T23:16:49+08:00-1518275809.mp4", "title")
	if np != "stories/username/title/username-2018-02-10T23:16:49+08:00-1518275809.mp4" {
		t.Error(np)
	}
}

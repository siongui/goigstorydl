package igstorydl

import (
	"github.com/siongui/goigstorylink"
	"os"
	"testing"
	"time"
)

func ExampleMonitorAndDownload(t *testing.T) {
	MonitorAndDownload(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
}

func ExampleDownloadHighlight(t *testing.T) {
	igstory.SetUserId(os.Getenv("IG_DS_USER_ID"))
	igstory.SetSessionId(os.Getenv("IG_SESSIONID"))
	igstory.SetCsrfToken(os.Getenv("IG_CSRFTOKEN"))
	for {
		DownloadHighlight(
			os.Getenv("IG_DS_USER_ID"),
			os.Getenv("IG_SESSIONID"),
			os.Getenv("IG_CSRFTOKEN"))
		time.Sleep(30 * time.Minute)
	}
}

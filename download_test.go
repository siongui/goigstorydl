package igstorydl

import (
	"os"
	"testing"
)

func ExampleMonitorAndDownload(t *testing.T) {
	MonitorAndDownload(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
}

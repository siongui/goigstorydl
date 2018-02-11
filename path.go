package igstorydl

import (
	"net/url"
	"path"
	"strconv"
	"time"
)

// Remove query string in the URL
func StripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
}

// Get file extension from URL
func FileExt(url string) string {
	filename := path.Base(StripQueryString(url))
	return path.Ext(filename)
}

func FormatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

// Default output folder: stories/
//
// TODO: should allow to change output folder.
func BuildOutputFilePath(username, url string, timestamp int64) string {
	dirname := path.Join("stories", username)
	CreateDirIfNotExist(dirname)
	ext := FileExt(url)
	ts := FormatTimestamp(timestamp)
	filename := username + "-" + ts + "-" + strconv.FormatInt(timestamp, 10) + ext
	p := path.Join(dirname, filename)
	return p
}

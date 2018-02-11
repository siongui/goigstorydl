// Package igstorydl helps you download Instagram stories.
package igstorydl

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/siongui/goigfollow"
	"github.com/siongui/goigstorylink"
	"os"
	"os/exec"
	"time"
)

// Call shell command wget to download. The reason to use wget is that wget
// supports automatically resume download. So this package only runs on Linux
// systems.
func Wget(url, filepath string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Print info of the story being downloaded, including username, timestamp,
// URL of the story, and saved file path in local machine.
func PrintDownloadInfo(username, url, filepath string, timestamp int64) {
	cc := color.New(color.FgCyan)
	rc := color.New(color.FgRed)
	fmt.Print("Username: ")
	rc.Println(username)
	fmt.Print("Story timestamp: ")
	rc.Println(FormatTimestamp(timestamp))
	fmt.Print("Download ")
	cc.Print(url)
	fmt.Print(" to ")
	rc.Print(filepath)
	fmt.Println(" ...")
}

func DownloadIGUser(user igstory.IGUser, isHighlight bool) {
	for _, story := range user.Stories {
		// BuildOutputFilePath also create dir if not exist
		p := BuildOutputFilePath(user.Username, story.Url, story.Timestamp)
		if isHighlight {
			p = AddTitleInPath(p, user.Title)
		}
		// check if file exist
		if _, err := os.Stat(p); os.IsNotExist(err) {
			// file not exists
			PrintDownloadInfo(user.Username, story.Url, p, story.Timestamp)
			err = Wget(story.Url, p)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// Download unexpired stories of users with unread stories.
func DownloadUnread() {
	users, err := igstory.GetUnreadStories()
	if err != nil {
		// return error? or just print?
		fmt.Println(err)
		return
	}

	for _, user := range users {
		DownloadIGUser(user, false)
	}
}

// Download all unexpired stories
func DownloadAll() {
	users, err := igstory.GetAllStories()
	if err != nil {
		// return error? or just print?
		fmt.Println(err)
		return
	}

	for _, user := range users {
		DownloadIGUser(user, false)
	}
}

// Download highlight stories of following users.
func DownloadHighlight(ds_user_id, sessionid, csrftoken string) {
	rf, err := igfollow.GetFollowing(ds_user_id, ds_user_id, sessionid, csrftoken)
	if err != nil {
		// return error? or just print?
		fmt.Println(err)
		return
	}

	for _, user := range rf.Users {
		husers, err := igstory.GetUserHighlightStories(user.Pk)
		if err != nil {
			// return error? or just print?
			fmt.Println(err)
			return
		}
		for _, huser := range husers {
			DownloadIGUser(huser, true)
		}
	}
}

// Given *ds_user_id*, *sessionid*, and *csrftoken* cookies, monitor and
// download stories automatically.
func MonitorAndDownload(userid, sessionid, csrftoken string) {
	igstory.SetUserId(userid)
	igstory.SetSessionId(sessionid)
	igstory.SetCsrfToken(csrftoken)

	cc := color.New(color.FgCyan)
	rc := color.New(color.FgRed)
	sleepInterval := 30 // seconds
	count := 0
	for {
		if count == 0 {
			DownloadAll()
			cc.Println("Download all stories finished")
		} else {
			DownloadUnread()
			cc.Println("Download unread stories finished")
		}
		count++
		count %= 5

		rc.Print(time.Now().Format(time.RFC3339))
		fmt.Print(": sleep ")
		cc.Print(sleepInterval)
		fmt.Println(" second")
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}

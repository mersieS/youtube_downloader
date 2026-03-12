package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func checkDependencies() error {
	if _, err := exec.LookPath("yt-dlp"); err != nil {
		return fmt.Errorf("yt-dlp bulunamadı. Kurmak için: brew install yt-dlp")
	}
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return fmt.Errorf("ffmpeg bulunamadı. Kurmak için: brew install ffmpeg")
	}
	return nil
}

func isValidYouTubeURL(url string) bool {
	prefixes := []string{
		"https://www.youtube.com/watch",
		"https://youtube.com/watch",
		"https://youtu.be/",
		"https://m.youtube.com/watch",
		"http://www.youtube.com/watch",
		"http://youtube.com/watch",
		"http://youtu.be/",
		"https://www.youtube.com/shorts/",
		"https://youtube.com/shorts/",
	}
	for _, p := range prefixes {
		if strings.HasPrefix(url, p) {
			return true
		}
	}
	return false
}

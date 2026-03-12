package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func downloadMP4(url, quality, outputDir string) error {
	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")

	var formatArg string
	if quality == "best" {
		formatArg = "bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best"
	} else {
		formatArg = fmt.Sprintf("bestvideo[height<=%s][ext=mp4]+bestaudio[ext=m4a]/bestvideo[height<=%s]+bestaudio/best[height<=%s]/best", quality, quality, quality)
	}

	fmt.Printf("\n  MP4 indiriliyor (%s)...\n", quality)

	cmd := exec.Command("yt-dlp",
		"--format", formatArg,
		"--merge-output-format", "mp4",
		"--embed-thumbnail",
		"--add-metadata",
		"--output", outputTemplate,
		"--no-playlist",
		"--progress",
		url,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp hatası: %w", err)
	}

	fmt.Printf("  Tamamlandı! Dosya: %s\n", outputDir)
	return nil
}

func downloadMP3(url, outputDir string) error {
	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")

	fmt.Println("\n  MP3 indiriliyor...")

	cmd := exec.Command("yt-dlp",
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "0",
		"--embed-thumbnail",
		"--add-metadata",
		"--output", outputTemplate,
		"--no-playlist",
		"--progress",
		url,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp hatası: %w", err)
	}

	fmt.Printf("  Tamamlandı! Dosya: %s\n", outputDir)
	return nil
}

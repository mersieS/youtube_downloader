package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader  *bufio.Reader
	version = "dev"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("ytmp3 v%s\n", version)
		return
	}

	if err := checkDependencies(); err != nil {
		fmt.Fprintf(os.Stderr, "Hata: %v\n", err)
		os.Exit(1)
	}

	reader = bufio.NewReader(os.Stdin)

	for {
		mode := selectMenuWithBanner("Platform seçin:", []string{
			"YouTube",
			"Diğer (Tüm siteler)",
		}, [][]string{bannerYouTube, bannerOther})

		if mode == -1 {
			fmt.Println("Görüşürüz!")
			break
		}

		isYouTube := mode == 0

		for {
			var prompt string
			if isYouTube {
				prompt = "YouTube URL girin (geri 'b' · çıkış 'q'): "
			} else {
				prompt = "Video URL girin (geri 'b' · çıkış 'q'): "
			}
			fmt.Print(prompt)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "Okuma hatası: %v\n", err)
				continue
			}

			input = strings.TrimSpace(input)
			if input == "q" || input == "quit" || input == "exit" {
				fmt.Println("Görüşürüz!")
				return
			}
			if input == "b" || input == "back" {
				fmt.Println()
				break
			}
			if input == "" {
				continue
			}

			if isYouTube && !isValidYouTubeURL(input) {
				fmt.Println("Geçersiz YouTube URL'si. Tekrar deneyin.")
				continue
			}
			if !isYouTube && !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
				fmt.Println("Geçersiz URL. http:// veya https:// ile başlamalı.")
				continue
			}

			fmt.Println()
			format := selectMenu("Format seçin:", []string{
				"MP4 (Video)",
				"MP3 (Ses)",
			})
			if format == -1 {
				fmt.Println()
				continue
			}

			var defaultDir string
			if isYouTube {
				if format == 0 {
					defaultDir = "Video"
				} else {
					defaultDir = "Muzik"
				}
			} else {
				if format == 0 {
					defaultDir = "Other/Video"
				} else {
					defaultDir = "Other/Muzik"
				}
			}
			outputDir := askOutputDir(defaultDir)
			if outputDir == "" {
				continue
			}

			switch format {
			case 0:
				fmt.Println()
				qi := selectMenu("Video kalitesi seçin:", []string{
					"2160p (4K)",
					"1440p (2K)",
					"1080p (Full HD)",
					"720p  (HD)",
					"480p",
					"360p",
					"En iyi kalite (otomatik)",
				})
				if qi == -1 {
					fmt.Println()
					continue
				}
				qualities := []string{"2160", "1440", "1080", "720", "480", "360", "best"}
				if err := downloadMP4(input, qualities[qi], outputDir); err != nil {
					fmt.Fprintf(os.Stderr, "İndirme hatası: %v\n", err)
				}
			case 1:
				if err := downloadMP3(input, outputDir); err != nil {
					fmt.Fprintf(os.Stderr, "İndirme hatası: %v\n", err)
				}
			}

			fmt.Println()
		}
	}
}

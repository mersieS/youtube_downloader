package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

// selectMenuWithBanner redraws a banner above the menu when selection changes.
func selectMenuWithBanner(title string, options []string, banners [][]string) int {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return -1
	}
	defer term.Restore(fd, oldState)

	selected := 0
	buf := make([]byte, 3)
	maxBanner := maxBannerLines()
	menuLines := len(options) + 3
	totalLines := maxBanner + 2 + menuLines
	firstRender := true

	render := func() {
		if !firstRender {
			fmt.Printf("\033[%dA\r", totalLines)
		}
		firstRender = false
		fmt.Print("\033[J")

		banner := banners[selected]
		fmt.Print("\r\n")
		for _, l := range banner {
			fmt.Printf("%s\r\n", l)
		}
		for i := len(banner); i < maxBanner; i++ {
			fmt.Print("\r\n")
		}
		fmt.Print("\r\n")

		fmt.Printf("  \033[1;37m%s\033[0m\r\n", title)
		for i, opt := range options {
			if i == selected {
				fmt.Printf("  \033[1;36m❯ %s\033[0m\r\n", opt)
			} else {
				fmt.Printf("    \033[0;37m%s\033[0m\r\n", opt)
			}
		}
		fmt.Printf("\r\n  \033[2m↑↓ seç · Enter onayla · q çık\033[0m\r\n")
	}

	render()

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return -1
		}

		moved := false

		if n == 1 {
			switch buf[0] {
			case 13:
				fmt.Printf("\033[%dA\r\033[J", totalLines)
				banner := banners[selected]
				fmt.Print("\r\n")
				for _, l := range banner {
					fmt.Printf("%s\r\n", l)
				}
				fmt.Print("\r\n")
				fmt.Printf("  \033[1;37m%s\033[0m \033[1;36m%s\033[0m\r\n", title, options[selected])
				return selected
			case 'q':
				fmt.Printf("\033[%dA\r\033[J", totalLines)
				return -1
			case 'k':
				if selected > 0 {
					selected--
					moved = true
				}
			case 'j':
				if selected < len(options)-1 {
					selected++
					moved = true
				}
			}
		}

		if n == 3 && buf[0] == 27 && buf[1] == 91 {
			switch buf[2] {
			case 65:
				if selected > 0 {
					selected--
					moved = true
				}
			case 66:
				if selected < len(options)-1 {
					selected++
					moved = true
				}
			}
		}

		if moved {
			render()
		}
	}
}

func selectMenu(title string, options []string) int {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return -1
	}
	defer term.Restore(fd, oldState)

	selected := 0
	buf := make([]byte, 3)
	totalLines := len(options) + 3
	firstRender := true

	render := func() {
		if !firstRender {
			fmt.Printf("\033[%dA\r", totalLines)
		}
		firstRender = false
		fmt.Print("\033[J")
		fmt.Printf("  \033[1;37m%s\033[0m\r\n", title)
		for i, opt := range options {
			if i == selected {
				fmt.Printf("  \033[1;36m❯ %s\033[0m\r\n", opt)
			} else {
				fmt.Printf("    \033[0;37m%s\033[0m\r\n", opt)
			}
		}
		fmt.Printf("\r\n  \033[2m↑↓ seç · Enter onayla · q iptal\033[0m\r\n")
	}

	render()

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return -1
		}

		moved := false

		if n == 1 {
			switch buf[0] {
			case 13:
				fmt.Printf("\033[%dA\r\033[J", totalLines)
				fmt.Printf("  \033[1;37m%s\033[0m \033[1;36m%s\033[0m\r\n", title, options[selected])
				return selected
			case 'q':
				fmt.Printf("\033[%dA\r\033[J", totalLines)
				return -1
			case 'k':
				if selected > 0 {
					selected--
					moved = true
				}
			case 'j':
				if selected < len(options)-1 {
					selected++
					moved = true
				}
			}
		}

		if n == 3 && buf[0] == 27 && buf[1] == 91 {
			switch buf[2] {
			case 65:
				if selected > 0 {
					selected--
					moved = true
				}
			case 66:
				if selected < len(options)-1 {
					selected++
					moved = true
				}
			}
		}

		if moved {
			render()
		}
	}
}

func askOutputDir(defaultSubdir string) string {
	home, _ := os.UserHomeDir()
	defaultPath := filepath.Join(home, "Music", "YouTube", defaultSubdir)

	fmt.Println()
	fmt.Printf("  Kayıt klasörü [Enter = %s]:\n", defaultPath)
	fmt.Print("  Klasör: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.TrimSpace(input)

	dir := defaultPath
	if input != "" {
		if strings.HasPrefix(input, "~") {
			input = filepath.Join(home, input[1:])
		}
		dir = input
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "  Klasör oluşturulamadı: %v\n", err)
		return ""
	}

	return dir
}

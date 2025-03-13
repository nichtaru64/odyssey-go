package ui

import (
	"fmt"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

var AsciiArt = []string{
	" ▄██████▄  ████████▄  ▄██   ▄      ▄████████    ▄████████    ▄████████ ▄██   ▄ ",
	"███    ███ ███   ▀███ ███   ██▄   ███    ███   ███    ███   ███    ███ ███   ██▄",
	"███    ███ ███    ███ ███▄▄▄███   ███    █▀    ███    █▀    ███    █▀  ███▄▄▄███",
	"███    ███ ███    ███ ▀▀▀▀▀▀███   ███          ███         ▄███▄▄▄     ▀▀▀▀▀▀███",
	"███    ███ ███    ███ ▄██   ███ ▀███████████ ▀███████████ ▀▀███▀▀▀     ▄██   ███",
	"███    ███ ███    ███ ███   ███          ███          ███   ███    █▄  ███   ███",
	"███    ███ ███   ▄███ ███   ███    ▄█    ███    ▄█    ███   ███    ███ ███   ███",
	" ▀██████▀  ████████▀   ▀█████▀   ▄████████▀   ▄████████▀    ██████████  ▀█████▀ ",
}

var FirstTime = true

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintAsciiArt() {
	fmt.Printf("%s", ColorBlue)
	for _, line := range AsciiArt {
		fmt.Println(line)
		if FirstTime {
			time.Sleep(400 * time.Millisecond)
		}
	}
	fmt.Printf("%s", ColorReset)
	FirstTime = false
}

func ShowMainMenu(selectedIndex int) {
	ClearScreen()
	PrintAsciiArt()

	fmt.Printf("\n%sHauptmenü%s\n", ColorPurple, ColorReset)
	items := []string{"Start", "Optionen", "Beenden"}

	for index, item := range items {
		if index == selectedIndex {
			fmt.Printf("%s> %s%s\n", ColorCyan, item, ColorReset)
		} else {
			fmt.Printf("%s  %s%s\n", ColorWhite, item, ColorReset)
		}
	}
}

func ShowOptionsMenu(currentSelection int) {
	ClearScreen()
	fmt.Printf("%sOptionen%s\n\n", ColorPurple, ColorReset)
	items := []string{"Zurück zum Hauptmenü", "Beenden"}
	for index, item := range items {
		if index == currentSelection {
			fmt.Printf("%s> %s%s\n", ColorCyan, item, ColorReset)
		} else {
			fmt.Printf("%s  %s%s\n", ColorWhite, item, ColorReset)
		}
	}
}

func SchreibAnimation(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(10 * time.Millisecond) // Fügt delay zwischen den zeichen hinzu
	}
	fmt.Println() // Zeilenumbruch
}

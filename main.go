package main

import (
	"math/rand"
	"odyssey/game"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	// Zufallszahlengenerator initialisieren
	rand.Seed(time.Now().UnixNano())

	// Tastatur initialisieren
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// Starte das Hauptmenü
	game.MainMenu()
}

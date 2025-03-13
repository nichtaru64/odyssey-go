package game

import (
	"fmt"
	"math/rand"
	"odyssey/models"
	"odyssey/ui"
	"time"

	"github.com/eiannone/keyboard"
)

// Allgemeine Spielwelt erstellt
var Welt models.Spielwelt

// Spielwelt initialisieren
func InitWelt() {
	// Zyklopen-Insel
	zyklopenWesen := []models.Wesen{
		{
			Name:         "Polyphem",
			Freundlich:   false,
			Staerke:      8,
			Beschreibung: "Ein einäugiger Riese, der Schafe züchtet und Menschen frisst.",
		},
	}

	zyklopenInsel := models.Insel{
		Name:         "Insel der Zyklopen",
		Groesse:      7,
		Bevoelkerung: 3,
		Kuesten:      8,
		Wesen:        zyklopenWesen,
		Besucht:      false,
	}

	// Kirke-Insel
	kirkeWesen := []models.Wesen{
		{
			Name:         "Kirke",
			Freundlich:   false,
			Staerke:      6,
			Beschreibung: "Eine Zauberin, die Menschen in Schweine verwandelt.",
		},
	}

	kirkeInsel := models.Insel{
		Name:         "Aia",
		Groesse:      5,
		Bevoelkerung: 2,
		Kuesten:      4,
		Wesen:        kirkeWesen,
		Besucht:      false,
	}

	// Sirenen-Insel
	sirenenWesen := []models.Wesen{
		{
			Name:         "Sirenen",
			Freundlich:   false,
			Staerke:      5,
			Beschreibung: "Wesen, die mit ihrem verführerischen Gesang Seeleute anlocken und töten.",
		},
	}

	sirenenInsel := models.Insel{
		Name:         "Insel der Sirenen",
		Groesse:      3,
		Bevoelkerung: 2,
		Kuesten:      6,
		Wesen:        sirenenWesen,
		Besucht:      false,
	}

	Mensch := []models.Wesen{
		{
			Name:         "Menschen",
			Freundlich:   true,
			Staerke:      5,
			Beschreibung: "Auf der Insel befinden sich verschiedene Menschen, die willig sind, sich der Reise des Odysseus anzuschließen.",
		},
	}

	phaeakenInsel := models.Insel{
		Name:         "Insel der Phäaken",
		Groesse:      3,
		Bevoelkerung: 20,
		Kuesten:      10,
		Wesen:        Mensch,
		Besucht:      false,
	}

	// Ithaka
	ithakaWesen := []models.Wesen{
		{
			Name:         "Penelope",
			Freundlich:   true,
			Staerke:      3,
			Beschreibung: "Die treue Frau des Odysseus.",
		},
		{
			Name:         "Freier",
			Freundlich:   false,
			Staerke:      7,
			Beschreibung: "Männer, die um Penelopes Hand anhalten und Odysseus Besitz verschwenden.",
		},
	}

	ithakaInsel := models.Insel{
		Name:         "Ithaka",
		Groesse:      8,
		Bevoelkerung: 7,
		Kuesten:      2,
		Wesen:        ithakaWesen,
		Besucht:      false,
	}

	// Alle Inseln zur Spielwelt hinzufügen
	Welt.Inseln = []models.Insel{phaeakenInsel, zyklopenInsel, kirkeInsel, sirenenInsel, ithakaInsel}

	// Spieler initialisieren
	Welt.Spieler = models.Spieler{
		Name:           "Odysseus",
		Crewmitglieder: 100,
		AktuelleInsel:  nil,
		BesuchteInseln: []*models.Insel{},
	}
}

// Eine zufällige Insel finden, die noch nicht besucht wurde
func FindeNaechsteInsel() *models.Insel {
	unbesuchteInseln := []models.Insel{}

	for i, insel := range Welt.Inseln {
		if !insel.Besucht {
			unbesuchteInseln = append(unbesuchteInseln, insel)
			return &Welt.Inseln[i] // Verweis auf die tatsächliche Insel in der Weltliste
		}
	}

	if len(unbesuchteInseln) == 0 {
		return nil // Keine unbesuchten Inseln mehr
	}

	// Zufällige unbesuchte Insel auswählen
	randomIndex := rand.Intn(len(unbesuchteInseln))

	// Finde die entsprechende Insel in der Weltliste
	for i, insel := range Welt.Inseln {
		if insel.Name == unbesuchteInseln[randomIndex].Name {
			return &Welt.Inseln[i]
		}
	}

	return nil
}

// Insel besuchen
func BesucheInsel(insel *models.Insel) {
	ui.ClearScreen()
	fmt.Printf("%sInsel entdeckt: %s%s\n", ui.ColorYellow, insel.Name, ui.ColorReset)
	time.Sleep(1 * time.Second)

	// Crew-Verluste durch Küstenschwierigkeit
	verluste := insel.Kuesten * 2

	if verluste > 0 {
		var neueVerluste int
		var endVerluste int
		if verluste > Welt.Spieler.Crewmitglieder {
			neueVerluste = verluste - Welt.Spieler.Crewmitglieder
			endVerluste = verluste - neueVerluste
		}
		if verluste < Welt.Spieler.Crewmitglieder {
			endVerluste = verluste
		}
		Welt.Spieler.Crewmitglieder -= endVerluste
		fmt.Printf("%sDas Ansteuern der Küste hat dich %s%d Crewmitglieder%s gekostet.\n",
			ui.ColorRed, ui.ColorWhite, endVerluste, ui.ColorRed)
		fmt.Printf("Du hast noch %s%d Crewmitglieder%s.\n\n",
			ui.ColorWhite, Welt.Spieler.Crewmitglieder, ui.ColorReset)
		time.Sleep(2 * time.Second)
	}

	// Insel beschreiben
	BeschreibeInsel(insel)

	// Mit Wesen interagieren
	for i := range insel.Wesen {
		InteragiereMitWesen(&insel.Wesen[i])
	}

	// Insel als besucht markieren
	insel.Besucht = true
	Welt.Spieler.BesuchteInseln = append(Welt.Spieler.BesuchteInseln, insel)
	Welt.Spieler.AktuelleInsel = insel

	fmt.Printf("\n%sDrücke eine Taste, um fortzufahren...%s", ui.ColorCyan, ui.ColorReset)
	keyboard.GetKey()
}

func BeschreibeInsel(insel *models.Insel) {
	fmt.Printf("%sDu erkundest %s...%s\n", ui.ColorGreen, insel.Name, ui.ColorReset)

	switch insel.Name {
	case "Insel der Zyklopen":
		ui.SchreibAnimation("Die Insel erscheint zunächst verlassen, doch in Höhlen leben riesige einäugige Kreaturen.")
		ui.SchreibAnimation("Deine Männer entdecken eine Höhle mit Käse und Lämmern und beschließen einzutreten.")
	case "Aia":
		ui.SchreibAnimation("Ein dichter Wald bedeckt die Insel. In der Mitte steht ein prächtiges Haus.")
		ui.SchreibAnimation("Aus dem Haus dringt verlockender Gesang einer Frauenstimme.")
	case "Insel der Sirenen":
		ui.SchreibAnimation("Die Insel ist von tückischen Klippen umgeben, an denen viele Schiffswracks liegen.")
		ui.SchreibAnimation("Aus der Ferne hörst du einen unwiderstehlich schönen Gesang.")
	case "Ithaka":
		ui.SchreibAnimation("Endlich, deine Heimat! Doch etwas stimmt nicht. Der Palast wird von Fremden bewohnt.")
		ui.SchreibAnimation("Du musst vorsichtig sein, wenn du dein Königreich zurückerobern willst.")
	default:
		ui.SchreibAnimation("Du erkundest die unbekannte Insel und triffst auf ihre Bewohner.")
	}

	fmt.Println()
}

func InteragiereMitWesen(wesen *models.Wesen) {
	fmt.Printf("%sDu begegnest: %s%s%s\n", ui.ColorYellow, ui.ColorWhite, wesen.Name, ui.ColorReset)
	ui.SchreibAnimation(wesen.Beschreibung)

	if wesen.Freundlich {
		fmt.Printf("%s%s scheint freundlich zu sein.%s\n", ui.ColorGreen, wesen.Name, ui.ColorReset)
		if wesen.Name == "Menschen" {
			var randZahl int = rand.Intn(20)
			Welt.Spieler.Crewmitglieder = Welt.Spieler.Crewmitglieder + randZahl
			fmt.Printf("Du hast jetzt %d Crewmitglieder! (+ %d)\n", Welt.Spieler.Crewmitglieder, randZahl)
		}
		// Positive Interaktion
		ui.SchreibAnimation("Die Begegnung verläuft friedlich und du erhältst wertvolle Informationen.")
	} else {
		fmt.Printf("%s%s ist feindlich gesinnt!%s\n", ui.ColorRed, wesen.Name, ui.ColorReset)
		// Kampf oder Flucht
		kampfErgebnis := wesen.Staerke - (rand.Intn(5) + 1) // Zufallsfaktor
		verluste := kampfErgebnis * 5

		if verluste > 0 {
			var neueVerluste int
			var endVerluste int
			if verluste > Welt.Spieler.Crewmitglieder {
				neueVerluste = verluste - Welt.Spieler.Crewmitglieder
				endVerluste = verluste - neueVerluste
			}
			if verluste < Welt.Spieler.Crewmitglieder {
				endVerluste = verluste
			}
			Welt.Spieler.Crewmitglieder -= endVerluste
			fmt.Printf("%sDu verlierst %d Crewmitglieder im Kampf!%s\n",
				ui.ColorRed, endVerluste, ui.ColorReset)
			fmt.Printf("Du hast noch %s%d Crewmitglieder%s.\n",
				ui.ColorWhite, Welt.Spieler.Crewmitglieder, ui.ColorReset)
		} else {
			fmt.Printf("%sDu besiegst %s ohne Verluste!%s\n",
				ui.ColorGreen, wesen.Name, ui.ColorReset)
		}
	}

	fmt.Println()
	time.Sleep(2 * time.Second)
}

// MainMenu ist das Hauptmenü des Spiels
func MainMenu() {
	currentSelection := 0
	ui.ShowMainMenu(currentSelection)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		switch {
		case key == keyboard.KeyArrowUp:
			if currentSelection > 0 {
				currentSelection--
			}
		case key == keyboard.KeyArrowDown:
			if currentSelection < 2 {
				currentSelection++
			}
		case key == keyboard.KeyEnter:
			switch currentSelection {
			case 0: // Start
				ui.ClearScreen()
				fmt.Printf("%sDas Spiel startet...%s\n", ui.ColorYellow, ui.ColorReset)
				time.Sleep(2 * time.Second)
				ui.ClearScreen()
				Game()
			case 1: // Optionen
				OptionsMenu()
				ui.ShowMainMenu(currentSelection)
			case 2: // Beenden
				ui.ClearScreen()
				fmt.Printf("%sWird beendet!%s\n", ui.ColorGreen, ui.ColorReset)
				return
			}
		}
		ui.ShowMainMenu(currentSelection)
	}
}

func OptionsMenu() {
	ui.ClearScreen()
	fmt.Printf("%sOptionen%s\n\n", ui.ColorPurple, ui.ColorReset)
	currentSelection := 0

	ui.ShowOptionsMenu(currentSelection)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch {
		case key == keyboard.KeyArrowUp:
			if currentSelection > 0 {
				currentSelection--
				ui.ShowOptionsMenu(currentSelection)
			}
		case key == keyboard.KeyArrowDown:
			if currentSelection < 1 {
				currentSelection++
				ui.ShowOptionsMenu(currentSelection)
			}
		case key == keyboard.KeyEnter:
			switch currentSelection {
			case 0: // zurück
				return
			case 1: // beenden
				ui.ClearScreen()
				fmt.Printf("%sWird beendet!%s\n", ui.ColorGreen, ui.ColorReset)
				keyboard.Close()
				panic("exit")
			}
		}
	}
}

// Game startet das Spiel
func Game() {
	// Initialisiere die Spielwelt
	InitWelt()

	for {
		naechsteInsel := FindeNaechsteInsel()
		if naechsteInsel == nil {
			// Spielende, alle Inseln besucht
			ui.ClearScreen()
			fmt.Printf("%sGlückwunsch! Du hast alle Inseln besucht und deine Heimreise abgeschlossen!%s\n",
				ui.ColorGreen, ui.ColorReset)
			fmt.Printf("\n%sDrücke eine Taste, um zum Hauptmenü zurückzukehren...%s",
				ui.ColorCyan, ui.ColorReset)
			keyboard.GetKey()
			return
		}

		// Besuche die nächste Insel
		BesucheInsel(naechsteInsel)

		// Überprüfe, ob der Spieler tot ist
		if Welt.Spieler.Crewmitglieder <= 0 {
			ui.ClearScreen()
			fmt.Printf("%sDu hast alle deine Crewmitglieder verloren! Spiel vorbei!%s\n",
				ui.ColorRed, ui.ColorReset)
			fmt.Printf("\n%sDrücke eine Taste, um zum Hauptmenü zurückzukehren...%s",
				ui.ColorCyan, ui.ColorReset)
			keyboard.GetKey()
			return
		}
	}
}

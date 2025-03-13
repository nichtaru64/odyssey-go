package models

// Insel repräsentiert eine Insel im Spiel
type Insel struct {
	Name         string
	Groesse      int // 1-10: Wie einfach kann man sie entdecken
	Bevoelkerung int // 1-10: Ein Modifikator zur Zutraulichkeit der Bewohner
	Kuesten      int // 1-10: Wie viele Crewmitglieder das Ansteuern kostet
	Wesen        []Wesen
	Besucht      bool
}

// Wesen im Spiel
type Wesen struct {
	Name         string
	Freundlich   bool
	Staerke      int // 1-10
	Beschreibung string
}

// Spieler/Hauptcharakter (Odysseus)
type Spieler struct {
	Name           string
	Crewmitglieder int
	AktuelleInsel  *Insel
	BesuchteInseln []*Insel
}

// Spielwelt enthält alle Inseln und Odysseus
type Spielwelt struct {
	Inseln  []Insel
	Spieler Spieler
}

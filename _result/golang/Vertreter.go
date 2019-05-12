package golang

type Vertreter struct {
	VertreterNr string      `json:VertreterNr`
	Name        string      `json:Name`
	Mitarbeiter Mitarbeiter `json:Mitarbeiter`
	Geloescht   bool        `json:Geloescht`
}

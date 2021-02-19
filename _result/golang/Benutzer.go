package golang

type Benutzer struct {
	BenutzerNr        string    `json:BenutzerNr`
	Passwort          string    `json:Passwort`
	ADUser            string    `json:ADUser`
	IstAdmin          bool      `json:IstAdmin`
	Bemerkungen       string    `json:Bemerkungen`
	EmailAbsender     string    `json:EmailAbsender`
	EmailAbsenderName string    `json:EmailAbsenderName`
	Lieferart         Lieferart `json:Lieferart`
	Lagerort          Lagerort  `json:Lagerort`
	Sprache           int       `json:Sprache`
	UnserZ            string    `json:UnserZ`
	IstBenutzer       bool      `json:IstBenutzer`
}

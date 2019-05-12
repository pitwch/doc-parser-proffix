package golang

type Standort struct {
	StandortNr    int         `json:StandortNr`
	Bezeichnung   string      `json:Bezeichnung`
	Hauptstandort bool        `json:Hauptstandort`
	Auftrag       Auftrag     `json:Auftrag`
	Intern        bool        `json:Intern`
	Kunde         Kunde       `json:Kunde`
	Prioritaet    Prioritaet  `json:Prioritaet`
	Vertragstyp   Vertragstyp `json:Vertragstyp`
	Geloescht     bool        `json:Geloescht`
	Parameter     string      `json:Parameter`
}

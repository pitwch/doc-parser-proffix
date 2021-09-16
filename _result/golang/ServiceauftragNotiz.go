package golang

type ServiceauftragNotiz struct {
	NotizNr        int            `json:"NotizNr"`
	Serviceauftrag Serviceauftrag `json:"Serviceauftrag"`
	Mitarbeiter    Mitarbeiter    `json:"Mitarbeiter"`
	Notizart       Notizart       `json:"Notizart"`
	Position       int            `json:"Position"`
	Intern         int            `json:"Intern"`
	Datum          string         `json:"Datum"`
	Termin         string         `json:"Termin"`
	Notiz          string         `json:"Notiz"`
	NotizRTF       string         `json:"NotizRTF"`
	AlarmDatumVon  string         `json:"AlarmDatumVon"`
	AlarmDatumBis  string         `json:"AlarmDatumBis"`
}

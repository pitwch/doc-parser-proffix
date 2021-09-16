package golang

type AuftragNotiz struct {
	AuftragNotizNr  int    `json:"AuftragNotizNr"`
	AlarmDatumBis   string `json:"AlarmDatumBis"`
	AlarmDatumVon   string `json:"AlarmDatumVon"`
	AuftragNotizArt string `json:"AuftragNotizArt"`
	AuftragNr       string `json:"AuftragNr"`
	Datum           string `json:"Datum"`
	ErstelltAm      string `json:"ErstelltAm"`
	ErstelltVon     string `json:"ErstelltVon"`
	Exportiert      string `json:"Exportiert"`
	Geaendert       string `json:"Geaendert"`
	GeaendertAm     string `json:"GeaendertAm"`
	GeaendertVon    string `json:"GeaendertVon"`
	ImportNr        int    `json:"ImportNr"`
	KontaktNr       int    `json:"KontaktNr"`
	MitarbeiterNr   string `json:"MitarbeiterNr"`
	Notiz           string `json:"Notiz"`
	NotizRTF        string `json:"NotizRTF"`
	OutlookID       string `json:"OutlookID"`
	Termin          string `json:"Termin"`
}

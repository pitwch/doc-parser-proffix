package golang

type Notizart struct {
	NotizartNr       string `json:"NotizartNr"`
	Bezeichnung      string `json:"Bezeichnung"`
	AuftragAlarm     bool   `json:"AuftragAlarm"`
	AlarmFarbe       int    `json:"AlarmFarbe"`
	AlarmReihenfolge int    `json:"AlarmReihenfolge"`
	AnzeigenAls      int    `json:"AnzeigenAls"`
	Aufgabe          bool   `json:"Aufgabe"`
	EinkaufAlarm     bool   `json:"EinkaufAlarm"`
	Erinnerung       int    `json:"Erinnerung"`
	Haupt            bool   `json:"Haupt"`
	Kategorien       string `json:"Kategorien"`
	PrioritaetCRM    int    `json:"PrioritaetCRM"`
	Privat           bool   `json:"Privat"`
	StatusCRM        int    `json:"StatusCRM"`
	Synchronisieren  bool   `json:"Synchronisieren"`
	Termin           bool   `json:"Termin"`
}

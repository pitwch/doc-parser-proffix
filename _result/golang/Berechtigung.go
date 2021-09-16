package golang

type Berechtigung struct {
	BerechtigungNr int            `json:"BerechtigungNr"`
	Benutzergruppe Benutzergruppe `json:"Benutzergruppe"`
	Benutzer       Benutzer       `json:"Benutzer"`
	Modul          string         `json:"Modul"`
	Form           string         `json:"Form"`
	Menu           string         `json:"Menu"`
	Recht          string         `json:"Recht"`
}

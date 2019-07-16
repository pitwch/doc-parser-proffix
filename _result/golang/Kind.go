package golang

type Kind struct {
	KindNr                    int         `json:KindNr`
	Ausbildungszulage1Von     string      `json:Ausbildungszulage1Von`
	Ausbildungszulage1Bis     string      `json:Ausbildungszulage1Bis`
	Ausbildungszulage2Von     string      `json:Ausbildungszulage2Von`
	Ausbildungszulage2Bis     string      `json:Ausbildungszulage2Bis`
	Bemerkungen               string      `json:Bemerkungen`
	Familienstatus            int         `json:Familienstatus`
	Geburtsdatum              string      `json:Geburtsdatum`
	Geschlecht                string      `json:Geschlecht`
	Mitarbeiter               Mitarbeiter `json:Mitarbeiter`
	Name                      string      `json:Name`
	Sozialversicherungsnummer string      `json:Sozialversicherungsnummer`
	Vorname                   string      `json:Vorname`
	Kinderzulage1Von          string      `json:Kinderzulage1Von`
	Kinderzulage1Bis          string      `json:Kinderzulage1Bis`
	Kinderzulage2Von          string      `json:Kinderzulage2Von`
	Kinderzulage2Bis          string      `json:Kinderzulage2Bis`
}

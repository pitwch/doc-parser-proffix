package golang

type Serviceauftrag struct {
	ServiceauftragNr      float64        `json:ServiceauftragNr`
	Kunde                 Kunde          `json:Kunde`
	Prioritaet            Prioritaet     `json:Prioritaet`
	Serviceart            Serviceart     `json:Serviceart`
	Kontaktperson         string         `json:Kontaktperson`
	KontaktpersonTel      string         `json:KontaktpersonTel`
	KontaktExtern         Adresse        `json:KontaktExtern`
	Referenztext          string         `json:Referenztext`
	Disponent             string         `json:Disponent`
	Status                Status         `json:Status`
	Datum                 string         `json:Datum`
	Kontakt               Kontakt        `json:Kontakt`
	Positionen            string         `json:Positionen`
	PositionNr            float64        `json:PositionNr`
	OriginalRapportNr     float64        `json:OriginalRapportNr`
	Installation          Position       `json:Installation`
	Prioritaet            Prioritaet     `json:Prioritaet`
	Betreueradresse       Adresse        `json:Betreueradresse`
	Externadresse         Adresse        `json:Externadresse`
	Bezeichnung1          string         `json:Bezeichnung1`
	Vertragstyp           Vertragstyp    `json:Vertragstyp`
	KonfigurationsID      string         `json:KonfigurationsID`
	Seriennummer          string         `json:Seriennummer`
	Produkt               string         `json:Produkt`
	Auftrag               Auftrag        `json:Auftrag`
	Standort              Standort       `json:Standort`
	Garantiezusatz        Garantiezusatz `json:Garantiezusatz`
	Installationstermin   string         `json:Installationstermin`
	Garantieende          string         `json:Garantieende`
	HerstellerVertrag     string         `json:HerstellerVertrag`
	AuftragProblem        string         `json:AuftragProblem`
	AuftragProblemRTF     string         `json:AuftragProblemRTF`
	Mitarbeiter           Mitarbeiter    `json:Mitarbeiter`
	Status                Status         `json:Status`
	DatumZeit             string         `json:DatumZeit`
	StartdatumZeit        string         `json:StartdatumZeit`
	EnddatumZeit          string         `json:EnddatumZeit`
	Termin                string         `json:Termin`
	FixTermin             bool           `json:FixTermin`
	Planzeit              float64        `json:Planzeit`
	Kundenbestellreferenz string         `json:Kundenbestellreferenz`
	Parameter             string         `json:Parameter`
}

package golang

type Auftrag struct {
	AuftragNr    string       `json:AuftragNr`
	Bezeichnung  string       `json:Bezeichnung`
	Kunde        Adresse      `json:Kunde`
	StartDatum   string       `json:StartDatum`
	EndDatum     string       `json:EndDatum`
	Positionsart Positionsart `json:Positionsart`
	Artikel      Artikel      `json:Artikel`
	Kostenstelle Kostenstelle `json:Kostenstelle`
	Kostenart    Kostenart    `json:Kostenart`
	Ertragskonto Konto        `json:Ertragskonto`
	Laufend      bool         `json:Laufend`
}

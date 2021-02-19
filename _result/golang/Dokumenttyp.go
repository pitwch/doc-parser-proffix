package golang

type Dokumenttyp struct {
	DokumenttypNr           string      `json:DokumenttypNr`
	Bezeichnung             string      `json:Bezeichnung`
	Inaktiv                 bool        `json:Inaktiv`
	ArtikeldokumenteDrucken bool        `json:ArtikeldokumenteDrucken`
	Aktivitaet              Aufgabe     `json:Aktivitaet`
	AuftragEroeffnen        bool        `json:AuftragEroeffnen`
	AutomatischInstallieren bool        `json:AutomatischInstallieren`
	AutomatischReservieren  bool        `json:AutomatischReservieren`
	AutomatischZahlen       bool        `json:AutomatischZahlen`
	Belegart                Belegart    `json:Belegart`
	Buchungsart             Buchungsart `json:Buchungsart`
	BuchungsartZahlung      Buchungsart `json:BuchungsartZahlung`
	IstDebitorenbuchung     bool        `json:IstDebitorenbuchung`
	Hauptdokumenttyp        bool        `json:Hauptdokumenttyp`
	KontoZahlung            Konto       `json:KontoZahlung`
	Lagerabtrag             bool        `json:Lagerabtrag`
	Liste1                  string      `json:Liste1`
	Liste2                  string      `json:Liste2`
	Liste3                  string      `json:Liste3`
	Liste4                  string      `json:Liste4`
	Liste5                  string      `json:Liste5`
	Liste6                  string      `json:Liste6`
	Liste7                  string      `json:Liste7`
	Liste8                  string      `json:Liste8`
	Liste9                  string      `json:Liste9`
	Liste10                 string      `json:Liste10`
	Negativ                 bool        `json:Negativ`
	RapportSpeichern        bool        `json:RapportSpeichern`
	Reihenfolge             int         `json:Reihenfolge`
	AktivitaetAnzeigen      bool        `json:AktivitaetAnzeigen`
	FormularauswahlAnzeigen bool        `json:FormularauswahlAnzeigen`
	ZahlungSperren          bool        `json:ZahlungSperren`
	IstTeilrechnung         bool        `json:IstTeilrechnung`
	Text                    Text        `json:Text`
	Parameter               string      `json:Parameter`
}

package golang

type Verbindungsprotokoll struct {
	VerbindungsprotokollNr int    `json:VerbindungsprotokollNr`
	Archivdaten            bool   `json:Archivdaten`
	Bank                   Bank   `json:Bank`
	Datei                  string `json:Datei`
	Datum                  string `json:Datum`
	EsrAbholen             bool   `json:EsrAbholen`
	KontoBestand           bool   `json:KontoBestand`
	Meldung                string `json:Meldung`
	SitzungId              int    `json:SitzungId`
	Typ                    int    `json:Typ`
	Zeit                   string `json:Zeit`
}

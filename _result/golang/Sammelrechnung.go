package golang

type Sammelrechnung struct {
	Dokumente                        int         `json:Dokumente`
	Dokumenttyp                      Dokumenttyp `json:Dokumenttyp`
	UnserZeichen                     string      `json:UnserZeichen`
	Referenztext                     string      `json:Referenztext`
	Kondition                        Kondition   `json:Kondition`
	Datum                            string      `json:Datum`
	Belegdatum                       string      `json:Belegdatum`
	EinzelnVerrechnen                bool        `json:EinzelnVerrechnen`
	AuftragNrNichtBeruecksichtigen   bool        `json:AuftragNrNichtBeruecksichtigen`
	AktivitaetenAufErledigtSetzen    bool        `json:AktivitaetenAufErledigtSetzen`
	RechnungsadressenZusammenfuehren bool        `json:RechnungsadressenZusammenfuehren`
	FusstextUebernehmen              bool        `json:FusstextUebernehmen`
	Startdatum                       string      `json:Startdatum`
	Lieferschein                     Dokument    `json:Lieferschein`
	Rechnung                         Dokument    `json:Rechnung`
	Fehler                           string      `json:Fehler`
}

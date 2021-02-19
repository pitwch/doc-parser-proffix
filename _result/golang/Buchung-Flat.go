package golang

type BuchungFlat struct {
	BuchungNr                      int          `json:BuchungNr`
	BuchungszeileNr                int          `json:BuchungszeileNr`
	Buchungsdatum                  string       `json:Buchungsdatum`
	Belegdatum                     string       `json:Belegdatum`
	Belegnummer                    int          `json:Belegnummer`
	Kondition                      Kondition    `json:Kondition`
	Buchungsart                    Buchungsart  `json:Buchungsart`
	EinzahlungMitteilung           string       `json:EinzahlungMitteilung`
	Verfalldatum                   string       `json:Verfalldatum`
	Adresse                        Adresse      `json:Adresse`
	Zahlungsart                    Zahlungsart  `json:Zahlungsart`
	EsrKodierzeile                 string       `json:EsrKodierzeile`
	EsrPruefziffer                 string       `json:EsrPruefziffer`
	Herkunft                       int          `json:Herkunft`
	Auftrag                        Auftrag      `json:Auftrag`
	Belegart                       Belegart     `json:Belegart`
	EsrNummer                      int          `json:EsrNummer`
	EinzahlungName                 string       `json:EinzahlungName`
	HabenKonto                     Konto        `json:HabenKonto`
	SollKonto                      Konto        `json:SollKonto`
	Waehrung                       Waehrung     `json:Waehrung`
	Kurs                           float64      `json:Kurs`
	MahnCode                       int          `json:MahnCode`
	MahnDatum                      string       `json:MahnDatum`
	Steuercode                     Steuercode   `json:Steuercode`
	SteuerbetragFW                 float64      `json:SteuerbetragFW`
	SteuerbetragSW                 float64      `json:SteuerbetragSW`
	SteuerbuchungFW                float64      `json:SteuerbuchungFW`
	SteuerbuchungSW                float64      `json:SteuerbuchungSW`
	SteuerbetragSaldobesteuerungFW float64      `json:SteuerbetragSaldobesteuerungFW`
	SteuerbetragSaldobesteuerungSW float64      `json:SteuerbetragSaldobesteuerungSW`
	Zahlungsdatum                  string       `json:Zahlungsdatum`
	IstSteuerbuchung               bool         `json:IstSteuerbuchung`
	IstErfassteBuchung             bool         `json:IstErfassteBuchung`
	BetragFW                       float64      `json:BetragFW`
	BetragSW                       float64      `json:BetragSW`
	Buchungstext                   string       `json:Buchungstext`
	Kostenart                      Kostenart    `json:Kostenart`
	Kostenstelle                   Kostenstelle `json:Kostenstelle`
	Belege                         string       `json:Belege`
	Belegstopp                     bool         `json:Belegstopp`
	StopBemerkung                  string       `json:StopBemerkung`
	Hauptbuchung                   int          `json:Hauptbuchung`
	IstErfassteBuchung             string       `json:IstErfassteBuchung`
}

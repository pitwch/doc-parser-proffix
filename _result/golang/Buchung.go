package golang

type Buchung struct {
	BuchungNr                      float64     `json:BuchungNr`
	Buchungsdatum                  string      `json:Buchungsdatum`
	Belegdatum                     string      `json:Belegdatum`
	Belegnummer                    float64     `json:Belegnummer`
	Kondition                      Kondition   `json:Kondition`
	EinzahlungMitteilung           string      `json:EinzahlungMitteilung`
	Verfalldatum                   string      `json:Verfalldatum`
	Adresse                        Adresse     `json:Adresse`
	Zahlungsart                    Zahlungsart `json:Zahlungsart`
	EsrKodierzeile                 string      `json:EsrKodierzeile`
	EsrPruefziffer                 string      `json:EsrPruefziffer`
	Herkunft                       float64     `json:Herkunft`
	Buchungen                      string      `json:Buchungen`
	BuchungszeileNr                float64     `json:BuchungszeileNr`
	Buchungsart                    Buchungsart `json:Buchungsart`
	Auftrag                        Auftrag     `json:Auftrag`
	Belegart                       Belegart    `json:Belegart`
	EsrNummer                      float64     `json:EsrNummer`
	EinzahlungName                 string      `json:EinzahlungName`
	HabenKonto                     Konto       `json:HabenKonto`
	SollKonto                      Konto       `json:SollKonto`
	Waehrung                       Waehrung    `json:Waehrung`
	Kurs                           float64     `json:Kurs`
	MahnCode                       float64     `json:MahnCode`
	MahnDatum                      string      `json:MahnDatum`
	Steuercode                     Steuercode  `json:Steuercode`
	SteuerbetragFW                 float64     `json:SteuerbetragFW`
	SteuerbetragSW                 float64     `json:SteuerbetragSW`
	SteuerbuchungFW                float64     `json:SteuerbuchungFW`
	SteuerbuchungSW                float64     `json:SteuerbuchungSW`
	SteuerbetragSaldobesteuerungFW float64     `json:SteuerbetragSaldobesteuerungFW`
	SteuerbetragSaldobesteuerungSW float64     `json:SteuerbetragSaldobesteuerungSW`
	Zahlungsdatum                  string      `json:Zahlungsdatum`
	IstSteuerbuchung               bool        `json:IstSteuerbuchung`
	IstErfassteBuchung             bool        `json:IstErfassteBuchung`
}

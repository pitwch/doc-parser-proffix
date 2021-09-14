package golang

type Zahlungsart struct {
	ZahlungsartNr   int     `json:"ZahlungsartNr"`
	Adresse         Adresse `json:"Adresse"`
	BankBelegNummer int     `json:"BankBelegNummer"`
	Bank            Bank    `json:"Bank"`
	Beguenstigter   Adresse `json:"Beguenstigter"`
	//BelastenVon       Zahlungsart `json:"BelastenVon"`
	EigeneZahlungsart bool   `json:"EigeneZahlungsart"`
	FormCode          string `json:"FormCode"`
	HauptZahlungsart  bool   `json:"HauptZahlungsart"`
	Inaktiv           bool   `json:"Inaktiv"`
	KontonummerBank   string `json:"KontonummerBank"`
	Konto             Konto  `json:"Konto"`
	PostcheckNummer   string `json:"PostcheckNummer"`
	Spesenregelung    int    `json:"Spesenregelung"`
	SwiftNummer       string `json:"SwiftNummer"`
	TeilnehmerNummer  string `json:"TeilnehmerNummer"`
	ZahlArt           int    `json:"ZahlArt"`
	ZahlungsartTyp    string `json:"ZahlungsartTyp"`
	Bemerkungen       string `json:"Bemerkungen"`
}

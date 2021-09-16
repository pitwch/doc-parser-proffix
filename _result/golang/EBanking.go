package golang

type EBanking struct {
	EBankingNr     int      `json:"EBankingNr"`
	Auszugsnummer  string   `json:"Auszugsnummer"`
	BankBuchungNr  string   `json:"BankBuchungNr"`
	Bank           Bank     `json:"Bank"`
	BelegBild      string   `json:"BelegBild"`
	Betrag         float64  `json:"Betrag"`
	BetragSW       float64  `json:"BetragSW"`
	BuchungNr      int      `json:"BuchungNr"`
	Buchungstext   string   `json:"Buchungstext"`
	Datum          string   `json:"Datum"`
	OpKonto        int      `json:"OpKonto"`
	Status         int      `json:"Status"`
	Saldo          float64  `json:"Saldo"`
	Valuta         string   `json:"Valuta"`
	Waehrung       Waehrung `json:"Waehrung"`
	Zahlungsgrund  string   `json:"Zahlungsgrund"`
	Zahlungsgrund2 string   `json:"Zahlungsgrund2"`
	ZeitBuchungNr  string   `json:"ZeitBuchungNr"`
	Taxen          float64  `json:"Taxen"`
	IstEsr         bool     `json:"IstEsr"`
	IstEs          bool     `json:"IstEs"`
	IsoMsgId       string   `json:"IsoMsgId"`
	IsoPmtInfId    string   `json:"IsoPmtInfId"`
	DetailBetrag   float64  `json:"DetailBetrag"`
	DetailKurs     float64  `json:"DetailKurs"`
	DetailWaehrung Waehrung `json:"DetailWaehrung"`
	IsoAcctSvcrRef string   `json:"IsoAcctSvcrRef"`
	IstCamt        bool     `json:"IstCamt"`
	IsoEndToEndId  string   `json:"IsoEndToEndId"`
}

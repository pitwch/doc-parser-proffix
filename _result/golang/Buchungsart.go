package golang

type Buchungsart struct {
	BuchungsartNr            string      `json:BuchungsartNr`
	Bezeichnung              string      `json:Bezeichnung`
	Bereich                  int         `json:Bereich`
	Typ                      int         `json:Typ`
	BetragsvorschlagFw       float64     `json:BetragsvorschlagFw`
	Belegart                 Belegart    `json:Belegart`
	BuchungVorzeichen        int         `json:BuchungVorzeichen`
	Buchungstextvorschlag    string      `json:Buchungstextvorschlag`
	Eroeffnungsbuchung       bool        `json:Eroeffnungsbuchung`
	Hauptbuchungsart         bool        `json:Hauptbuchungsart`
	Inaktiv                  bool        `json:Inaktiv`
	KeinKontoVorschlag       bool        `json:KeinKontoVorschlag`
	Sammelbuchung            bool        `json:Sammelbuchung`
	SammelbuchungBuchungsart Buchungsart `json:SammelbuchungBuchungsart`
	SkontoRueckbuchung       bool        `json:SkontoRueckbuchung`
	Steuercode               Steuercode  `json:Steuercode`
	SteuerRueckbuchung       bool        `json:SteuerRueckbuchung`
	UmsatzBuchung            bool        `json:UmsatzBuchung`
	BuchungsartenKonten      bool        `json:BuchungsartenKonten`
}

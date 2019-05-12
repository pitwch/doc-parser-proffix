package golang

type Installation struct {
	InstallationNr       int                `json:InstallationNr`
	Standort             Standort           `json:Standort`
	Betreueradresse      Adresse            `json:Betreueradresse`
	Externadresse        Adresse            `json:Externadresse`
	Prioritaet           Prioritaet         `json:Prioritaet`
	Artikel              Artikel            `json:Artikel`
	Bezeichnung1         string             `json:Bezeichnung1`
	Bezeichnung2         string             `json:Bezeichnung2`
	Bezeichnung3         string             `json:Bezeichnung3`
	Bezeichnung4         string             `json:Bezeichnung4`
	Bezeichnung5         string             `json:Bezeichnung5`
	Gewicht              float64            `json:Gewicht`
	Seriennummer         string             `json:Seriennummer`
	KonfigurationsID     string             `json:KonfigurationsID`
	ProduktNr            string             `json:ProduktNr`
	Installationstermin  string             `json:Installationstermin`
	Garantieende         string             `json:Garantieende`
	Garantiezusatz       Garantiezusatz     `json:Garantiezusatz`
	Auftrag              Auftrag            `json:Auftrag`
	Vertragstyp          Vertragstyp        `json:Vertragstyp`
	Herstellervertrag    string             `json:Herstellervertrag`
	Suchfeld             string             `json:Suchfeld`
	Artikelklasse        Artikelklasse      `json:Artikelklasse`
	Artikelgruppe        Artikelgruppe      `json:Artikelgruppe`
	Artikeluntergruppe   Artikeluntergruppe `json:Artikeluntergruppe`
	Menge                string             `json:Menge`
	Anlage               Anlage             `json:Anlage`
	Inhaber              Auftrag            `json:Inhaber`
	IstLagerinstallation bool               `json:IstLagerinstallation`
	Geloescht            bool               `json:Geloescht`
	Global               bool               `json:Global`
	Standard             bool               `json:Standard`
	Parameter            string             `json:Parameter`
}

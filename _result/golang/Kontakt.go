package golang

type Kontakt struct {
	KontaktNr                 int        `json:KontaktNr`
	Kontakttyp                Kontakttyp `json:Kontakttyp`
	Adresse                   Adresse    `json:Adresse`
	Titel                     string     `json:Titel`
	Name                      string     `json:Name`
	Vorname                   string     `json:Vorname`
	Strasse                   string     `json:Strasse`
	PLZ                       string     `json:PLZ`
	Ort                       string     `json:Ort`
	Geloescht                 bool       `json:Geloescht`
	Anrede                    string     `json:Anrede`
	Bemerkungen               string     `json:Bemerkungen`
	EMail                     string     `json:EMail`
	TelDirekt                 string     `json:TelDirekt`
	TelZentrale               string     `json:TelZentrale`
	TelPrivat                 string     `json:TelPrivat`
	Fax                       string     `json:Fax`
	Mobiltelefon              string     `json:Mobiltelefon`
	Homepage                  string     `json:Homepage`
	Region                    Region     `json:Region`
	Land                      Land       `json:Land`
	Hauptadresse              bool       `json:Hauptadresse`
	AdresseGleichHauptadresse bool       `json:AdresseGleichHauptadresse`
	Adresszeile1              string     `json:Adresszeile1`
	Adresszeile2              string     `json:Adresszeile2`
	WebshopBenutzer           string     `json:WebshopBenutzer`
	WebshopEMail              string     `json:WebshopEMail`
	WebshopPasswort           string     `json:WebshopPasswort`
	WebshopPasswortAlt        string     `json:WebshopPasswortAlt`
	WebshopKonditionen        string     `json:WebshopKonditionen`
	Briefanrede               string     `json:Briefanrede`
}

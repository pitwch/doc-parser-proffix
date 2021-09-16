package golang

type Adresse struct {
	AdressNr                            int          `json:"AdressNr"`
	Name                                string       `json:"Name"`
	Vorname                             string       `json:"Vorname"`
	Strasse                             string       `json:"Strasse"`
	PLZ                                 string       `json:"PLZ"`
	Ort                                 string       `json:"Ort"`
	Geloescht                           bool         `json:"Geloescht"`
	Adressgruppen                       string       `json:"Adressgruppen"`
	Anrede                              string       `json:"Anrede"`
	Bemerkungen                         string       `json:"Bemerkungen"`
	EMail                               string       `json:"EMail"`
	TelDirekt                           string       `json:"TelDirekt"`
	TelZentrale                         string       `json:"TelZentrale"`
	TelPrivat                           string       `json:"TelPrivat"`
	Fax                                 string       `json:"Fax"`
	Mobiltelefon                        string       `json:"Mobiltelefon"`
	Homepage                            string       `json:"Homepage"`
	PLZPostfach                         string       `json:"PLZPostfach"`
	Postfach                            string       `json:"Postfach"`
	Region                              Region       `json:"Region"`
	Land                                Land         `json:"Land"`
	Longitude                           float64      `json:"Longitude"`
	Latitude                            float64      `json:"Latitude"`
	DebitorenSteuercode                 Steuercode   `json:"DebitorenSteuercode"`
	DebitorenSteuercodeVerwenden        bool         `json:"DebitorenSteuercodeVerwenden"`
	DebitorenErtragskonto               Konto        `json:"DebitorenErtragskonto"`
	DebitorenErtragskontoVerwenden      bool         `json:"DebitorenErtragskontoVerwenden"`
	DebitorenKostenstelle               Kostenstelle `json:"DebitorenKostenstelle"`
	DebitorenKostenstelleVerwenden      bool         `json:"DebitorenKostenstelleVerwenden"`
	DebitorenKostenart                  Kostenart    `json:"DebitorenKostenart"`
	DebitorenKostenartVerwenden         bool         `json:"DebitorenKostenartVerwenden"`
	DebitorenRabatt                     float64      `json:"DebitorenRabatt"`
	Vertreter                           Vertreter    `json:"Vertreter"`
	DebitorenWaehrung                   Waehrung     `json:"DebitorenWaehrung"`
	Lagerpreis                          int          `json:"Lagerpreis"`
	Geburtsdatum                        string       `json:"Geburtsdatum"`
	Adresszeile1                        string       `json:"Adresszeile1"`
	Adresszeile2                        string       `json:"Adresszeile2"`
	Adresszeile3                        string       `json:"Adresszeile3"`
	Adresszeile4                        string       `json:"Adresszeile4"`
	Adresszeile5                        string       `json:"Adresszeile5"`
	WebshopBenutzer                     string       `json:"WebshopBenutzer"`
	WebshopEMail                        string       `json:"WebshopEMail"`
	WebshopPasswort                     string       `json:"WebshopPasswort"`
	WebshopPasswortAlt                  string       `json:"WebshopPasswortAlt"`
	WebshopKonditionen                  string       `json:"WebshopKonditionen"`
	Lieferadresse                       Adresse      `json:"Lieferadresse"`
	Rechnungsadresse                    Adresse      `json:"Rechnungsadresse"`
	Sprache                             Sprache      `json:"Sprache"`
	Suchfeld                            string       `json:"Suchfeld"`
	DebitorenSammelkonto                Konto        `json:"DebitorenSammelkonto"`
	KreditorenSammelkonto               Konto        `json:"KreditorenSammelkonto"`
	Briefanrede                         string       `json:"Briefanrede"`
	GLN                                 string       `json:"GLN"`
	DebitorenKondition                  Kondition    `json:"DebitorenKondition"`
	DebitorenAuftrag                    Auftrag      `json:"DebitorenAuftrag"`
	DebitorenKreditlimite               float64      `json:"DebitorenKreditlimite"`
	DebitorenMWSTID                     string       `json:"DebitorenMWSTID"`
	DebitorenMWSTUID                    string       `json:"DebitorenMWSTUID"`
	Kleinmengenzuschlag                 Zuschlag     `json:"Kleinmengenzuschlag"`
	ESRNummer                           ESRNummer    `json:"ESRNummer"`
	Buchungssperre                      bool         `json:"Buchungssperre"`
	Photo                               string       `json:"Photo"`
	LSV                                 bool         `json:"LSV"`
	DebitorenLieferart                  Lieferart    `json:"DebitorenLieferart"`
	KreditorenAufwandkonto              Konto        `json:"KreditorenAufwandkonto"`
	KreditorenBuchungssperre            bool         `json:"KreditorenBuchungssperre"`
	KreditorenKostenart                 Kostenart    `json:"KreditorenKostenart"`
	KreditorenKonditionen               Kondition    `json:"KreditorenKonditionen"`
	KreditorenKostenstelle              Kostenstelle `json:"KreditorenKostenstelle"`
	KreditorenKundenNrBeiLieferant      string       `json:"KreditorenKundenNrBeiLieferant"`
	KreditorenRabatt                    float64      `json:"KreditorenRabatt"`
	KreditorenSteuercodeVerwenden       bool         `json:"KreditorenSteuercodeVerwenden"`
	KreditorenMWSTIDNummer              string       `json:"KreditorenMWSTIDNummer"`
	KreditorenUIDMWST                   string       `json:"KreditorenUIDMWST"`
	KreditorenWaehrung                  Waehrung     `json:"KreditorenWaehrung"`
	KreditorenBelegVorErfassen          bool         `json:"KreditorenBelegVorErfassen"`
	KreditorenEinkaufsSperre            bool         `json:"KreditorenEinkaufsSperre"`
	KreditorenVerguetungsSperre         bool         `json:"KreditorenVerguetungsSperre"`
	KreditorenVerknuepfungMitBestellung bool         `json:"KreditorenVerknuepfungMitBestellung"`
	KreditorenKstAufteilung1            string       `json:"KreditorenKstAufteilung1"`
	KreditorenKstAufteilung2            string       `json:"KreditorenKstAufteilung2"`
	KreditorenKstAufteilung3            string       `json:"KreditorenKstAufteilung3"`
	KreditorenKstAufteilung4            string       `json:"KreditorenKstAufteilung4"`
	KreditorenKstAufteilung5            string       `json:"KreditorenKstAufteilung5"`
	KreditorenKstAufteilung6            string       `json:"KreditorenKstAufteilung6"`
	KreditorenKstAufteilung7            string       `json:"KreditorenKstAufteilung7"`
	KreditorenKstAufteilung8            string       `json:"KreditorenKstAufteilung8"`
	KreditorenKstAufteilung9            string       `json:"KreditorenKstAufteilung9"`
	KreditorenKstAufteilung10           string       `json:"KreditorenKstAufteilung10"`
	Parameter                           string       `json:"Parameter"`
}

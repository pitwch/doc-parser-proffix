package golang

type Ressource struct {
	RessourceNr     int          `json:RessourceNr`
	Anzahl          int          `json:Anzahl`
	Artikel         Artikel      `json:Artikel`
	Aufgabe         Aufgabe      `json:Aufgabe`
	Bezeichnung     string       `json:Bezeichnung`
	Einheit         Einheit      `json:Einheit`
	Ertragskonto    Konto        `json:Ertragskonto`
	Eskalation      int          `json:Eskalation`
	Planobjektfarbe int          `json:Planobjektfarbe`
	Geloescht       bool         `json:Geloescht`
	Kostenart       Kostenart    `json:Kostenart`
	Kostenstelle    Kostenstelle `json:Kostenstelle`
	Mindestdauer    int          `json:Mindestdauer`
	Mitarbeiter     Mitarbeiter  `json:Mitarbeiter`
	Nachlaufzeit    int          `json:Nachlaufzeit`
	Installation    Installation `json:Installation`
	Preiseinheit    string       `json:Preiseinheit`
	PreisFW         float64      `json:PreisFW`
	PreisSW         float64      `json:PreisSW`
	Steuercode      Steuercode   `json:Steuercode`
	Verrechnen      bool         `json:Verrechnen`
	Vorlaufzeit     int          `json:Vorlaufzeit`
	Waehrung        Waehrung     `json:Waehrung`
	Parameter       string       `json:Parameter`
}

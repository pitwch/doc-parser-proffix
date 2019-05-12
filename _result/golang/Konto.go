package golang

type Konto struct {
	KontoNr                       string       `json:KontoNr`
	Bezeichnung                   string       `json:Bezeichnung`
	Steuercode                    Steuercode   `json:Steuercode`
	Kontoklasse                   Kontoklasse  `json:Kontoklasse`
	Kostenstelle                  Kostenstelle `json:Kostenstelle`
	KostenstelleZwingend          int          `json:KostenstelleZwingend`
	KostenstelleFix               bool         `json:KostenstelleFix`
	Kostenart                     Kostenart    `json:Kostenart`
	KostenartZwingend             int          `json:KostenartZwingend`
	KostenartFix                  bool         `json:KostenartFix`
	Waehrung                      Waehrung     `json:Waehrung`
	Auftrag                       Auftrag      `json:Auftrag`
	AuftragZwingend               int          `json:AuftragZwingend`
	AuftragFix                    bool         `json:AuftragFix`
	Lohnart                       Lohnart      `json:Lohnart`
	LohnartZwingend               int          `json:LohnartZwingend`
	LohnartFix                    bool         `json:LohnartFix`
	BuchenInFinanzbuchhaltung     bool         `json:BuchenInFinanzbuchhaltung`
	BuchenInDebitorenbuchhaltung  bool         `json:BuchenInDebitorenbuchhaltung`
	BuchenInKreditorenbuchhaltung bool         `json:BuchenInKreditorenbuchhaltung`
	BuchenInAuftragsbearbeitung   bool         `json:BuchenInAuftragsbearbeitung`
	BuchenInLeistungsverwaltung   bool         `json:BuchenInLeistungsverwaltung`
	BuchenInAnlagekonto           bool         `json:BuchenInAnlagekonto`
	Sperre                        string       `json:Sperre`
}

package golang

type Artikel struct {
	ArtikelNr                      string             `json:"ArtikelNr"`
	Bezeichnung1                   string             `json:"Bezeichnung1"`
	Bezeichnung2                   string             `json:"Bezeichnung2"`
	Bezeichnung3                   string             `json:"Bezeichnung3"`
	Bezeichnung4                   string             `json:"Bezeichnung4"`
	Bezeichnung5                   string             `json:"Bezeichnung5"`
	Geloescht                      bool               `json:"Geloescht"`
	Waehrung                       Waehrung           `json:"Waehrung"`
	EinheitLager                   Einheit            `json:"EinheitLager"`
	EinheitRechnung                Einheit            `json:"EinheitRechnung"`
	Dim1                           float64            `json:"Dim1"`
	Dim2                           float64            `json:"Dim2"`
	Dim3                           float64            `json:"Dim3"`
	Verkaufspreis1                 float64            `json:"Verkaufspreis1"`
	Verkaufspreis2                 float64            `json:"Verkaufspreis2"`
	Verkaufspreis3                 float64            `json:"Verkaufspreis3"`
	Verkaufspreis4                 float64            `json:"Verkaufspreis4"`
	Verkaufspreis5                 float64            `json:"Verkaufspreis5"`
	Steuercode                     Steuercode         `json:"Steuercode"`
	Ertragskonto                   Konto              `json:"Ertragskonto"`
	Kostenstelle                   Kostenstelle       `json:"Kostenstelle"`
	Kostenart                      Kostenart          `json:"Kostenart"`
	KeinRabatt                     bool               `json:"KeinRabatt"`
	Barcode                        string             `json:"Barcode"`
	Lagerort                       Lagerort           `json:"Lagerort"`
	Lagerplatz                     Lagerplatz         `json:"Lagerplatz"`
	SerieNrVerwenden               bool               `json:"SerieNrVerwenden"`
	ChargeVerwenden                bool               `json:"ChargeVerwenden"`
	Steuercode1                    Steuercode         `json:"Steuercode1"`
	Steuercode2                    Steuercode         `json:"Steuercode2"`
	Steuercode3                    Steuercode         `json:"Steuercode3"`
	Steuercode4                    Steuercode         `json:"Steuercode4"`
	Steuercode5                    Steuercode         `json:"Steuercode5"`
	Auftrag                        Auftrag            `json:"Auftrag"`
	Aufwandskonto                  Konto              `json:"Aufwandskonto"`
	BasisInstallieren              bool               `json:"BasisInstallieren"`
	BasisNurEinmalInstallieren     bool               `json:"BasisNurEinmalInstallieren"`
	Bestandeskonto                 Konto              `json:"Bestandeskonto"`
	Bestellpunkt                   float64            `json:"Bestellpunkt"`
	Bestellrhythmus                float64            `json:"Bestellrhythmus"`
	Bestellvorschlag               int                `json:"Bestellvorschlag"`
	BasispreisProzentsatzBewertung int                `json:"BasispreisProzentsatzBewertung"`
	Bewertungspreis                float64            `json:"Bewertungspreis"`
	ProzentsatzBewertung           float64            `json:"ProzentsatzBewertung"`
	GTIN                           string             `json:"GTIN"`
	GueltigVonEinkauf              string             `json:"GueltigVonEinkauf"`
	GueltigBisEinkauf              string             `json:"GueltigBisEinkauf"`
	Einkaufspreis                  float64            `json:"Einkaufspreis"`
	SteuercodeEinkauf              Steuercode         `json:"SteuercodeEinkauf"`
	Einstandspreis                 float64            `json:"Einstandspreis"`
	Ersatzartikel                  Artikel            `json:"Ersatzartikel"`
	Faktor1                        float64            `json:"Faktor1"`
	Faktor2                        float64            `json:"Faktor2"`
	Faktor3                        float64            `json:"Faktor3"`
	Faktor4                        float64            `json:"Faktor4"`
	Faktor5                        float64            `json:"Faktor5"`
	FaktorBewertung                float64            `json:"FaktorBewertung"`
	Faktor1Fixiert                 bool               `json:"Faktor1Fixiert"`
	Faktor2Fixiert                 bool               `json:"Faktor2Fixiert"`
	Faktor3Fixiert                 bool               `json:"Faktor3Fixiert"`
	Faktor4Fixiert                 bool               `json:"Faktor4Fixiert"`
	Faktor5Fixiert                 bool               `json:"Faktor5Fixiert"`
	FaktorBewertungFixiert         bool               `json:"FaktorBewertungFixiert"`
	Garantie                       int                `json:"Garantie"`
	Gebindeeinheit                 Einheit            `json:"Gebindeeinheit"`
	Gewicht                        float64            `json:"Gewicht"`
	Artikelgruppe                  Artikelgruppe      `json:"Artikelgruppe"`
	Hersteller                     Hersteller         `json:"Hersteller"`
	KeinBestand                    bool               `json:"KeinBestand"`
	KeinKassenverkauf              bool               `json:"KeinKassenverkauf"`
	Artikelklasse                  Artikelklasse      `json:"Artikelklasse"`
	Lagerartikel                   bool               `json:"Lagerartikel"`
	LetzterEinkaufspreis           float64            `json:"LetzterEinkaufspreis"`
	Minimalbestand                 float64            `json:"Minimalbestand"`
	Maximalbestand                 float64            `json:"Maximalbestand"`
	Mindestbestellmenge            float64            `json:"Mindestbestellmenge"`
	Negativbestand                 bool               `json:"Negativbestand"`
	Optimalbestellmenge            float64            `json:"Optimalbestellmenge"`
	Packungseinheit                float64            `json:"Packungseinheit"`
	MitPackungseinheitRechnen      bool               `json:"MitPackungseinheitRechnen"`
	PreisAufAnfrageWebshop         bool               `json:"PreisAufAnfrageWebshop"`
	Sammelartikel                  Sammelartikel      `json:"Sammelartikel"`
	Webshop                        bool               `json:"Webshop"`
	AktionWebshop                  bool               `json:"AktionWebshop"`
	WebshopbildErstellt            bool               `json:"WebshopbildErstellt"`
	Sicherheitstage                int                `json:"Sicherheitstage"`
	StuecklisteAufloesen           bool               `json:"StuecklisteAufloesen"`
	StuecklisteBuendeln            bool               `json:"StuecklisteBuendeln"`
	IstStuecklistenkopf            bool               `json:"IstStuecklistenkopf"`
	StuecklistenpreiseBerechnen    bool               `json:"StuecklistenpreiseBerechnen"`
	Tara                           float64            `json:"Tara"`
	Artikeluntergruppe             Artikeluntergruppe `json:"Artikeluntergruppe"`
	GueltigVonVerkauf              string             `json:"GueltigVonVerkauf"`
	GueltigBisVerkauf              string             `json:"GueltigBisVerkauf"`
	Verpackung                     Verpackung         `json:"Verpackung"`
	Basispreis1                    int                `json:"Basispreis1"`
	Basispreis2                    int                `json:"Basispreis2"`
	Basispreis3                    int                `json:"Basispreis3"`
	Basispreis4                    int                `json:"Basispreis4"`
	Basispreis5                    int                `json:"Basispreis5"`
	BasispreisBewertung            int                `json:"BasispreisBewertung"`
	Wiederbeschaffung              int                `json:"Wiederbeschaffung"`
	Suchfeld                       string             `json:"Suchfeld"`
	Ursprungsland                  Land               `json:"Ursprungsland"`
	TarifNr                        float64            `json:"TarifNr"`
	Tarifschlüssel                 int                `json:"Tarifschlüssel"`
}

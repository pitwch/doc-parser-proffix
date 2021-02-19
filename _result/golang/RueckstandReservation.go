package golang

type RueckstandReservation struct {
	RueckstandReservationNr int              `json:RueckstandReservationNr`
	Adresse                 Adresse          `json:Adresse`
	Anzahl                  float64          `json:Anzahl`
	Artikel                 Artikel          `json:Artikel `
	Bemerkung               string           `json:Bemerkung`
	BemerkungRTF            string           `json:BemerkungRTF`
	Bestellmenge            float64          `json:Bestellmenge`
	Bestellung              Bestellung       `json:Bestellung`
	Charge                  Charge           `json:Charge`
	Dokument                Dokument         `json:Dokument`
	DokumentpositionNr      Dokumentposition `json:DokumentpositionNr`
	ErstelltAm              string           `json:ErstelltAm`
	ErstelltVon             string           `json:ErstelltVon`
	GeaendertAm             string           `json:GeaendertAm`
	GeaendertVon            string           `json:GeaendertVon`
	GtinStufe               string           `json:GtinStufe`
	Lagerort                Lagerort         `json:Lagerort`
	Lagerplatz              Lagerplatz       `json:Lagerplatz`
	Liefertermin            string           `json:Liefertermin`
	Menge                   float64          `json:Menge`
	ReservierenAb           string           `json:ReservierenAb`
	Typ                     int              `json:Typ`
	UngueltigAb             string           `json:UngueltigAb`
	Verarbeitet             bool             `json:Verarbeitet`
	Verpackungen            int              `json:Verpackungen`
	Woche                   string           `json:Woche`
	Parameter               string           `json:Parameter`
}

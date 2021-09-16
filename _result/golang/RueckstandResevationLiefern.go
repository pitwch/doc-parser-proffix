package golang

type RueckstandResevationLiefern struct {
	AktivitaetenAufErledigtSetzen bool                  `json:"AktivitaetenAufErledigtSetzen"`
	Belegdatum                    string                `json:"Belegdatum"`
	Datum                         string                `json:"Datum"`
	Dokumenttyp                   Dokumenttyp           `json:"Dokumenttyp"`
	EinzelnLiefern                bool                  `json:"EinzelnLiefern"`
	RueckstandReservation         string                `json:"RueckstandReservation"`
	Startdatum                    string                `json:"Startdatum"`
	UnserZeichen                  string                `json:"UnserZeichen"`
	Lagerbewegungen               Lagerbewegungen       `json:"Lagerbewegungen"`
	RueckstandReservation         RueckstandReservation `json:"RueckstandReservation"`
	Fehler                        string                `json:"Fehler"`
	Lieferschein                  Dokumenttyp           `json:"Lieferschein"`
	RueckstandReservationen       string                `json:"RueckstandReservationen"`
}

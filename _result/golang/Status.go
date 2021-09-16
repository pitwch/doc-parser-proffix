package golang

type Status struct {
	StatusNr               string `json:"StatusNr"`
	Bezeichnung            string `json:"Bezeichnung"`
	Dateien                string `json:"Dateien"`
	IstEskaliert           bool   `json:"IstEskaliert"`
	Farbe                  int    `json:"Farbe"`
	InstallationDokDrucken bool   `json:"InstallationDokDrucken"`
	NichtEskalieren        bool   `json:"NichtEskalieren"`
	MeldungIntern          bool   `json:"MeldungIntern"`
	MeldungExtern          bool   `json:"MeldungExtern"`
	Reihenfolge            int    `json:"Reihenfolge"`
	ServiceDokDrucken      bool   `json:"ServiceDokDrucken"`
	ShowDialog             bool   `json:"ShowDialog"`
	Sperren                bool   `json:"Sperren"`
	Statusart              int    `json:"Statusart"`
}

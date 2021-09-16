package golang

type Kinderzulage struct {
	KinderzulageNr             int     `json:"KinderzulageNr"`
	Adresse                    Adresse `json:"Adresse"`
	Ausbildungszulage1Betrag   float64 `json:"Ausbildungszulage1Betrag"`
	Ausbildungszulage1AlterVon int     `json:"Ausbildungszulage1AlterVon"`
	Ausbildungszulage1AlterBis int     `json:"Ausbildungszulage1AlterBis"`
	Ausbildungszulage2Betrag   float64 `json:"Ausbildungszulage2Betrag"`
	Ausbildungszulage2AlterVon int     `json:"Ausbildungszulage2AlterVon"`
	Ausbildungszulage2AlterBis int     `json:"Ausbildungszulage2AlterBis"`
	Beitragssatz               float64 `json:"Beitragssatz"`
	Region                     Region  `json:"Region"`
	MitgliedNr                 string  `json:"MitgliedNr"`
	Nummer                     string  `json:"Nummer"`
	Subnummer                  string  `json:"Subnummer"`
	Kinderzulage1Betrag        float64 `json:"Kinderzulage1Betrag"`
	Kinderzulage1AlterVon      int     `json:"Kinderzulage1AlterVon"`
	Kinderzulage1AlterBis      int     `json:"Kinderzulage1AlterBis"`
	Kinderzulage2Betrag        float64 `json:"Kinderzulage2Betrag"`
	Kinderzulage2AlterVon      int     `json:"Kinderzulage2AlterVon"`
	Kinderzulage2AlterBis      int     `json:"Kinderzulage2AlterBis"`
}

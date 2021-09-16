package golang

type Planobjekt struct {
	PlanobjektNr              int              `json:"PlanobjektNr"`
	Adresse                   Adresse          `json:"Adresse"`
	Aufgabe                   Aufgabe          `json:"Aufgabe"`
	Auftrag                   Auftrag          `json:"Auftrag"`
	Serviceauftrag            Serviceauftrag   `json:"Serviceauftrag"`
	Beginn                    string           `json:"Beginn"`
	Bemerkungen               string           `json:"Bemerkungen"`
	Bezeichnung               string           `json:"Bezeichnung"`
	Rechnungsdatum            string           `json:"Rechnungsdatum"`
	Dokument                  Dokument         `json:"Dokument"`
	Ende                      string           `json:"Ende"`
	Farbe                     int              `json:"Farbe"`
	Ganztaegig                bool             `json:"Ganztaegig"`
	Kontakt                   Kontakt          `json:"Kontakt"`
	Menge                     int              `json:"Menge"`
	PlanobjektgruppenNr       int              `json:"PlanobjektgruppenNr"`
	ServiceauftragspositionNr int              `json:"ServiceauftragspositionNr"`
	Ressource                 Ressource        `json:"Ressource"`
	Planobjektstatus          Planobjektstatus `json:"Planobjektstatus"`
	Planobjektgruppe          int              `json:"Planobjektgruppe"`
	Parameter                 string           `json:"Parameter"`
}

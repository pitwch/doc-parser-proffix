package golang

type Stundenantraege struct {
	StundenNr   int         `json:"StundenNr"`
	Mitarbeiter Mitarbeiter `json:"Mitarbeiter"`
	Stundenart  Stundenart  `json:"Stundenart"`
	Datum       string      `json:"Datum"`
	StartZeit   string      `json:"StartZeit"`
	EndZeit     string      `json:"EndZeit"`
	Stunden     float64     `json:"Stunden"`
	Pause       float64     `json:"Pause"`
	KaWo        string      `json:"KaWo"`
	Quelle      string      `json:"Quelle"`
	Bemerkungen string      `json:"Bemerkungen"`
	Longitude   float64     `json:"Longitude"`
	Latitude    float64     `json:"Latitude"`
	Accuracy    float64     `json:"Accuracy"`
	Status      int         `json:"Status"`
}

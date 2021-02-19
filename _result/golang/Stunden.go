package golang

type Stunden struct {
	StundenNr     int         `json:StundenNr`
	Mitarbeiter   Mitarbeiter `json:Mitarbeiter`
	Stundenart    Stundenart  `json:Stundenart`
	Datum         string      `json:Datum`
	StartZeit     string      `json:StartZeit`
	EndZeit       string      `json:EndZeit`
	Stunden       float64     `json:Stunden`
	Pause         float64     `json:Pause`
	KaWo          string      `json:KaWo`
	Quelle        string      `json:Quelle`
	QuelleStop    string      `json:QuelleStop`
	Bemerkungen   string      `json:Bemerkungen`
	Longitude     float64     `json:Longitude`
	Latitude      float64     `json:Latitude`
	Accuracy      float64     `json:Accuracy`
	LongitudeStop float64     `json:LongitudeStop`
	LatitudeStop  float64     `json:LatitudeStop`
	AccuracyStop  float64     `json:AccuracyStop`
	Kuerzung      bool        `json:Kuerzung`
	KuerzungMonat bool        `json:KuerzungMonat`
	Gesperrt      bool        `json:Gesperrt`
	Status        int         `json:Status`
}

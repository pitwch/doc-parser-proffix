package golang

type Dokument struct {
	DokumentNr                      float64         `json:DokumentNr`
	Adresse                         Adresse         `json:Adresse`
	Adresszeile1                    string          `json:Adresszeile1`
	Adresszeile2                    string          `json:Adresszeile2`
	Adresszeile3                    string          `json:Adresszeile3`
	Adresszeile4                    string          `json:Adresszeile4`
	Adresszeile5                    string          `json:Adresszeile5`
	Anrede                          string          `json:Anrede`
	Auftrag                         Auftrag         `json:Auftrag`
	Buchung                         Buchung         `json:Buchung`
	Datum                           string          `json:Datum`
	Dokumentstatus                  Dokumentstatus  `json:Dokumentstatus`
	Dokumenttyp                     Dokumenttyp     `json:Dokumenttyp`
	ESRNummer                       ESRNummer       `json:ESRNummer`
	Fusstext                        string          `json:Fusstext`
	FusstextRTF                     string          `json:FusstextRTF`
	Gedruckt                        bool            `json:Gedruckt`
	Gesperrt                        bool            `json:Gesperrt`
	Gueltigkeit                     string          `json:Gueltigkeit`
	IhrZeichen                      string          `json:IhrZeichen`
	Region                          Region          `json:Region`
	Kondition                       Kondition       `json:Kondition`
	Kontaktname                     string          `json:Kontaktname`
	Kontakt                         Kontakt         `json:Kontakt`
	Kurs                            float64         `json:Kurs`
	Land                            Land            `json:Land`
	Lieferadresszeile1              string          `json:Lieferadresszeile1`
	Lieferadresszeile2              string          `json:Lieferadresszeile2`
	Lieferadresszeile3              string          `json:Lieferadresszeile3`
	Lieferadresszeile4              string          `json:Lieferadresszeile4`
	Lieferadresszeile5              string          `json:Lieferadresszeile5`
	Lieferadresse                   Adresse         `json:Lieferadresse`
	LieferadressAnrede              string          `json:LieferadressAnrede`
	Lieferart                       Lieferart       `json:Lieferart`
	LieferadressEMail               string          `json:LieferadressEMail`
	LieferadressRegion              Region          `json:LieferadressRegion`
	LieferadressKontaktname         string          `json:LieferadressKontaktname`
	LieferadressKontakt             Kontakt         `json:LieferadressKontakt`
	LieferadressLand                Land            `json:LieferadressLand`
	LieferadressName                string          `json:LieferadressName`
	LieferadressOrt                 string          `json:LieferadressOrt`
	LieferadressPLZ                 string          `json:LieferadressPLZ`
	LieferadressPLZPostfach         string          `json:LieferadressPLZPostfach`
	LieferadressPostfach            string          `json:LieferadressPostfach`
	LieferadressStrasse             string          `json:LieferadressStrasse`
	LieferadressTelefon             string          `json:LieferadressTelefon`
	Liefertermin                    string          `json:Liefertermin`
	LieferadressVorname             string          `json:LieferadressVorname`
	Marge                           float64         `json:Marge`
	Name                            string          `json:Name`
	Ort                             string          `json:Ort`
	PLZ                             string          `json:PLZ`
	PLZPostfach                     string          `json:PLZPostfach`
	PortoErtragskonto               Konto           `json:PortoErtragskonto`
	PortoFW                         float64         `json:PortoFW`
	PortoKostenart                  Kostenart       `json:PortoKostenart`
	PortoKostenstelle               Kostenstelle    `json:PortoKostenstelle`
	PortoMWStFW                     float64         `json:PortoMWStFW`
	PortoMWStSW                     float64         `json:PortoMWStSW`
	PortoSteuercode                 Steuercode      `json:PortoSteuercode`
	PortoSW                         float64         `json:PortoSW`
	Postfach                        string          `json:Postfach`
	Belegdatum                      string          `json:Belegdatum`
	Rechnungsadresszeile1           string          `json:Rechnungsadresszeile1`
	Rechnungsadresszeile2           string          `json:Rechnungsadresszeile2`
	Rechnungsadresszeile3           string          `json:Rechnungsadresszeile3`
	Rechnungsadresszeile4           string          `json:Rechnungsadresszeile4`
	Rechnungsadresszeile5           string          `json:Rechnungsadresszeile5`
	Rechnungsadresse                Adresse         `json:Rechnungsadresse`
	RechnungsadressAnrede           string          `json:RechnungsadressAnrede`
	RechnungsadressRegion           Region          `json:RechnungsadressRegion`
	RechnungsadressKontaktname      string          `json:RechnungsadressKontaktname`
	RechnungsadressKontakt          Kontakt         `json:RechnungsadressKontakt`
	RechnungsadressLand             Land            `json:RechnungsadressLand`
	RechnungsadressName             string          `json:RechnungsadressName`
	RechnungsadressOrt              string          `json:RechnungsadressOrt`
	RechnungsadressPLZ              string          `json:RechnungsadressPLZ`
	RechnungsadressPLZPostfach      string          `json:RechnungsadressPLZPostfach`
	RechnungsadressPostfach         string          `json:RechnungsadressPostfach`
	RechnungsadressStrasse          string          `json:RechnungsadressStrasse`
	RechnungsadressVorname          string          `json:RechnungsadressVorname`
	Referenztext                    string          `json:Referenztext`
	Sammelkonto                     Konto           `json:Sammelkonto`
	Stapelrechnung                  Dokument        `json:Stapelrechnung`
	SteuerFW1                       float64         `json:SteuerFW1`
	SteuerFW2                       float64         `json:SteuerFW2`
	SteuerFW3                       float64         `json:SteuerFW3`
	SteuerFW4                       float64         `json:SteuerFW4`
	SteuerFW5                       float64         `json:SteuerFW5`
	Steuercode1                     Steuercode      `json:Steuercode1`
	Steuercode2                     Steuercode      `json:Steuercode2`
	Steuercode3                     Steuercode      `json:Steuercode3`
	Steuercode4                     Steuercode      `json:Steuercode4`
	Steuercode5                     Steuercode      `json:Steuercode5`
	SteuerInklusiv1                 bool            `json:SteuerInklusiv1`
	SteuerInklusiv2                 bool            `json:SteuerInklusiv2`
	SteuerInklusiv3                 bool            `json:SteuerInklusiv3`
	SteuerInklusiv4                 bool            `json:SteuerInklusiv4`
	SteuerInklusiv5                 bool            `json:SteuerInklusiv5`
	SteuerMargen1                   bool            `json:SteuerMargen1`
	SteuerMargen2                   bool            `json:SteuerMargen2`
	SteuerMargen3                   bool            `json:SteuerMargen3`
	SteuerMargen4                   bool            `json:SteuerMargen4`
	SteuerMargen5                   bool            `json:SteuerMargen5`
	SteuerSW1                       float64         `json:SteuerSW1`
	SteuerSW2                       float64         `json:SteuerSW2`
	SteuerSW3                       float64         `json:SteuerSW3`
	SteuerSW4                       float64         `json:SteuerSW4`
	SteuerSW5                       float64         `json:SteuerSW5`
	Steuertext1                     string          `json:Steuertext1`
	Steuertext2                     string          `json:Steuertext2`
	Steuertext3                     string          `json:Steuertext3`
	Steuertext4                     string          `json:Steuertext4`
	Steuertext5                     string          `json:Steuertext5`
	SteuertotalFW1                  float64         `json:SteuertotalFW1`
	SteuertotalFW2                  float64         `json:SteuertotalFW2`
	SteuertotalFW3                  float64         `json:SteuertotalFW3`
	SteuertotalFW4                  float64         `json:SteuertotalFW4`
	SteuertotalFW5                  float64         `json:SteuertotalFW5`
	SteuertotalSW1                  float64         `json:SteuertotalSW1`
	SteuertotalSW2                  float64         `json:SteuertotalSW2`
	SteuertotalSW3                  float64         `json:SteuertotalSW3`
	SteuertotalSW4                  float64         `json:SteuertotalSW4`
	SteuertotalSW5                  float64         `json:SteuertotalSW5`
	Strasse                         string          `json:Strasse`
	SubtotaldifferenzFW             float64         `json:SubtotaldifferenzFW`
	SubtotaldifferenzSW             float64         `json:SubtotaldifferenzSW`
	TotalEinkaufspreisFW            float64         `json:TotalEinkaufspreisFW`
	TotalEinkaufspreisSW            float64         `json:TotalEinkaufspreisSW`
	TotalFW                         float64         `json:TotalFW`
	Totalgewicht                    float64         `json:Totalgewicht`
	TotalsteuerFW                   float64         `json:TotalsteuerFW`
	TotalsteuerSW                   float64         `json:TotalsteuerSW`
	TotalSW                         float64         `json:TotalSW`
	UnserZeichen                    string          `json:UnserZeichen`
	Vertreter                       Vertreter       `json:Vertreter`
	Verwaltungsadresse              Adresse         `json:Verwaltungsadresse`
	Vorname                         string          `json:Vorname`
	Zahlungen                       string          `json:Zahlungen`
	Vorauszahlungen                 string          `json:Vorauszahlungen`
	Waehrung                        Waehrung        `json:Waehrung`
	Zuschlag                        Zuschlag        `json:Zuschlag`
	ZuschlagErtragskonto            Konto           `json:ZuschlagErtragskonto`
	ZuschlagFW                      float64         `json:ZuschlagFW`
	ZuschlagKostenart               Kostenart       `json:ZuschlagKostenart`
	ZuschlagKostenstelle            Kostenstelle    `json:ZuschlagKostenstelle`
	ZuschlagMWStFW                  float64         `json:ZuschlagMWStFW`
	ZuschlagMWStSW                  float64         `json:ZuschlagMWStSW`
	ZuschlagSteuercode              Steuercode      `json:ZuschlagSteuercode`
	ZuschlagSW                      float64         `json:ZuschlagSW`
	Positionen                      string          `json:Positionen`
	PositionNr                      float64         `json:PositionNr`
	Artikel                         Artikel         `json:Artikel`
	Auftrag                         Auftrag         `json:Auftrag`
	Serviceauftrag                  Serviceauftrag  `json:Serviceauftrag`
	Bezeichnung1                    string          `json:Bezeichnung1`
	Bezeichnung2                    string          `json:Bezeichnung2`
	Bezeichnung3                    string          `json:Bezeichnung3`
	Bezeichnung4                    string          `json:Bezeichnung4`
	Bezeichnung5                    string          `json:Bezeichnung5`
	Bild                            string          `json:Bild`
	Dim2                            float64         `json:Dim2`
	Eingangsstempel                 string          `json:Eingangsstempel`
	Lagereinheit                    Einheit         `json:Lagereinheit`
	Rechnungseinheit                Einheit         `json:Rechnungseinheit`
	EinkaufspreisFW                 float64         `json:EinkaufspreisFW`
	EinkaufspreisSW                 float64         `json:EinkaufspreisSW`
	Ertragskonto                    Konto           `json:Ertragskonto`
	Gebinde                         float64         `json:Gebinde`
	Gewicht                         float64         `json:Gewicht`
	GtinStufe                       string          `json:GtinStufe`
	Gutschein                       Gutschein       `json:Gutschein`
	Inhalt                          float64         `json:Inhalt`
	IstOption                       bool            `json:IstOption`
	Kundenbestellnummer             string          `json:Kundenbestellnummer`
	Kundenbestellposition           string          `json:Kundenbestellposition`
	Kostenart                       Kostenart       `json:Kostenart`
	Kostenstelle                    Kostenstelle    `json:Kostenstelle`
	Dim1                            float64         `json:Dim1`
	Lagerpreis                      float64         `json:Lagerpreis`
	Lieferant                       Adresse         `json:Lieferant`
	Liefertermin                    string          `json:Liefertermin`
	Marge                           float64         `json:Marge`
	Menge                           float64         `json:Menge`
	MengeOptional                   float64         `json:MengeOptional`
	MengeVerr                       float64         `json:MengeVerr`
	MengeVerrOptional               float64         `json:MengeVerrOptional`
	OccasionsObjektVerkauf          Occasionsobjekt `json:OccasionsObjektVerkauf`
	OccasionsObjektAnkauf           Occasionsobjekt `json:OccasionsObjektAnkauf`
	OccasionsObjektSerienummer      string          `json:OccasionsObjektSerienummer`
	NichtAnzeigen                   bool            `json:NichtAnzeigen`
	ExterneNotizen                  string          `json:ExterneNotizen`
	ExterneNotizenRTF               string          `json:ExterneNotizenRTF`
	InterneNotizen                  string          `json:InterneNotizen`
	InterneNotizenRTF               string          `json:InterneNotizenRTF`
	Optional                        bool            `json:Optional`
	Positionsart                    Positionsart    `json:Positionsart`
	Kundeninstallationen            string          `json:Kundeninstallationen`
	ServiceauftragspositionNr       float64         `json:ServiceauftragspositionNr`
	Positionstyp                    float64         `json:Positionstyp`
	Preisdefinition                 string          `json:Preisdefinition`
	PreisdefinitionOptional         string          `json:PreisdefinitionOptional`
	PreisdefinitionZusatzartikel    string          `json:PreisdefinitionZusatzartikel`
	PreisFW                         float64         `json:PreisFW`
	PreisSW                         float64         `json:PreisSW`
	RabattSW                        float64         `json:RabattSW`
	RabattFW                        float64         `json:RabattFW`
	Rapport                         Rapport         `json:Rapport`
	AutomatischRechnenVon           float64         `json:AutomatischRechnenVon`
	Rekapitulation                  bool            `json:Rekapitulation`
	RueckstandMenge                 float64         `json:RueckstandMenge`
	Steuercode                      Steuercode      `json:Steuercode`
	Leistungsverwaltungstyp         float64         `json:Leistungsverwaltungstyp`
	SteuerFW                        float64         `json:SteuerFW`
	SteuerSW                        float64         `json:SteuerSW`
	Stuecklistenkopf                Artikel         `json:Stuecklistenkopf`
	Subtotalbezeichnungen           string          `json:Subtotalbezeichnungen`
	Dim3                            float64         `json:Dim3`
	TotalFW                         float64         `json:TotalFW`
	TotalInklusivFW                 float64         `json:TotalInklusivFW`
	TotalInklusivSW                 float64         `json:TotalInklusivSW`
	TotalOptional                   float64         `json:TotalOptional`
	TotalSW                         float64         `json:TotalSW`
	Verpackungen                    float64         `json:Verpackungen`
	Vertreter                       Vertreter       `json:Vertreter`
	Zusatzartikelreferenz           float64         `json:Zusatzartikelreferenz`
	Rechnen                         bool            `json:Rechnen`
	GesamtrabattVerbuchungAufteilen bool            `json:GesamtrabattVerbuchungAufteilen`
	BetragFW                        float64         `json:BetragFW`
	BetragDW                        float64         `json:BetragDW`
	BetragSW                        float64         `json:BetragSW`
	Buchung                         Buchung         `json:Buchung`
	Buchungstext                    string          `json:Buchungstext`
	Datum                           string          `json:Datum`
	ZahlungNr                       float64         `json:ZahlungNr`
	Waehrung                        Waehrung        `json:Waehrung`
	Buchungsart                     Buchungsart     `json:Buchungsart`
	Konto                           Konto           `json:Konto`
	Kurs                            float64         `json:Kurs`
	Gutschein                       Gutschein       `json:Gutschein`
	BetragFW                        float64         `json:BetragFW`
	BetragDW                        float64         `json:BetragDW`
	BetragSW                        float64         `json:BetragSW`
	Buchung                         Buchung         `json:Buchung`
	Buchungstext                    string          `json:Buchungstext`
	Datum                           string          `json:Datum`
	Belegnummer                     float64         `json:Belegnummer`
	Waehrung                        Waehrung        `json:Waehrung`
	Buchungsart                     Buchungsart     `json:Buchungsart`
	Konto                           Konto           `json:Konto`
	Kurs                            float64         `json:Kurs`
	BuchungsartVorauszahlung        Buchungsart     `json:BuchungsartVorauszahlung`
	KontoVorauszahlung              Konto           `json:KontoVorauszahlung`
	SteuercodeVorauszahlung         Steuercode      `json:SteuercodeVorauszahlung`
	Parameter                       string          `json:Parameter`
}
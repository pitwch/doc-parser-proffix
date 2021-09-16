package model

type DocsAll struct {
	Entities []DocLink
}
type DocLink struct {
	Name string
	Link string
	Doc  Doc
}
type Doc struct {
	TableName  string
	PrimaryKey string
	Fields     []DocFields
	Methods    []DocMethods
	Parameter  []DocParameter
}
type DocFields struct {
	Feld        string
	Datentyp    string
	NamePROFFIX string
	Besonderes  string
	seitVersion string
}

type DocMethods struct {
	Endpunkt         string
	Rueckgabewert    string
	Beschreibung     string
	BenoetigteLizenz string
	seitVersion      string
}
type DocParameter struct {
	Parameter              string
	Format                 string
	Beschreibung           string
	VerhaltenOhneParameter string
	seitVersion            string
}

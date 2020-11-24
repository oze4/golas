package file

// LASFile represents a .las file
type LASFile struct {
	VersionInformation   Section
	WellInformation      Section
	CurveInformation     Section
	ParameterInformation Section
	Other                Section
	ASCIILogData         LogData
	CustomSections       []Section
}

// Line represents a header line in a .las file section
type Line struct {
	Mnem        string
	Units       string
	Data        string
	Description string
}

// LogData represents a row in the ASCII Log Data section ('~A')
type LogData []interface{}

// Section represents a .las file section
type Section struct {
	Name     string
	Data     []Line
	Comments string
}

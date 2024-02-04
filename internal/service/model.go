package service

const (
	rootDirName   = "Папка"
	parentDirName = "Подпапка"
	fileName      = "Название файла"

	Organization   = "cOrganization"
	ModelTxtName   = "cModelTxtName"
	Programm       = "cProgramm"
	SessionTime    = "tSessionTime"
	SessionDate    = "dSessionDate"
	SessionDateUTC = "dSessionDateUTC"
	SessionTimeUTC = "tSessionTimeUTC"
	DataFileName   = "cDataFileName"
	ProcLevel      = "cProcLevel"

	LUpLat    = "nLUpLat"
	LUpLon    = "nLUpLon"
	RUpLat    = "nRUpLat"
	RUpLon    = "nRUpLon"
	RDownLat  = "nRDownLat"
	RDownLon  = "nRDownLon"
	LDownLat  = "nLDownLat"
	LDownLon  = "nLDownLon"
	LUpNord   = "nLUpNord"
	LUpEast   = "nLUpEast"
	RUpNord   = "nRUpNord"
	RUpEast   = "nRUpEast"
	RDownNord = "nRDownNord"
	RDownEast = "nRDownEast"
	LDownNord = "nLDownNord"
	LDownEast = "nLDownEast"
)

var Titles = []string{
	rootDirName,
	parentDirName,
	fileName,

	Organization,
	ModelTxtName,
	Programm,
	SessionTime,
	SessionDate,
	SessionDateUTC,
	SessionTimeUTC,
	DataFileName,
	ProcLevel,

	LUpLat,
	LUpLon,
	RUpLat,
	RUpLon,
	RDownLat,
	RDownLon,
	LDownLat,
	LDownLon,
	LUpNord,
	LUpEast,
	RUpNord,
	RUpEast,
	RDownNord,
	RDownEast,
	LDownNord,
	LDownEast,
}

type XMLData struct {
	data []byte

	rootDirName   string
	parentDirName string
	fileName      string
}

type XLSXData struct {
	rootDirName   string
	parentDirName string
	fileName      string

	organization   string
	modelTxtName   string
	programm       string
	sessionTime    string
	sessionDate    string
	sessionDateUTC string
	sessionTimeUTC string
	dataFileName   string
	procLevel      string

	lUpLat    string
	lUpLon    string
	rUpLat    string
	rUpLon    string
	rDownLat  string
	rDownLon  string
	lDownLat  string
	lDownLon  string
	lUpNord   string
	lUpEast   string
	rUpNord   string
	rUpEast   string
	rDownNord string
	rDownEast string
	lDownNord string
	lDownEast string
}

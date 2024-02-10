package service

type XMLData struct {
	Data []byte

	RootDirName   string
	ParentDirName string
	FileName      string
}

type XLSXData struct {
	Data map[string]string

	RootDirName   string
	ParentDirName string
	FileName      string
}

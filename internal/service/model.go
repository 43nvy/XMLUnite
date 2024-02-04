package service

type XMLData struct {
	data []byte

	rootDirName   string
	parentDirName string
	fileName      string
}

type XLSXData struct {
	data map[string]string

	rootDirName   string
	parentDirName string
}

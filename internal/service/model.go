package service

import "fmt"

type XMLData struct {
	Data []byte

	RootDirName   string
	ParentDirName string
	FileName      string
}

type XLSXData struct {
	fieldsData    map[string]string
	tagFieldsData map[string]map[string]string

	RootDirName   string
	ParentDirName string
}

type fieldTag string

func (t *fieldTag) openTag() string {
	return fmt.Sprintf("<%s>", *t)
}

func (t *fieldTag) closeTag() string {
	return fmt.Sprintf("</%s>", *t)
}

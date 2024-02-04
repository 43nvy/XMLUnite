package service

import (
	consoleUI "github.com/43nvy/XMLUnite/internal/ui"
)

type ui interface {
	InputData(data string)
	OutputData(data string)
}

type Service interface {
	FindFiles() ([]string, error)
	ReadXMLFile(filePath string) (*XMLData, error)
	ParseXMLFile(xmlData *XMLData) *XLSXData
	ExtractToXLSX(xmlData []*XLSXData) error
}

type service struct {
	ui consoleUI.ConsoleUI

	rootDataDir    string
	outputFileName string

	fields []string
}

func New(ui consoleUI.ConsoleUI, rootDataDir string, outputFileName string, fields []string) Service {
	return &service{
		ui:             ui,
		rootDataDir:    rootDataDir,
		outputFileName: outputFileName,
		fields:         fields,
	}
}

func (s *service) fieldsWithSpace() []string {
	newFields := make([]string, len(s.fields))

	for i, field := range s.fields {
		newFields[i] = field + " "
	}

	return newFields
}

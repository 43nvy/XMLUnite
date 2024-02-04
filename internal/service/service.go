package service

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	consoleUI "github.com/43nvy/XMLUnite/internal/ui"
	"github.com/tealeg/xlsx"
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

	rootDir        string
	outputFileName string
}

func New(ui consoleUI.ConsoleUI, rootDir string, outputFileName string) Service {
	return &service{
		ui:             ui,
		rootDir:        rootDir,
		outputFileName: outputFileName,
	}
}

func (s *service) FindFiles() ([]string, error) {
	var xmlFilesList []string

	err := filepath.Walk(s.rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".xml") {
			xmlFilesList = append(xmlFilesList, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("FindFiles error: %w", err)
	}

	return xmlFilesList, nil
}

func (s *service) ReadXMLFile(filePath string) (*XMLData, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ReadXMLFile error: %w", err)
	}

	fileName := filepath.Base(filePath)
	parentDirName := filepath.Dir(filePath)
	rootDirName := filepath.Dir(parentDirName)

	result := &XMLData{
		data: fileData,

		rootDirName:   filepath.Base(rootDirName),
		parentDirName: filepath.Base(parentDirName),
		fileName:      fileName,
	}

	return result, nil
}

func (s *service) ParseXMLFile(xmlData *XMLData) *XLSXData {
	lines := bytes.Split(xmlData.data, []byte("\n"))

	dataMap := make(map[string]string)

	for _, line := range lines {
		for _, title := range Titles {
			if bytes.Contains(line, []byte(title)) {
				equalIndex := bytes.Index(line, []byte("="))
				if equalIndex == -1 {
					continue
				}

				value := strings.TrimSpace(string(line[equalIndex+1:]))
				dataMap[title] = value
				break
			}
		}
	}

	return &XLSXData{
		rootDirName:   xmlData.rootDirName,
		parentDirName: xmlData.parentDirName,
		fileName:      xmlData.fileName,

		organization:   dataMap[Organization],
		modelTxtName:   dataMap[ModelTxtName],
		programm:       dataMap[Programm],
		sessionTime:    dataMap[SessionDate],
		sessionDateUTC: dataMap[SessionDateUTC],
		sessionTimeUTC: dataMap[SessionTimeUTC],
		dataFileName:   dataMap[DataFileName],
		procLevel:      dataMap[ProcLevel],
		lUpLat:         dataMap[LUpLat],
		lUpLon:         dataMap[LUpLon],
		rUpLat:         dataMap[RUpLat],
		rUpLon:         dataMap[RUpLon],
		rDownLat:       dataMap[RDownLat],
		rDownLon:       dataMap[RDownLon],
		lDownLat:       dataMap[LDownLat],
		lDownLon:       dataMap[LDownLon],
		lUpNord:        dataMap[LUpNord],
		lUpEast:        dataMap[LUpEast],
		rUpNord:        dataMap[RUpNord],
		rUpEast:        dataMap[RUpEast],
		rDownNord:      dataMap[RDownNord],
		rDownEast:      dataMap[RDownEast],
		lDownNord:      dataMap[LDownNord],
		lDownEast:      dataMap[LDownEast],
	}
}

func (s *service) ExtractToXLSX(xmlData []*XLSXData) error {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Лист1")
	if err != nil {
		return fmt.Errorf("ExtractToXLSX add sheet error: %w", err)
	}

	s.fillHeaderRow(sheet)

	for _, data := range xmlData {
		if data != nil {
			s.fillXLSXSheet(sheet, data)
		}
	}

	err = file.Save(s.outputFileName)
	if err != nil {
		return fmt.Errorf("ExtractToXLSX save xlsx file error: %w", err)
	}

	return nil
}

func (s *service) fillHeaderRow(sheet *xlsx.Sheet) {
	headerRow := sheet.AddRow()

	for _, value := range Titles {
		headerRow.AddCell().SetValue(value)
	}
}

func (s *service) fillXLSXSheet(sheet *xlsx.Sheet, xlsxData *XLSXData) {
	dataRow := sheet.AddRow()

	dataRow.AddCell().SetValue(xlsxData.rootDirName)
	dataRow.AddCell().SetValue(xlsxData.parentDirName)
	dataRow.AddCell().SetValue(xlsxData.fileName)

	dataRow.AddCell().SetValue(xlsxData.organization)
	dataRow.AddCell().SetValue(xlsxData.modelTxtName)
	dataRow.AddCell().SetValue(xlsxData.programm)
	dataRow.AddCell().SetValue(xlsxData.sessionTime)
	dataRow.AddCell().SetValue(xlsxData.sessionDate)
	dataRow.AddCell().SetValue(xlsxData.sessionDateUTC)
	dataRow.AddCell().SetValue(xlsxData.sessionTimeUTC)
	dataRow.AddCell().SetValue(xlsxData.dataFileName)
	dataRow.AddCell().SetValue(xlsxData.procLevel)

	dataRow.AddCell().SetValue(xlsxData.lUpLat)
	dataRow.AddCell().SetValue(xlsxData.lUpLon)
	dataRow.AddCell().SetValue(xlsxData.rUpLat)
	dataRow.AddCell().SetValue(xlsxData.rUpLon)
	dataRow.AddCell().SetValue(xlsxData.rDownLat)
	dataRow.AddCell().SetValue(xlsxData.rDownLon)
	dataRow.AddCell().SetValue(xlsxData.lDownLat)
	dataRow.AddCell().SetValue(xlsxData.lDownLon)
	dataRow.AddCell().SetValue(xlsxData.lUpNord)
	dataRow.AddCell().SetValue(xlsxData.lUpEast)
	dataRow.AddCell().SetValue(xlsxData.rUpNord)
	dataRow.AddCell().SetValue(xlsxData.rUpEast)
	dataRow.AddCell().SetValue(xlsxData.rDownNord)
	dataRow.AddCell().SetValue(xlsxData.rDownEast)
	dataRow.AddCell().SetValue(xlsxData.lDownNord)
	dataRow.AddCell().SetValue(xlsxData.lDownEast)
}

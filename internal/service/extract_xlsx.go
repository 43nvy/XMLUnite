package service

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

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

	err = file.Save(s.outputFileName + ".xlsx")
	if err != nil {
		return fmt.Errorf("ExtractToXLSX save xlsx file error: %w", err)
	}

	return nil
}

func (s *service) fillHeaderRow(sheet *xlsx.Sheet) {
	headerRow := sheet.AddRow()

	headerRow.AddCell().SetValue("Папка")
	headerRow.AddCell().SetValue("Подпапка")

	for _, title := range s.fields {
		if title == "cDataFileName" {
			title = "Название файла"
		}

		headerRow.AddCell().SetValue(title)
	}
}

func (s *service) fillXLSXSheet(sheet *xlsx.Sheet, xlsxData *XLSXData) {
	dataRow := sheet.AddRow()

	dataRow.AddCell().SetValue(xlsxData.rootDirName)
	dataRow.AddCell().SetValue(xlsxData.parentDirName)

	for _, val := range s.fields {
		dataRow.AddCell().SetValue(xlsxData.data[val])
	}
}

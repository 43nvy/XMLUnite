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
	headerRow.AddCell().SetValue("Файл")

	for _, title := range s.tagNames {
		headerRow.AddCell().SetValue(title)
	}
}

func (s *service) fillXLSXSheet(sheet *xlsx.Sheet, xlsxData *XLSXData) {
	dataRow := sheet.AddRow()

	dataRow.AddCell().SetValue(xlsxData.RootDirName)
	dataRow.AddCell().SetValue(xlsxData.ParentDirName)
	dataRow.AddCell().SetValue(xlsxData.FileName)

	for _, field := range s.tags {
		dataRow.AddCell().SetValue(xlsxData.Data[field])
	}
}

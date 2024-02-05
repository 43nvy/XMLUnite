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

	for _, title := range s.fieldNames {
		headerRow.AddCell().SetValue(title)
	}

	if len(s.tagFieldNames) != 0 {
		for tagName, tagFieldNameSlice := range s.tagFieldNames {
			for _, fieldName := range tagFieldNameSlice {
				headerTitle := fmt.Sprintf("[%s] %s", tagName, fieldName)
				headerRow.AddCell().SetValue(headerTitle)
			}
		}
	}
}

func (s *service) fillXLSXSheet(sheet *xlsx.Sheet, xlsxData *XLSXData) {
	dataRow := sheet.AddRow()

	dataRow.AddCell().SetValue(xlsxData.rootDirName)
	dataRow.AddCell().SetValue(xlsxData.parentDirName)
	dataRow.AddCell().SetValue(xlsxData.fileName)

	for _, field := range s.fieldNames {
		dataRow.AddCell().SetValue(xlsxData.fieldsData[field])
	}

	if len(s.tagFieldNames) != 0 {
		for tagName, tagFieldNameSlice := range s.tagFieldNames {
			for _, fieldName := range tagFieldNameSlice {
				dataRow.AddCell().SetValue(xlsxData.tagFieldsData[tagName][fieldName])
			}
		}
	}
}

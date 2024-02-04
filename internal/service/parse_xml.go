package service

import (
	"bytes"
	"strings"
)

func (s *service) ParseXMLFile(xmlData *XMLData) *XLSXData {
	lines := bytes.Split(xmlData.data, []byte("\n"))

	dataMap := make(map[string]string)

	fields := s.fieldsWithSpace() // Так как в источнике данных есть вхождения между искомыми строками, нужно убедиться, что берется нужное поле

	for _, line := range lines {
		for _, field := range fields {
			if bytes.Contains(line, []byte(field)) {
				equalIndex := bytes.Index(line, []byte("="))
				if equalIndex == -1 {
					continue
				}

				value := strings.TrimSpace(string(line[equalIndex+1:])) // Чтобы пропустить пробел, берем значение после "=" +1
				dataMap[strings.TrimSpace(field)] = value
				break // Уже нашли значение по списку, можно идти на следующую строку
			}
		}
	}

	return &XLSXData{
		rootDirName:   xmlData.rootDirName,
		parentDirName: xmlData.parentDirName,

		data: dataMap,
	}
}

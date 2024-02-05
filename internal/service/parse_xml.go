package service

import (
	"bytes"
	"strings"
)

func (s *service) ParseXMLFile(data *XMLData) *XLSXData {
	lines := bytes.Split(data.Data, []byte("\n"))

	fieldsData := s.parseFields(lines, s.namesWithSpace(s.fieldNames)) // Так как в источнике данных есть вхождения между искомыми строками, нужно убедиться, что берется нужное поле
	tagFieldsData := s.parseTagFields(lines)

	return &XLSXData{
		rootDirName:   data.RootDirName,
		parentDirName: data.ParentDirName,
		fileName:      data.FileName,

		fieldsData:    fieldsData,
		tagFieldsData: tagFieldsData,
	}
}

func (s *service) parseTagFields(lines [][]byte) map[string]map[string]string {
	tagFieldsData := make(map[string]map[string]string)

	if len(s.tagFieldNames) != 0 {

		for _, tag := range s.tagNames {
			var tagLines [][]byte
			isInsideTag := false
			for _, line := range lines {

				if bytes.Contains(line, []byte(tag.openTag())) {
					isInsideTag = true
					continue
				}

				if isInsideTag {
					tagLines = append(tagLines, line)
				}

				if bytes.Contains(line, []byte(tag.closeTag())) {
					isInsideTag = false
				}
			}

			newTagFieldsNames := s.namesWithSpace(s.tagFieldNames[string(tag)])
			tagData := s.parseFields(tagLines, newTagFieldsNames)
			tagFieldsData[string(tag)] = tagData
		}
	}

	return tagFieldsData
}

func (s *service) parseFields(lines [][]byte, fieldNames []string) map[string]string {
	dataMap := make(map[string]string)

	for _, line := range lines {
		for _, field := range fieldNames {
			if bytes.Contains(line, []byte(field)) {
				value := s.takeValueFromLine(line)
				dataMap[strings.TrimSpace(field)] = value
				break // Уже нашли значение по списку, можно идти на следующую строку
			}
		}
	}

	return dataMap
}

func (s *service) takeValueFromLine(line []byte) string {
	equalIndex := bytes.Index(line, []byte("="))
	if equalIndex == -1 {
		return ""
	}

	value := strings.TrimSpace(string(line[equalIndex+1:])) // Чтобы пропустить пробел, берем значение после "=" +1

	return value
}

package service

import (
	"sort"
)

type Service interface {
	ParseXMLFile(xmlData *XMLData) *XLSXData
	ExtractToXLSX(data []*XLSXData) error
}

type service struct {
	outputFileName string

	fieldNames    []string
	tagNames      []fieldTag
	tagFieldNames map[string][]string
}

func New(outputFileName string, fields []string, tagFields map[string][]string) Service {
	var strTags []string

	for key := range tagFields {
		strTags = append(strTags, key)
	}

	sort.Strings(strTags)

	tags := make([]fieldTag, len(strTags))

	for i, strTag := range strTags {
		tags[i] = fieldTag(strTag)
	}

	return &service{
		outputFileName: outputFileName,
		fieldNames:     fields,
		tagNames:       tags,
		tagFieldNames:  tagFields,
	}
}

func (s *service) namesWithSpace(fieldNames []string) []string {
	newFieldNames := make([]string, len(fieldNames))

	for i, name := range fieldNames {
		newFieldNames[i] = name + " "
	}

	return newFieldNames
}

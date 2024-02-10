package service

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

func (s *service) ParseXMLFile(xmlData *XMLData) (*XLSXData, error) {
	doc, err := xmlquery.Parse(strings.NewReader(string(xmlData.Data)))
	if err != nil {
		return nil, fmt.Errorf("ParseXMLFile parse error: %w", err)
	}

	data := make(map[string]string)

	for _, v := range s.tags {
		queryTag := "//" + v
		node := xmlquery.FindOne(doc, queryTag)
		if node != nil {
			data[v] = strings.TrimSpace(node.InnerText())
		}
	}

	return &XLSXData{
		Data: data,

		RootDirName:   xmlData.RootDirName,
		ParentDirName: xmlData.ParentDirName,
		FileName:      xmlData.FileName,
	}, nil
}

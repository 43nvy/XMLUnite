package service

type Service interface {
	ParseXMLFile(xmlData *XMLData) (*XLSXData, error)
	ExtractToXLSX(data []*XLSXData) error
}

type service struct {
	outputFileName string

	tags     []string
	tagNames []string
}

func New(outputFileName string, tags []string, tagNames []string) Service {
	return &service{
		outputFileName: outputFileName,

		tags:     tags,
		tagNames: tagNames,
	}
}

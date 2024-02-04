package service

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/charmap"
)

func (s *service) ReadXMLFile(filePath string) (*XMLData, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ReadXMLFile error: %w", err)
	}

	utf8Data, err := charmap.Windows1251.NewDecoder().Bytes(fileData)
	if err != nil {
		return nil, fmt.Errorf("Error converting encoding: %w", err)
	}

	fileName := filepath.Base(filePath)
	parentDirName := filepath.Dir(filePath)
	rootDirName := filepath.Dir(parentDirName)

	result := &XMLData{
		data: utf8Data,

		rootDirName:   filepath.Base(rootDirName),
		parentDirName: filepath.Base(parentDirName),
		fileName:      fileName,
	}

	return result, nil
}

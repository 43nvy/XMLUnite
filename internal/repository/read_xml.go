package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/43nvy/XMLUnite/internal/service"
	"golang.org/x/text/encoding/charmap"
)

func (r *repos) ReadXMLFile(filePath string) (*service.XMLData, error) {
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

	return &service.XMLData{
		Data: utf8Data,

		RootDirName:   filepath.Base(rootDirName),
		ParentDirName: filepath.Base(parentDirName),
		FileName:      fileName,
	}, nil
}

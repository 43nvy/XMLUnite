package service

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func (s *service) FindFiles() ([]string, error) {
	var xmlFilesList []string

	err := filepath.Walk(s.rootDataDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".xml") {
			xmlFilesList = append(xmlFilesList, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("FindFiles error: %w", err)
	}

	return xmlFilesList, nil
}

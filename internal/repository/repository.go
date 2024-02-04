package repository

import "github.com/43nvy/XMLUnite/internal/service"

type Repos interface {
	FindFiles() ([]string, error)
	ReadXMLFile(filePath string) (*service.XMLData, error)
}

type repos struct {
	rootDataDir string
}

func New(rootDataDir string) *repos {
	return &repos{
		rootDataDir: rootDataDir,
	}
}

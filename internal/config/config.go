package config

import (
	"fmt"
	"os"
	"strings"
)

type FiledsConfig struct {
	Fileds []string
}

func NewConfig(cfgPath string) (*FiledsConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("NewConfig read file error: %w", err)
	}

	fields := strings.Split(string(data), "\n")

	return &FiledsConfig{
		Fileds: fields,
	}, nil
}

package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type FieldsConfig struct {
	Tags     []string
	TagNames []string
}

func NewConfig(cfgPath string) (*FieldsConfig, error) {
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		err := os.WriteFile(cfgPath, nil, 0644)
		if err != nil {
			return nil, fmt.Errorf("NewConfig create file error: %w", err)
		}

		return nil, fmt.Errorf("файл config.txt не был найден в директории с программой, он был создан заново.")
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("NewConfig read file error: %w", err)
	}

	tags, tagNames := formateFields(data)

	return &FieldsConfig{
		Tags:     tags,
		TagNames: tagNames,
	}, nil
}

func formateFields(cfgFileData []byte) ([]string, []string) {
	scanner := bufio.NewScanner(bytes.NewReader(cfgFileData))

	var tags []string
	var tagNames []string

	for scanner.Scan() {
		line := scanner.Text()
		tags = append(tags, line)
		tagNames = append(tagNames, cleanTag(line))
	}

	return tags, tagNames
}

func cleanTag(tag string) string {
	index := strings.Index(tag, ":")
	if index == -1 {
		return tag
	}

	return tag[index+1:]
}

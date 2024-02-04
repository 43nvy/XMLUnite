package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type FieldsConfig struct {
	Fields    []string
	TagFileds map[string][]string
}

func NewConfig(cfgPath string) (*FieldsConfig, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("NewConfig read file error: %w", err)
	}

	fields, tagFields := formateFields(data)

	return &FieldsConfig{
		Fields:    fields,
		TagFileds: tagFields,
	}, nil
}

func formateFields(cfgFileData []byte) ([]string, map[string][]string) {
	scanner := bufio.NewScanner(bytes.NewReader(cfgFileData))

	var fields []string
	tagFields := make(map[string][]string)

	var inTag bool
	var currTag string

	for scanner.Scan() {
		line := scanner.Text()

		if inTag {
			if line == fmt.Sprintf("</%s>", currTag) {
				inTag = false
				currTag = ""
			} else {
				tagFields[currTag] = append(tagFields[currTag], line)
			}
		} else {
			if len(line) > 2 && line[0] == '<' && line[len(line)-1] == '>' {
				inTag = true
				currTag = line[1 : len(line)-1]
				tagFields[currTag] = []string{}
			} else {
				fields = append(fields, line)
			}
		}
	}

	return fields, tagFields
}

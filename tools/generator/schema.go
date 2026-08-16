package main

import (
	"encoding/json"
	"os"
)

type propertyType struct {
	Type       string        `json:"type"`
	Format     string        `json:"format"`
	Nullable   bool          `json:"nullable"`
	ObjectName string        `json:"objectName"`
	EnumName   string        `json:"enumName"`
	Class      *propertyType `json:"class"`
}

type property struct {
	Description string       `json:"description"`
	Type        propertyType `json:"type"`
	Update      string       `json:"update"`
	Enterprise  bool         `json:"enterprise"`
}

type fieldGroup struct {
	Properties map[string]property `json:"properties"`
	Defaults   map[string]any      `json:"defaults"`
}

type variant struct {
	Name       string `json:"name"`
	Label      string `json:"label"`
	SchemaName string `json:"schemaName"`
}

type schemaEntry struct {
	Type       string    `json:"type"`
	SchemaName string    `json:"schemaName"`
	Variants   []variant `json:"variants"`
}

type enumValue struct {
	Name string `json:"name"`
}

type objectEntry struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Enterprise  bool   `json:"enterprise"`
}

type schemaFile struct {
	Objects map[string]objectEntry `json:"objects"`
	Fields  map[string]fieldGroup  `json:"fields"`
	Schemas map[string]schemaEntry `json:"schemas"`
	Enums   map[string][]enumValue `json:"enums"`
}

func loadSchema(path string) (schemaFile, error) {
	var file schemaFile

	raw, err := os.ReadFile(path)
	if err != nil {
		return file, err
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		return file, err
	}

	return file, nil
}

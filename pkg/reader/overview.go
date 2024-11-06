package reader

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	DOCTOREAD_YML  = "overview.yml"
	DOCTOREAD_YAML = "overview.yaml"
)

type Section struct {
	Title string `yaml:"title"`
	Text  string `yaml:"text"`
}

type Description struct {
	Name     string    `yaml:"name"`
	Sections []Section `yaml:"sections"`
}

type Document struct {
	Description Description `yaml:"description"`
}

func ReadOverview(rootDir string) ([]Document, error) {
	var documentSpecification []Document

	docPathYML := filepath.Join(rootDir, DOCTOREAD_YML)
	docPathYAML := filepath.Join(rootDir, DOCTOREAD_YAML)

	var docPath string
	if _, err := os.Stat(docPathYML); err == nil {
		docPath = docPathYML
	} else if _, err := os.Stat(docPathYAML); err == nil {
		docPath = docPathYAML
	} else {
		log.Printf("File %s or %s not found", docPathYML, docPathYAML)
		return nil, os.ErrNotExist
	}

	data, err := os.ReadFile(docPath)
	if err != nil {
		return nil, err
	}

	var specification Document
	err = yaml.Unmarshal(data, &specification)
	if err != nil {
		return nil, err
	}

	documentSpecification = append(documentSpecification, specification)

	return documentSpecification, nil
}

func ReadSection(rootDir, sectionTitle string) (*Section, error) {
	documents, err := ReadOverview(rootDir)
	if err != nil {
		return nil, err
	}

	for _, document := range documents {
		for _, section := range document.Description.Sections {
			if section.Title == sectionTitle {
				return &section, nil
			}
		}
	}

	return nil, os.ErrNotExist
}
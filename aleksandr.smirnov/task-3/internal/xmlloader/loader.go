package xmlloader

import (
	"encoding/xml"
	"fmt"
	"os"
)

func LoadFile(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read XML file: %w", err)
	}

	if err := xml.Unmarshal(data, target); err != nil {
		return fmt.Errorf("parse XML data: %w", err)
	}

	return nil
}

package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func DecodeXML(r io.Reader) (*Bank, error) {
	decoder := xml.NewDecoder(r)

	var data Bank
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("xml decode failed: %w", err)
	}

	return &data, nil
}

func LoadFromFile(path string) (*Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	return DecodeXML(file)
}

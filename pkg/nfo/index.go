package nfo

import (
	"encoding/xml"
	"os"
)

func Read(path string, v interface{}) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, v)
}

func Write(path string, v interface{}) error {
	bytes, err := xml.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}

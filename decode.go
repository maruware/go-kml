package kml

import (
	"encoding/xml"
	"io"
	"os"
)

func Decode(r io.Reader) (KML, error) {
	var kml KML
	d := xml.NewDecoder(r)
	err := d.Decode(&kml)
	if err != nil {
		return kml, err
	}

	return kml, nil
}

func DecodeFile(p string) (KML, error) {
	f, err := os.Open(p)
	if err != nil {
		return KML{}, err
	}
	defer f.Close()

	return Decode(f)
}

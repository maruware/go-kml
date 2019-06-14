package kml

import (
	"encoding/xml"
	"io"
)

func Encode(w io.Writer, kml KML) error {
	d := xml.NewEncoder(w)
	err := d.Encode(kml)
	return err
}

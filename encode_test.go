package kml

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestEncode(t *testing.T) {
	dir, _ := os.Getwd()
	kmlFile := filepath.Join(dir, "KML_Samples.kml")
	kml, _ := DecodeFile(kmlFile)

	buf := &bytes.Buffer{}
	err := Encode(buf, kml)
	if err != nil {
		t.Fatalf("failed to encode: %v", err)
	}
	if buf.Len() == 0 {
		t.Errorf("encoded data is empty")
	}
}

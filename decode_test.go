package kml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDecodeFile(t *testing.T) {
	dir, _ := os.Getwd()
	kmlFile := filepath.Join(dir, "KML_Samples.kml")
	kml, err := DecodeFile(kmlFile)

	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}
	if kml.Document.Name != "KML Samples" {
		t.Errorf("name do not match expected: %v", kml.Document.Name)
	}
	if len(kml.Document.Styles) != 12 {
		t.Errorf("styles should be 12 but %v", len(kml.Document.Styles))
	}
	expectHref := "http://maps.google.com/mapfiles/kml/pal4/icon28.png"
	if v := kml.Document.Styles[0].IconStyle.Icon.Href; v != expectHref {
		t.Errorf("icon href: expect %v, but %v", expectHref, v)
	}
	if len(kml.Document.Folders) != 6 {
		t.Errorf("folders should be 6 but %v", len(kml.Document.Folders))
	}
	expectName := "Placemarks"
	if v := kml.Document.Folders[0].Name; v != expectName {
		t.Errorf("folder name: expect %v, but %v", expectName, v)
	}
}

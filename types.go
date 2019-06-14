package kml

import (
	"strconv"
	"strings"
)

type named struct {
	Name string `xml:"name"`
}

type Icon struct {
	Href string `xml:"href"`
}

type HotSpot struct {
	X      int    `xml:"x,attr"`
	XUnits string `xml:"xunits,attr"`
	Y      int    `xml:"y,attr"`
	YUnits string `xml:"yunits,attr"`
}

type IconStyle struct {
	Color   string  `xml:"color"`
	Scale   string  `xml:"scale"`
	Icon    Icon    `xml:"Icon"`
	HotSpot HotSpot `xml:"hotSpot"`
}

type BallonStyle struct {
	Text string `xml:text`
}

type PolyStyle struct {
	Color string `xml:text`
}

type LineStyle struct {
	Color string  `xml:"color"`
	Width float32 `xml:"width"`
}

type Style struct {
	IconStyle   IconStyle   `xml:"IconStyle"`
	BallonStyle BallonStyle `xml:"BalloonStyle"`
	LineStyle   LineStyle   `xml:"LineStyle"`
	PolyStyle   PolyStyle   `xml:"PolyStyle"`
}

type Coordinate struct {
	Latitude  float32
	Longitude float32
	Altitude  float32
}

type CoodinatesOne string
type CoodinatesMulti string

type Point struct {
	Coordinates CoodinatesOne `xml:"coordinates"`
}

func parseCoodinates(text string, one bool) ([]Coordinate, error) {
	coordinates := []Coordinate{}
	lines := strings.Split(text, "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		nums := strings.Split(l, ",")
		lng, err := strconv.ParseFloat(nums[0], 32)
		if err != nil {
			return nil, err
		}
		lat, err := strconv.ParseFloat(nums[1], 32)
		if err != nil {
			return nil, err
		}
		alt, err := strconv.ParseFloat(nums[2], 32)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, Coordinate{
			Longitude: float32(lng),
			Latitude:  float32(lat),
			Altitude:  float32(alt),
		})
		if one {
			return coordinates, nil
		}
	}
	return coordinates, nil
}

func (c CoodinatesOne) Parse() (Coordinate, error) {
	coordinates, err := parseCoodinates(string(c), true)
	if len(coordinates) > 0 {
		return coordinates[0], err
	}
	return Coordinate{}, err
}

func (c CoodinatesMulti) Parse() ([]Coordinate, error) {
	return parseCoodinates(string(c), false)
}

type styleURLContainer struct {
	StyleURL string `xml:"styleUrl"`
}

type Placemark struct {
	named
	styleURLContainer
	Point  Point  `xml:"Point"`
	LookAt LookAt `xml:"LookAt"`
}

type LookAt struct {
	Coordinate
	Heading float32 `xml:"heading"`
	Tilt    float32 `xml:"tilt"`
	Range   float32 `xml:"range"`
}

type Folder struct {
	named
	styleURLContainer
	Description string      `xml:"description"`
	Placemarks  []Placemark `xml:"Placemark"`
	LookAt      LookAt      `xml:"LookAt"`
}

type Document struct {
	named
	Description string   `xml:"description"`
	Styles      []Style  `xml:"Style"`
	Folders     []Folder `xml:"Folder"`
}

type KML struct {
	Document Document `xml:"Document"`
}

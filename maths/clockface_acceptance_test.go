package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/biggsean/learn-go-with-tests2/maths/clockface"
)

// SVG was generated 2024-11-10 13:51:56 by https://xml-to-go.github.io/ in Ukraine.
type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}
type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	var tests = []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tt.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", tt.line, svg.Line)
			}

		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	var tests = []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tt.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", tt.line, svg.Line)
			}

		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	var tests = []struct {
		time time.Time
		line Line
	}{
		{simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tt.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", tt.line, svg.Line)
			}

		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

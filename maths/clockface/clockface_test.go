package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	var tests = []struct {
		name     string
		expected float64
		given    time.Time
	}{
		{"thirty seconds in radians", math.Pi, simpleTime(0, 0, 30)},
		{"zero seconds in radians", 0, simpleTime(0, 0, 0)},
		{"forty-five seconds in radians", (math.Pi / 2) * 3, simpleTime(0, 0, 45)},
		{"seven seconds in radians", (math.Pi / 30) * 7, simpleTime(0, 0, 7)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := secondsInRadians(tt.given)
			if !roughlyEqualFloat64(actual, tt.expected) {
				t.Errorf("(%s): expected %v, actual %v", tt.given.Format("15:04:05"), tt.expected, actual)
			}

		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	var tests = []struct {
		name     string
		expected Point
		given    time.Time
	}{
		{"point at thirty seconds", Point{0, -1}, simpleTime(0, 0, 30)},
		{"point at forty-five seconds", Point{-1, 0}, simpleTime(0, 0, 45)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := secondHandPoint(tt.given)
			if !roughlyEqualPoint(actual, tt.expected) {
				t.Errorf("(%s): expected %v, actual %v", tt.given.Format("15:04:05"), tt.expected, actual)
			}

		})
	}

}

func TestMinutesInRadians(t *testing.T) {
	var tests = []struct {
		name     string
		expected float64
		given    time.Time
	}{
		{"thirty minutes in radians", math.Pi, simpleTime(0, 30, 0)},
		{"seven seconds in radians", (math.Pi / (30 * 60)) * 7, simpleTime(0, 0, 7)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := minutesInRadians(tt.given)
			if !roughlyEqualFloat64(actual, tt.expected) {
				t.Errorf("(%s): expected %v, actual %v", tt.given.Format("15:04:05"), tt.expected, actual)
			}

		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	var tests = []struct {
		name     string
		expected Point
		given    time.Time
	}{
		{"point at thirty minutes", Point{0, -1}, simpleTime(0, 30, 0)},
		{"point at forty-five minutes", Point{-1, 0}, simpleTime(0, 45, 0)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := minuteHandPoint(tt.given)
			if !roughlyEqualPoint(actual, tt.expected) {
				t.Errorf("(%s): expected %v, actual %v", tt.given.Format("15:04:05"), tt.expected, actual)
			}

		})
	}
}

func TestHoursInRadians(t *testing.T) {
	var tests = []struct {
		name     string
		expected float64
		given    time.Time
	}{
		{"six in hours", math.Pi, simpleTime(6, 0, 0)},
		{"midnight in hours", 0, simpleTime(0, 0, 0)},
		{"nine PM in hours", math.Pi * 1.5, simpleTime(21, 0, 0)},
		{"one minute and thrty seconds in hours", math.Pi / ((6 * 60 * 60) / 90), simpleTime(0, 1, 30)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := hoursInRadians(tt.given)
			if !roughlyEqualFloat64(actual, tt.expected) {
				t.Errorf("(%s): expected %v, actual %v", tt.given.Format("15:04:05"), tt.expected, actual)
			}

		})
	}
}

func TestHourHandPoint(t *testing.T) {
	var tests = []struct {
		name     string
		expected Point
		given    time.Time
	}{
		{"point at six am", Point{0, -1}, simpleTime(6, 0, 0)},
		{"point at nine pm", Point{-1, 0}, simpleTime(21, 0, 0)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := hourHandPoint(tt.given)
			if !roughlyEqualPoint(actual, tt.expected) {
				t.Errorf("(%s): expected %v, actual %v", tt.given.Format("15:04:05"), tt.expected, actual)
			}

		})
	}
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(1337, time.October, 28, h, m, s, 0, time.UTC)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

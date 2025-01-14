package app

import (
	"fmt"
	"math"
	"strings"
)

type UnitConverter struct {
	from string
	to   string
}

type converterFunc func(float64) float64

var unitConversions = map[string]map[UnitConverter]converterFunc{
	"temperature": {
		{"kelvin", "rankine"}:    func(val float64) float64 { return val * 1.8 },
		{"kelvin", "celsius"}:    func(val float64) float64 { return val - 273.15 },
		{"kelvin", "fahrenheit"}: func(val float64) float64 { return (val-273.15)*9/5 + 32 },

		{"rankine", "kelvin"}:     func(val float64) float64 { return val / 1.8 },
		{"rankine", "celsius"}:    func(val float64) float64 { return (val - 491.67) * 5 / 9 },
		{"rankine", "fahrenheit"}: func(val float64) float64 { return val - 459.67 },

		{"celsius", "kelvin"}:     func(val float64) float64 { return val + 273.15 },
		{"celsius", "rankine"}:    func(val float64) float64 { return (val + 273.15) * 1.8 },
		{"celsius", "fahrenheit"}: func(val float64) float64 { return (val * 9 / 5) + 32 },

		{"fahrenheit", "kelvin"}:  func(val float64) float64 { return (val-32)*5/9 + 273.15 },
		{"fahrenheit", "rankine"}: func(val float64) float64 { return val + 459.67 },
		{"fahrenheit", "celsius"}: func(val float64) float64 { return (val - 32) * 5 / 9 },
	},
	"volume": {
		{"liters", "tablespoons"}:  func(val float64) float64 { return val * 67.628045404 },
		{"liters", "cubic inches"}: func(val float64) float64 { return val * 61.023744095 },
		{"liters", "cups"}:         func(val float64) float64 { return val * 4.2267528377 },
		{"liters", "cubic feet"}:   func(val float64) float64 { return val / 28.316846592 },
		{"liters", "gallons"}:      func(val float64) float64 { return val / 3.785411784 },

		{"tablespoons", "liters"}:       func(val float64) float64 { return val / 67.628045404 },
		{"tablespoons", "cubic inches"}: func(val float64) float64 { return val / 1.108225108 },
		{"tablespoons", "cups"}:         func(val float64) float64 { return val / 16 },
		{"tablespoons", "cubic feet"}:   func(val float64) float64 { return val / 1915.0129863 },
		{"tablespoons", "gallons"}:      func(val float64) float64 { return val / 256 },

		{"cubic inches", "liters"}:      func(val float64) float64 { return val / 61.023744095 },
		{"cubic inches", "tablespoons"}: func(val float64) float64 { return val * 1.108225108 },
		{"cubic inches", "cups"}:        func(val float64) float64 { return val / 14.4375 },
		{"cubic inches", "cubic feet"}:  func(val float64) float64 { return val / 1728 },
		{"cubic inches", "gallons"}:     func(val float64) float64 { return val / 231 },

		{"cups", "liters"}:       func(val float64) float64 { return val / 4.2267528377 },
		{"cups", "tablespoons"}:  func(val float64) float64 { return val * 16 },
		{"cups", "cubic inches"}: func(val float64) float64 { return val * 14.4375 },
		{"cups", "cubic feet"}:   func(val float64) float64 { return val / 119.688311688 },
		{"cups", "gallons"}:      func(val float64) float64 { return val / 16 },

		{"cubic feet", "liters"}:       func(val float64) float64 { return val * 28.316846592 },
		{"cubic feet", "tablespoons"}:  func(val float64) float64 { return val * 1915.0129863 },
		{"cubic feet", "cubic inches"}: func(val float64) float64 { return val * 1728 },
		{"cubic feet", "cups"}:         func(val float64) float64 { return val * 119.688311688 },
		{"cubic feet", "gallons"}:      func(val float64) float64 { return val * 7.48051948 },

		{"gallons", "liters"}:       func(val float64) float64 { return val * 3.785411784 },
		{"gallons", "tablespoons"}:  func(val float64) float64 { return val * 256 },
		{"gallons", "cubic inches"}: func(val float64) float64 { return val * 231 },
		{"gallons", "cups"}:         func(val float64) float64 { return val * 16 },
		{"gallons", "cubic feet"}:   func(val float64) float64 { return val / 7.48051948 },
	},
}

var roundFunc = func(value float64) float64 {
	return math.Floor(value*10+0.5) / 10
}

func ConvertUnits(from, to string, val float64) (float64, error) {
	lowerFrom := strings.ToLower(from)
	lowerTo := strings.ToLower(to)
	if lowerFrom == lowerTo {
		return roundFunc(val), nil
	}

	unitConverter := UnitConverter{
		lowerFrom,
		lowerTo,
	}
	if convertFunc, ok := unitConversions["temperature"][unitConverter]; ok {
		return roundFunc(convertFunc(val)), nil
	}
	if convertFunc, ok := unitConversions["volume"][unitConverter]; ok {
		return roundFunc(convertFunc(val)), nil
	}

	return -1, fmt.Errorf("invalid conversion: from %s, to %s", from, to)
}

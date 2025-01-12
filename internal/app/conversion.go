package app

import (
	"fmt"
	"math"
	"strings"
)

type UnitConverter struct {
	from     string
	to       string
	unitType string
}

type converterFuncs struct {
	ToBaseFunc   func(float64) float64
	FromBaseFunc func(float64) float64
}

var unitConversions = map[string]map[string]converterFuncs{
	"temperature": {
		// treat kelvin as base temperature unit
		"kelvin": {
			func(val float64) float64 { return val },
			func(val float64) float64 { return val },
		},
		"rankine": {
			func(val float64) float64 { return val * 0.555556 },
			func(val float64) float64 { return val * 1.8 },
		},
		"celsius": {
			func(val float64) float64 { return val + 273.15 },
			func(val float64) float64 { return val - 273.15 },
		},
		"fahrenheit": {
			func(val float64) float64 { return (val-32)*5/9 + 273.15 },
			func(val float64) float64 { return (val-273.15)*9/5 + 32 },
		},
	},
	"volume": {
		// treat liters as base volume unit
		"liters": {
			func(val float64) float64 { return val },
			func(val float64) float64 { return val },
		},
		"gallons": {
			func(val float64) float64 { return val * 3.78541 },
			func(val float64) float64 { return val / 3.78541 },
		},
		"cups": {
			func(val float64) float64 { return val / 4.227 },
			func(val float64) float64 { return val * 4.227 },
		},
		"tablespoons": {
			func(val float64) float64 { return val / 67.628 },
			func(val float64) float64 { return val * 67.628 },
		},
		"cubic feet": {
			func(val float64) float64 { return val * 28.317 },
			func(val float64) float64 { return val / 28.317 },
		},
		"cubic inches": {
			func(val float64) float64 { return val / 61.024 },
			func(val float64) float64 { return val * 61.024 },
		},
	},
}

var roundFunc = func(value float64) float64 {
	return math.Round(value*10) / 10
}

func NewConverter(from, to string) (*UnitConverter, error) {
	lowerFrom := strings.ToLower(from)
	lowerTo := strings.ToLower(to)
	_, fromIsTemp := unitConversions["temperature"][lowerFrom]
	_, toIsTemp := unitConversions["temperature"][lowerTo]
	if fromIsTemp && toIsTemp {
		return &UnitConverter{
			from:     lowerFrom,
			to:       lowerTo,
			unitType: "temperature",
		}, nil
	}

	_, fromIsVolume := unitConversions["volume"][lowerFrom]
	_, toIsVolume := unitConversions["temperature"][lowerTo]
	if fromIsVolume && toIsVolume {
		return &UnitConverter{
			from:     lowerFrom,
			to:       lowerTo,
			unitType: "volume",
		}, nil
	}

	return nil, fmt.Errorf("invalid conversion: from %s, to %s", from, to)
}

func (c *UnitConverter) Convert(val float64) float64 {
	baseVal := unitConversions[c.unitType][c.from].ToBaseFunc(val)
	return roundFunc(unitConversions[c.unitType][c.to].FromBaseFunc(baseVal))
}

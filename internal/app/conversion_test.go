package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllConversions(t *testing.T) {
	type UandV struct {
		Unit     string
		Val      float64
		Expected float64
	}

	conversionsToTest := map[string][]UandV{
		// temperatures
		"kelvin": {
			{"kelvin", 100, 100},
			{"rankine", 100, 180.0},
			{"celsius", 100, -173.15},
			{"fahrenheit", 100, -279.67},
		},
		"rankine": {
			{"kelvin", 100, 55.55},
			{"rankine", 100, 100},
			{"celsius", 100, -217.594},
			{"fahrenheit", 100, -359.67},
		},
		"celsius": {
			{"kelvin", 100, 373.15},
			{"rankine", 100, 671.67},
			{"celsius", 100, 100},
			{"fahrenheit", 100, 212},
		},
		"fahrenheit": {
			{"kelvin", 100, 310.928},
			{"rankine", 100, 559.67},
			{"celsius", 100, 37.7778},
			{"fahrenheit", 100, 100},
		},

		// volumes
		"liters": {
			{"liters", 100, 100},
			{"gallons", 100, 26.4172},
			{"cups", 100, 422.675},
			{"tablespoons", 100, 6762.8},
			{"cubic feet", 100, 3.5315},
			{"cubic inches", 100, 6102.37},
		},
		"gallons": {
			{"liters", 100, 378.541},
			{"gallons", 100, 100},
			{"cups", 100, 1600},
			{"tablespoons", 100, 25600},
			{"cubic feet", 100, 13.3681},
			{"cubic inches", 100, 23100},
		},
		"cups": {
			{"liters", 100, 23.6588},
			{"gallons", 100, 6.25},
			{"cups", 100, 100},
			{"tablespoons", 100, 1600},
			{"cubic feet", 100, 0.8355},
			{"cubic inches", 100, 1443.75},
		},
		"tablespoons": {
			{"liters", 100, 1.4787},
			{"gallons", 100, 0.3906},
			{"cups", 100, 6.25},
			{"tablespoons", 100, 100},
			{"cubic feet", 100, 0.0522},
			{"cubic inches", 100, 90.2344},
		},
		"cubic feet": {
			{"liters", 100, 2831.68},
			{"gallons", 100, 748.052},
			{"cups", 100, 11968.8},
			{"tablespoons", 100, 191501.3},
			{"cubic feet", 100, 100},
			{"cubic inches", 100, 172800},
		},
		"cubic inches": {
			{"liters", 100, 1.63871},
			{"gallons", 100, 0.4329},
			{"cups", 100, 6.92641},
			{"tablespoons", 100, 110.823},
			{"cubic feet", 100, 0.0578704},
			{"cubic inches", 100, 100},
		},
	}
	for fromUnit, toUandV := range conversionsToTest {
		for _, to := range toUandV {
			result, err := ConvertUnits(fromUnit, to.Unit, to.Val)
			assert.Nil(t, err)
			assert.Equal(t, roundFunc(to.Expected), result, []string{fromUnit, to.Unit})
		}
	}
}

func TestInvalidConversion(t *testing.T) {
	_, err := ConvertUnits("not a unit", "also not a unit", 123.123)
	assert.Error(t, err)

	_, err = ConvertUnits("celsius", "cubic inches", 234.567)
	assert.Error(t, err)

	_, err = ConvertUnits("liters", "rankine", 234.567)
	assert.Error(t, err)

	_, err = ConvertUnits("liters", "not even close to a unit 0.0", 123.123)
	assert.Error(t, err)
}

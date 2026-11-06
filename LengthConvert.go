package main

import (
	"fmt"
	"strconv"
)

var lengthToMeters = map[string]float64{
	"millimeter":  0.001,
	"millimeters": 0.001,
	"centimeter":  0.01,
	"centimeters": 0.01,
	"meter":       1.0,
	"meters":      1.0,
	"kilometer":   1000.0,
	"kilometers":  1000.0,
	"inch":        0.0254,
	"inches":      0.0254,
	"foot":        0.3048,
	"feet":        0.3048,
	"yard":        0.9144,
	"yards":       0.9144,
	"mile":        1609.344,
	"miles":       1609.344,
}

func convertLength(value float64, fromUnit string, toUnit string) (float64, error) {
	fromFactor, checkOne := lengthToMeters[fromUnit]
	toFactor, checkTwo := lengthToMeters[toUnit]

	if !checkOne || !checkTwo {
		return 0, fmt.Errorf("unsupported conversion from %s to %s", fromUnit, toUnit)
	}

	meters := value * fromFactor

	result := fmt.Sprintf("%.2f", meters/toFactor)
	floatResult, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting result to float: %v", err)
	}

	return floatResult, nil
}

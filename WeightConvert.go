package main

import (
	"fmt"
	"strconv"
)

var weightToGrams = map[string]float64{
	"milligram": 0.001,
	"gram":      1.0,
	"kilogram":  1000.0,
	"ounce":     28.349523125,
	"pound":     453.59237,
}

func convertWeight(value float64, fromUnit string, toUnit string) (float64, error) {
	fromFactor, ok1 := weightToGrams[fromUnit]
	toFactor, ok2 := weightToGrams[toUnit]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("unsupported conversion from %s to %s", fromUnit, toUnit)
	}

	// Convert to base unit (grams), then to target unit
	grams := value * fromFactor
	returns := fmt.Sprintf("%.2f", grams/toFactor)
	floatResult, err := strconv.ParseFloat(returns, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting result to float: %v", err)
	}

	return floatResult, nil
}

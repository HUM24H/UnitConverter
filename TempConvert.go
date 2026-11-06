package main

import "fmt"

func convertTemperature(value float64, fromUnit string, toUnit string) (float64, error) {
	var celsius float64
	switch fromUnit {
	case "Celsius", "celsius":
		celsius = value
	case "Fahrenheit", "fahrenheit":
		celsius = (value - 32) * 5.0 / 9.0
	case "Kelvin", "kelvin":
		celsius = value - 273.15
	default:
		return 0, fmt.Errorf("unsupported source temperature unit: %s", fromUnit)
	}

	switch toUnit {
	case "Celsius", "celsius":
		return celsius, nil
	case "Fahrenheit", "fahrenheit":
		return (celsius * 9.0 / 5.0) + 32, nil
	case "Kelvin", "kelvin":
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("unsupported target temperature unit: %s", toUnit)
	}
}

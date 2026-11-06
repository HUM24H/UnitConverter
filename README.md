# Unit Converter Web App - Built in Go
## Introduction
Unit Converter Web App built using Go

**Note: This CLI app was built for the purpose of learning Go. This is not meant for any production or public use**

## Description
The Unit Converter is a simple Web App that allows a user to convert from 3 types of units:
- Length
- Weight
- Temperature

## How to Use
In order to run this Web App, you will first need the "web-app" executable and the "/scripts" webpage.

This needs to be run within a Linux environment, simply execute the "web-app" executable and you will be presented with the following:

`2026/11/06 14:51:20 Starting server on :8080`

Then in a Web Browser within the same machine the Go Webserver is running, navigate to: http://localhost:8080

From here you will get a Webpage, like below:

<img width="1685" height="1093" alt="image" src="https://github.com/user-attachments/assets/89f50e1b-d37d-4ec9-9454-366f5bd76689" />

Across the Web Page are three types of units; Length, Weight & Temperature.

Below is an easy step-by-step of what to do:
- Enter a number into one of the "VALUE" fields
- Select desired "FROM UNIT"
- Select desired "TO UNIT"
- Click "Convert" button

From here the results will be displayed their respective "Result" box, each one is rounded to 2 decimal places.

Finally, below is a list of all supported units:
* **Length**
  * Millimeter
  * Centimeter
  * Meter
  * Kilometer
  * Inch
  * Foot
  * Yard
  * Mile
* **Weight**
  * Milligram
  * Gram
  * Kilogram
  * Ounce
  * Pound
* **Temperature**
  * Celsius
  * Fahrenheit
  * Kelvin

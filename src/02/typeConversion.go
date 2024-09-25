package main

import "strconv"

func main() {
	// Tip dönüşümleri
	allConversion()

	println("--------------------------------------------------")

	// String to integer
	stringToInteger()

	println("--------------------------------------------------")

	// String to float
	stringToFloat()

	println("--------------------------------------------------")

	// Integer to string
	integerToString()

	println("--------------------------------------------------")

	// Float to string
	floatToString()

	println("--------------------------------------------------")

	// Boolean to string
	booleanToString()

	println("--------------------------------------------------")

	// String to boolean
	stringToBoolean()

	println("--------------------------------------------------")

	// String to byte
	stringToByte()

	println("--------------------------------------------------")

	// Byte to string
	byteToString()
}

func allConversion() {
	// Tip dönüşümleri
	var a int = 10
	var b float64 = float64(a)
	var c int = int(b)
	println(a, b, c)
}

func stringToInteger() {
	// String to integer
	var str string = "10"
	a, _ := strconv.Atoi(str)
	println(a)
}

func stringToFloat() {
	// String to float
	var str string = "10.5"
	a, _ := strconv.ParseFloat(str, 64)
	println(a)
}

func integerToString() {
	// Integer to string
	var a int = 10
	str := strconv.Itoa(a)
	println(str)
}

func floatToString() {
	// Float to string
	var a float64 = 10.5
	str := strconv.FormatFloat(a, 'f', 6, 64)
	println(str)
}

func booleanToString() {
	// Boolean to string
	var a bool = true
	str := strconv.FormatBool(a)
	println(str)
}

func stringToBoolean() {
	// String to boolean
	var str string = "true"
	a, _ := strconv.ParseBool(str)
	println(a)
}

func stringToByte() {
	// String to byte
	var str string = "A"
	a := []byte(str)
	println(a)
}

func byteToString() {
	// Byte to string
	var a []byte = []byte{'A'}
	str := string(a)
	println(str)
}

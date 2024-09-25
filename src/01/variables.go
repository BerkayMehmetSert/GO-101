package main

func main() {
	// String değişken tanımlama
	stringVariables() // stringVariables fonksiyonunu çağırma

	println("----------------")

	// Integer değişken tanımlama
	intVariable() // intVariable fonksiyonunu çağırma

	println("----------------")

	// Float değişken tanımlama
	floatVariable() // floatVariable fonksiyonunu çağırma

	println("----------------")

	// Boolean değişken tanımlama
	boolVariable() // boolVariable fonksiyonunu çağırma

	println("----------------")

	// Sabit değişken tanımlama
	constantVariables() // constantVariables fonksiyonunu çağırma
}

func stringVariables() {
	// String değişken tanımlama
	var stringVariable = "Hello World!"

	// Alternatif tanımlama
	stringVariable2 := "Hello World!"

	// Alternatif tanımlama
	var stringVariable3 string = "Hello World!"

	// Ekran çıktısı
	println(stringVariable)
	println(stringVariable2)
	println(stringVariable3)
}

func intVariable() {
	// Integer değişken tanımlama
	var intVariable = 10

	// Alternatif tanımlama
	intVariable2 := 10

	// Alternatif tanımlama
	var intVariable3 int = 10

	// Ekran çıktısı
	println(intVariable)
	println(intVariable2)
	println(intVariable3)
}

func floatVariable() {
	// Float değişken tanımlama
	var floatVariable = 10.5

	// Alternatif tanımlama
	floatVariable2 := 10.5

	// Alternatif tanımlama
	var floatVariable3 float64 = 10.5

	// Ekran çıktısı
	println(floatVariable)
	println(floatVariable2)
	println(floatVariable3)
}

func boolVariable() {
	// Boolean değişken tanımlama
	var boolVariable = true

	// Alternatif tanımlama
	boolVariable2 := false

	// Alternatif tanımlama
	var boolVariable3 bool = true

	// Ekran çıktısı
	println(boolVariable)
	println(boolVariable2)
	println(boolVariable3)
}

func constantVariables() {
	// Sabit değişken tanımlama
	const constantVariable = "Hello World!"
	const constantIntegerVariable = 3.14

	// Alternatif tanımlama
	const constantVariable2 string = "Hello World!"
	const constantIntegerVariable2 = 3.14

	// Ekran çıktısı
	println(constantVariable)
	println(constantVariable2)
	println(constantIntegerVariable)
	println(constantIntegerVariable2)
}

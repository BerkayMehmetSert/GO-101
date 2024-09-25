package main

func main() {
	// Basit bir array tanımlama
	simpleArray()

	println("----------------")

	// Alternatif basit bir array tanımlama
	alternativeSimpleArray()

	println("----------------")

	// Parametreli array tanımlama
	sampleArrayWithParameters("Hello", "World", "!")

	println("----------------")

	// Otomatik boyutlandırma (dizi boyutunu belirtmeye gerek yok)
	automaticSizeArray()

	// Gerçek hayat uygulamalarında array'lerin kullanımı
	newLanguage := Language{
		Name:     "Golang",
		Symbol:   "Go",
		Paradigm: "Concurrent, Imperative, Structured, Object-Oriented",
	}

	newLanguageForCSharp := Language{
		Name:     "CSharp",
		Symbol:   "C#",
		Paradigm: "Structured, Imperative, Object-Oriented, Event-Driven, Functional, Generic, Reflective, Concurrent",
	}

	languages := [...]Language{newLanguage, newLanguageForCSharp}

	for _, language := range languages {
		println(language.Name + " (" + language.Symbol + ") - " + language.Paradigm)
	}
}

func simpleArray() {
	simples := [3]string{}
	simples[0] = "Hello"
	simples[1] = "World"
	simples[2] = "!"

	simplesToString := simples[0] + " " + simples[1] + simples[2]

	println(simplesToString)
}

func alternativeSimpleArray() {
	simples := [3]string{"Hello", "World", "!"}

	simplesToString := simples[0] + " " + simples[1] + simples[2]

	println(simplesToString)
}

func sampleArrayWithParameters(item1, item2, item3 string) {
	samples := [3]string{item1, item2, item3}

	samplesToString := samples[0] + " " + samples[1] + samples[2]

	println(samplesToString)
}

func automaticSizeArray() {
	// Otomatik boyutlandırma
	autoSize := [...]string{"Hello", "World"}

	autoSizeToString := autoSize[0] + " " + autoSize[1]

	println(autoSizeToString)
}

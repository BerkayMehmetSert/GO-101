package main

func main() {
	// Struct tanımlama
	language1 := Language{"Golang", "Go"}

	// Ekran çıktısı
	println("Language1:", language1.Name, language1.Symbol)

	println("----------------")

	// Struct örneği oluşturma
	language2 := new(Language)
	language2.Name = "CSharp"
	language2.Symbol = "C#"

	// Ekran çıktısı
	println("Language2:", language2.Name, language2.Symbol)

	println("----------------")

	// Boş yapıcı fonksiyon
	language3 := NewLanguage()
	language3.Name = "Java"
	language3.Symbol = "J"

	// Ekran çıktısı
	println("Language3:", language3.Name, language3.Symbol)

	println("----------------")

	// Parametreli yapıcı fonksiyon
	language4 := NewLanguageWithParams("Python", "Py")
	println("Language4:", language4.Name, language4.Symbol)

	println("----------------")

	// Parametreli yapıcı fonksiyon alternatif
	language5 := NewLanguageWithParamsShortcut("JavaScript", "JS")
	println("Language5:", language5.Name, language5.Symbol)

	println("----------------")

	// Parametreli yapıcı fonksiyon alternatif
	language6 := NewLanguageWithParamsShortcut2("TypeScript", "TS")
	println("Language6:", language6.Name, language6.Symbol)

}

type Language struct {
	Name   string
	Symbol string
}

// Boş yapıcı fonksiyon
func NewLanguage() *Language {
	return new(Language)
}

// Parametreli yapıcı fonksiyon
func NewLanguageWithParams(name string, symbol string) *Language {
	return &Language{name, symbol}
}

// Parametreli yapıcı fonksiyon alternatif
func NewLanguageWithParamsShortcut(name string, symbol string) *Language {
	return &Language{Name: name, Symbol: symbol}
}

// Parametreli yapıcı fonksiyon alternatif
func NewLanguageWithParamsShortcut2(name, symbol string) *Language {
	return &Language{Name: name, Symbol: symbol}
}

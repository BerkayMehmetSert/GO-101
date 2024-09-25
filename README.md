# Go 101

## Go Nedir?

Go, basit, hızlı ve güvenilir yazılım geliştirmeyi amaçlayan açık kaynak bir programlama dilidir. Hızlı derleme süreci,
hafif söz dizimi ve etkin bir çöp toplama mekanizması gibi özellikleriyle dikkat çeker.

Go'nun kullanım alanı oldukça geniştir ve çeşitli geliştirme araçları, paketleri ve modülleri içerir. Web uygulamaları,
API'ler, veritabanı sistemleri, ağ yazılımı gibi çeşitli uygulama türlerini geliştirmek için tercih edilir.

Dilin temel özelliklerinden biri, hızlı derleme sürecidir. Benzerlikleriyle C diline benzeyen söz dizimi ve C++'taki
nesne yönelimli programlama özelliklerini kullanır. Ayrıca, çöp toplama, hafif iş parçacıkları, kapsüllü tipler ve
dinamik bellek yönetimi gibi özellikler de bulunur.

Go'nun önemli bir özelliği, etkin paket yönetimi sistemidir. Modüller ve paketler, geliştiricilerin kodlarını organize
etmelerini ve yeniden kullanılabilir bileşenler oluşturmalarını sağlar.

Sonuç olarak, Go dilinin hızlı derleme süreci, hafif söz dizimi, etkin paket yönetimi, çöp toplama ve kapsüllü tipler
gibi özellikleri, profesyonel yazılım geliştirme süreçlerine katkıda bulunur. Bu özellikler sayesinde, çeşitli uygulama
türlerini geliştirmek için ideal bir seçenek sunar.

### Go Programlama Diline Giriş: Değişkenler ve Sabitler

Go dilinde değişkenler ve sabitler kullanarak programlarımızı esnek ve güçlü hale getirebiliriz. Bu öğreticide, farklı
veri türlerinde değişkenlerin tanımlanması ve kullanılması üzerine örnekler göstereceğiz.

Öncelikle, bir Go programında değişken tanımlamak için `var` anahtar kelimesi veya `:=` kısaltması kullanılır. Bir
değişkenin türünü belirtmek için var anahtar kelimesi kullanılırken, `:=` kısaltması değişkenin türünü otomatik olarak
belirler.

```go
package main

func main() {
	// String değişkenlerin tanımlanması ve kullanılması
	var stringVariable = "Hello World!"
	stringVariable2 := "Hello World!"
	var stringVariable3 string = "Hello World!"

	println(stringVariable)
	println(stringVariable2)
	println(stringVariable3)
}
```

Yukarıdaki örnekte, üç farklı şekilde bir string değişken
tanımlıyoruz: `var stringVariable`, `stringVariable2 :=`, `var stringVariable3 string`. Ardından, bu
değişkenler `println()` fonksiyonu ile ekrana yazdırılıyor.

Benzer şekilde, integer, float, ve boolean değişkenler de tanımlanabilir ve kullanılabilir:

```go
package main

func main() {
	// Integer değişkenlerin tanımlanması ve kullanılması
	var intVariable = 10
	intVariable2 := 10
	var intVariable3 int = 10

	println(intVariable)
	println(intVariable2)
	println(intVariable3)

	// Float değişkenlerin tanımlanması ve kullanılması
	var floatVariable = 10.5
	floatVariable2 := 10.5
	var floatVariable3 float64 = 10.5

	println(floatVariable)
	println(floatVariable2)
	println(floatVariable3)

	// Boolean değişkenlerin tanımlanması ve kullanılması
	var boolVariable = true
	boolVariable2 := false
	var boolVariable3 bool = true

	println(boolVariable)
	println(boolVariable2)
	println(boolVariable3)
}
```

Son olarak, Go dilinde sabitler de tanımlanabilir. Sabitler, değerleri değişmeyen ve program boyunca sabit kalan
değerlerdir:

```go
package main

func main() {
	// Sabit değişkenlerin tanımlanması ve kullanılması
	const constantVariable = "Hello World!"
	const constantVariable2 string = "Hello World!"

	println(constantVariable)
	println(constantVariable2)
}
```

Yukarıdaki örnek, iki farklı şekilde bir sabit tanımlar: `const constantVariable`, `const constantVariable2 string`.
Ardından, bu sabitler `println()` fonksiyonu ile ekrana yazdırılır.

### Go Programlama Diline Giriş: Tip Dönüşümleri

Go dilinde veri türleri arasında dönüşümler yapmak sıkça karşılaşılan bir ihtiyaçtır. Bu öğreticide, farklı veri türleri
arasında dönüşümler yapmayı göstereceğiz.

Öncelikle, farklı veri türleri arasında genel dönüşümler yapmak için `strconv` paketini kullanırız. Bu paket, stringleri
integer veya float değerlere dönüştürme gibi işlemleri gerçekleştirmemize olanak tanır.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Tip dönüşümleri
	allConversion()

	fmt.Println("--------------------------------------------------")

	// String'ten integer'a dönüşüm
	stringToInteger()

	fmt.Println("--------------------------------------------------")

	// String'ten float'a dönüşüm
	stringToFloat()

	fmt.Println("--------------------------------------------------")

	// Integer'dan string'e dönüşüm
	integerToString()

	fmt.Println("--------------------------------------------------")

	// Float'tan string'e dönüşüm
	floatToString()

	fmt.Println("--------------------------------------------------")

	// Boolean'dan string'e dönüşüm
	booleanToString()

	fmt.Println("--------------------------------------------------")

	// String'ten boolean'a dönüşüm
	stringToBoolean()

	fmt.Println("--------------------------------------------------")

	// String'ten byte dizisine dönüşüm
	stringToByte()

	fmt.Println("--------------------------------------------------")

	// Byte dizisinden string'e dönüşüm
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
```

Yukarıdaki örnekte, farklı dönüşümler için bir dizi fonksiyon tanımlanmıştır. Bunlar arasında string'i integer veya
float'a dönüştürme, integer'ı string'e dönüştürme, float'ı string'e dönüştürme, boolean'ı string'e dönüştürme, string'i
boolean'a dönüştürme, string'i byte dizisine dönüştürme ve byte dizisini string'e dönüştürme gibi işlemler yer alır. Bu
fonksiyonlar, `strconv` paketini kullanarak dönüşümleri gerçekleştirirler.

### Go Programlama Diline Giriş: Fonksiyonlar

Go dilinde fonksiyonlar, programınızın modülerliğini artıran ve kod tekrarını önleyen önemli yapı taşlarıdır. Bu
öğreticide, farklı türde fonksiyonların nasıl tanımlanacağını ve kullanılacağını göstereceğiz.

Öncelikle, basit bir fonksiyon tanımlayalım ve onu çağıralım:

```go
package main

import "fmt"

func main() {
	// Fonksiyon çağırma
	sayHello()
}

func sayHello() {
	fmt.Println("Hello World!")
}
```

Yukarıdaki örnekte, `sayHello` adında basit bir fonksiyon tanımladık ve `main` fonksiyonunda çağırdık. Bu fonksiyon,
ekrana **"Hello World!"** yazdıracaktır.

Şimdi, parametre alan ve parametre döndüren fonksiyonları inceleyelim:

```go
package main

import "fmt"

func main() {
	// Fonksiyon çağırma
	sayHelloTo("Go")
	fmt.Println(sum(10, 20))
}

// Parametreli fonksiyon
func sayHelloTo(name string) {
	fmt.Println("Hello", name)
}

// Değer döndüren fonksiyon
func sum(a int, b int) int {
	return a + b
}
```

Yukarıdaki örnekte, `sayHelloTo` fonksiyonu `name` adında bir parametre alır ve bu ismi kullanarak bir mesaj yazdırır.
sum fonksiyonu ise `a` ve `b` adında iki parametre alır ve bu parametrelerin toplamını döndürür.

Fonksiyonlar, genellikle programın farklı bölümlerini modüler hale getirmek için kullanılır ve kodun daha okunabilir ve
bakımı daha kolay hale getirilmesine yardımcı olur.

### Go Programlama Diline Giriş: Struct ve Yapı Yöntemleri

Go dilinde, struct'lar karmaşık veri yapılarını temsil etmek için kullanılır. Bu öğreticide, bir struct tanımlamak, bu
struct'tan örnekler oluşturmak ve yapı yöntemleri kullanarak bu örneklerle çalışmak üzerinde duracağız.

Öncelikle, basit bir `struct` tanımlayalım ve bu struct'tan örnekler oluşturalım:

```go
package main

import "fmt"

func main() {
	// Struct tanımlama
	language1 := Language{"Golang", "Go"}

	// Ekran çıktısı
	fmt.Println("Language1:", language1.Name, language1.Symbol)

	fmt.Println("----------------")

	// Struct örneği oluşturma
	language2 := new(Language)
	language2.Name = "CSharp"
	language2.Symbol = "C#"

	// Ekran çıktısı
	fmt.Println("Language2:", language2.Name, language2.Symbol)

	fmt.Println("----------------")

	// Boş yapıcı fonksiyon
	language3 := NewLanguage()
	language3.Name = "Java"
	language3.Symbol = "J"

	// Ekran çıktısı
	fmt.Println("Language3:", language3.Name, language3.Symbol)

	fmt.Println("----------------")

	// Parametreli yapıcı fonksiyon
	language4 := NewLanguageWithParams("Python", "Py")
	fmt.Println("Language4:", language4.Name, language4.Symbol)

	fmt.Println("----------------")

	// Parametreli yapıcı fonksiyon alternatif
	language5 := NewLanguageWithParamsShortcut("JavaScript", "JS")
	fmt.Println("Language5:", language5.Name, language5.Symbol)

	fmt.Println("----------------")

	// Parametreli yapıcı fonksiyon alternatif
	language6 := NewLanguageWithParamsShortcut2("TypeScript", "TS")
	fmt.Println("Language6:", language6.Name, language6.Symbol)
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
```

Yukarıdaki örnekte, Language adında bir struct tanımladık. Bu struct, bir programlama dilinin adını ve sembolünü temsil
eder. Daha sonra, bu struct'tan birkaç örnek oluşturduk ve farklı yapı yöntemlerini kullanarak bu örneklerle çalıştık.

Bu örnekte kullanılan yapı yöntemleri şunlardır:

1. **Boş yapıcı fonksiyon:** `NewLanguage` fonksiyonu, boş bir `Language` örneği oluşturur.

2. **Parametreli yapıcı fonksiyon:** `NewLanguageWithParams` fonksiyonu, verilen ad ve sembol değerleriyle
   bir `Language` örneği oluşturur.

3. **Parametreli yapıcı fonksiyon alternatif:** `NewLanguageWithParamsShortcut` ve `NewLanguageWithParamsShortcut2`
   fonksiyonları, ad ve sembol değerlerini alan ve bu değerlerle bir `Language` örneği oluşturan alternatif
   yöntemlerdir.

Struct'lar ve yapı yöntemleri, Go dilinde karmaşık veri yapılarını kolayca temsil etmenize ve işlemenize olanak tanır.

### Go Programlama Diline Giriş: For Döngüsü

Go dilinde döngüler, belirli bir koşul sağlandığı sürece bir bloğu tekrar tekrar çalıştırmak için kullanılır. Bu
öğreticide, `for` döngüsünün farklı kullanımlarını ve özelliklerini göstereceğiz.

Öncelikle, basit bir for döngüsü kullanımını inceleyelim:

```go
package main

import "fmt"

func main() {
	// For döngüsü
	forLoop() // forLoop fonksiyonunu çağırma
}

func forLoop() {
	for i := 0; i < 10; i++ {
		value := i
		fmt.Println("Value:", value)
	}
}
```

Yukarıdaki örnekte, basit bir `for` döngüsü kullanımı gösterilmektedir. Bu döngü, 0'dan 9'a kadar olan sayıları ekrana
yazdıracaktır.

Alternatif olarak, `for` döngüsünü başlatma, koşulu belirtme ve artış adımını döngü dışında gerçekleştirebiliriz:

```go
package main

import "fmt"

func main() {
	// For döngüsü
	forLoopAlternative() // forLoopAlternative fonksiyonunu çağırma
}

func forLoopAlternative() {
	i := 0
	for i < 10 {
		value := i
		fmt.Println("Value:", value)
		i++
	}
}
```

Bu örnekte, `for` döngüsünün başlangıç durumunu döngü dışında tanımladık ve döngü koşulunu doğrudan for ifadesi içinde
belirttik.

`continue` ifadesi, döngünün bir sonraki iterasyona geçmesini sağlar:

```go
package main

import "fmt"

func main() {
	// For döngüsü
	forLoopContinue() // forLoopContinue fonksiyonunu çağırma
}

func forLoopContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		value := i
		fmt.Println("Value:", value)
	}
}
```

Bu örnekte, `i` değeri 5 olduğunda döngü, geri kalan kodu atlayacak ve bir sonraki iterasyona geçecektir.

`break` ifadesi, döngüden çıkılmasını sağlar:

```go
package main

import "fmt"

func main() {
	// For döngüsü
	forLoopBreak() // forLoopBreak fonksiyonunu çağırma
}

func forLoopBreak() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		value := i
		fmt.Println("Value:", value)
	}
}
```

Bu örnekte, `i` değeri 5 olduğunda döngü sona erecek ve döngüden çıkılacaktır.

`for` döngüsü, programınızda belirli işlemlerin tekrarlanmasını sağlar ve koşullara bağlı olarak farklı davranışlar
sergileyebilir.

### Go Programlama Diline Giriş: Diziler (Arrays)

Go dilinde, diziler aynı türden verilerin bir koleksiyonunu saklamak için kullanılır. Diziler sabit bir boyuta sahiptir
ve boyutlarını tanımlarken belirtilirler. Bu öğreticide, farklı türde dizilerin nasıl tanımlanacağını ve kullanılacağını
göstereceğiz.

Öncelikle, basit bir dizi tanımlayıp içine elemanlar ekleyerek başlayalım:

```go
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
```

Yukarıdaki örnekte, basit bir dizi tanımlaması ve farklı şekillerde diziye eleman eklemeleri gösterilmektedir.

1. **simpleArray** fonksiyonunda, boş bir dizi oluşturulur ve daha sonra elemanlar tek tek atanır.

2. **alternativeSimpleArray** fonksiyonunda, elemanlar dizi tanımının içinde doğrudan belirtilir.

3. **sampleArrayWithParameters** fonksiyonunda, fonksiyon parametreleri kullanılarak bir dizi oluşturulur.

4. **automaticSizeArray** fonksiyonunda, dizi boyutunu belirtmeden otomatik boyutlandırma yapılır.

Şimdi, gerçek hayatta dizi kullanımını göstermek için bir örnek daha ekleyelim:

```go
type Language struct {
Name     string
Symbol   string
Paradigm string
}

func main() {
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
fmt.Println(language.Name + " (" + language.Symbol + ") - " + language.Paradigm)
}
}
```

Bu örnekte, `Language` adında bir yapı tanımlıyoruz. Daha sonra, bu yapıyı kullanarak programlama dillerini temsil eden
bir dizi oluşturuyoruz. Her dil için bir `Language` yapı örneği oluşturup bunları bir diziye ekliyoruz. Son
olarak, `for` döngüsüyle bu dillerin özelliklerini ekrana yazdırıyoruz. Diziler, programlama dillerinde veri depolamanın
ve işlemenin temel yollarından biridir ve iyi bir şekilde anlaşılması önemlidir.

### Go Programlama Diline Giriş: Slice'lar

Go dilinde, slice'lar, dinamik boyutlara sahip bir dizi türüdür. Slice'lar, bir dizinin bir alt kümesini referans alır
ve bu alt küme üzerinde çalışır. Bu öğreticide, slice'ların nasıl oluşturulacağını, elemanlarının nasıl
değiştirileceğini ve yeni elemanların nasıl eklenebileceğini göstereceğiz.

Öncelikle, basit bir slice oluşturup onu yazdıralım:

```go
package main

import "fmt"

func main() {
	// Basit bir slice tanımlama
	simpleSlice()

	fmt.Println("----------------")

	// Slice elemanlarını değiştirme
	changeSliceElements()

	fmt.Println("----------------")

	// Slice'e eleman ekleme
	appendToSlice()
}

func simpleSlice() {
	// Basit bir slice tanımlama
	simpleSlices := []string{"Hello", "World", "!"}

	// Slice elemanlarını yazdırma
	fmt.Println(simpleSlices)
}

func changeSliceElements() {
	// Slice elemanlarını değiştirme
	changeSlice := []string{"Hello", "World", "!"}
	changeSlice[2] = "Go"

	// Slice elemanlarını yazdırma
	fmt.Println(changeSlice)
}

func appendToSlice() {
	// Slice'e eleman ekleme
	appendSlice := []string{"Hello", "World"}
	appendSlice = append(appendSlice, "Go")

	// Slice elemanlarını yazdırma
	fmt.Println(appendSlice)
}
```

Yukarıdaki örnekte, basit bir slice tanımlaması ve bu slice üzerinde değişiklik yapma işlemleri gösterilmektedir.

1. **simpleSlice** fonksiyonunda, basit bir slice tanımlanır ve elemanları ekrana yazdırılır.
2. **changeSliceElements** fonksiyonunda, slice elemanları değiştirilir ve sonuç ekrana yazdırılır.
3. **appendToSlice** fonksiyonunda, slice'e yeni bir eleman eklenir ve sonuç ekrana yazdırılır.

Slice'lar, Go dilinde dinamik boyutlara sahip veri yapıları oluşturmanın temel yoludur ve veri koleksiyonlarını daha
esnek bir şekilde işlemenizi sağlar.

### Go Programlama Diline Giriş: Map'ler

Go dilinde, map'ler (haritalar), anahtar-değer çiftlerini saklamak için kullanılan dinamik bir veri yapısıdır. Her
anahtar benzersiz olmalıdır ve her anahtarın bir değeri vardır. Bu öğreticide, map'lerin nasıl tanımlanacağını,
elemanlarının nasıl eklenip çıkarılacağını ve map'te eleman aramanın nasıl yapıldığını göstereceğiz.

1. **Basit Bir Map Tanımlama:**

Bir map tanımlamak için make fonksiyonunu kullanabiliriz.

```go
package main

import "fmt"

func main() {
	// Basit bir map tanımlama
	simpleMap()
}

func simpleMap() {
	// Map tanımlama
	myMap := make(map[string]string)

	// Map'e eleman ekleme
	myMap["Key1"] = "Value1"
	myMap["Key2"] = "Value2"

	// Map'teki elemanları döngü ile yazdırma
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
```

Yukarıdaki örnekte, `simpleMap` fonksiyonuyla bir map tanımlıyoruz, elemanlar ekliyoruz ve ardından döngü ile bu
elemanları yazdırıyoruz.

2. **Map'ten Eleman Arama:**

Bir anahtarı kullanarak bir map'te eleman arayabiliriz.

```go
package main

import "fmt"

func main() {
	// Map'te eleman arama
	searchInMap("Key1")
}

func searchInMap(key string) {
	// Map tanımlama
	myMap := make(map[string]string)

	// Map'e eleman ekleme
	myMap["Key1"] = "Value1"
	myMap["Key2"] = "Value2"

	// Map'te eleman arama
	searchValue, isExists := myMap[key]

	// Eleman varsa
	if isExists {
		fmt.Println("Aranan Değer:", searchValue)
	} else {
		fmt.Println("Eleman bulunamadı")
	}
}
```

Yukarıdaki örnekte, searchInMap fonksiyonuyla bir map'te belirli bir anahtarın var olup olmadığını kontrol ediyoruz.
Eğer varsa, değeri yazdırıyoruz; yoksa, bir mesaj yazdırıyoruz. Map'ler, anahtar-değer çiftlerini saklamak ve hızlı bir
şekilde erişim sağlamak için yaygın olarak kullanılır.
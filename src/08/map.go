package main

func main() {
	// Basit bir map tanımlama
	simpleMap()

	// Map'te eleman arama
	searchInMap("Key1")
}

func simpleMap() {
	// Map tanımlama
	myMap := make(map[string]string)

	// Map'e eleman ekleme
	myMap["Key1"] = "Value1"
	myMap["Key2"] = "Value2"

	// Map'teki elemanları döngü ile yazdırma
	for key, value := range myMap {
		println(key, value)
	}
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
		println(searchValue)
	} else {
		println("Eleman bulunamadı")
	}
}

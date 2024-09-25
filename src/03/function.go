package main

func main() {
	// Fonksiyon çağırma
	sayHello()

	println("--------------------------------------------------")

	// Parametreli fonksiyon çağırma
	sayHelloTo("Golang")

	println("--------------------------------------------------")

	message := "Golang"
	sayHelloTo(message)

	println("--------------------------------------------------")

	// Değer döndüren fonksiyon
	result := sum(10, 20)
	calculated(result)
}

func sayHello() {
	println("Hello World!")
}

// Parametreli fonksiyon
func sayHelloTo(name string) {
	println("Hello", name)
}

func calculated(value int) {
	println("Value:", value)
}

// Değer döndüren fonksiyon
func sum(a int, b int) int {
	return a + b
}

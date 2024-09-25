package main

import "fmt"

func main() {
	// Slice oluşturma
	sliceArray()

	// slice yeni eleman ekleme
	sliceAppend()

	// slice oluşturma make fonksiyonu ile
	sliceWithMake()

	// slice kapasitesi
	capacityOfSlice()

	// slice uzunluğu
	sliceLength()
}

func sliceArray() {
	integerArray := []int{45, 23, 43}

	// Slice oluşturma
	slice := integerArray[:]

	// Ekran çıktısı
	fmt.Println(slice)
}

func sliceAppend() {
	integerArray := []int{45, 23, 43}

	// Slice oluşturma
	slice := integerArray[:]

	// Slice'e yeni eleman ekleme
	slice = append(slice, 12)

	// Ekran çıktısı
	fmt.Println(slice)
}

func sliceWithMake() {
	// Slice oluşturma
	slice := make([]int, 3)

	// Slice'e eleman ekleme
	slice[0] = 45
	slice[1] = 23
	slice[2] = 43

	// Ekran çıktısı
	fmt.Println(slice)
}

func capacityOfSlice() {
	// Slice oluşturma
	slice := make([]int, 2)

	// Eleman ekleme
	slice = append(slice, 12)

	// Kapasite
	fmt.Println("Kapasite: ", cap(slice))
}

func sliceLength() {
	// Slice oluşturma
	slice := make([]int, 2)

	// Eleman ekleme
	slice = append(slice, 12)

	// Uzunluk
	fmt.Println("Uzunluk: ", len(slice))
}

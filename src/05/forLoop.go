package main

func main() {
	// For döngüsü
	forLoop() // forLoop fonksiyonunu çağırma

	println("----------------")

	// For döngüsü alternatif
	forLoopAlternative() // forLoopAlternative fonksiyonunu çağırma

	println("----------------")

	// For döngüsü continue
	forLoopContinue() // forLoopContinue fonksiyonunu çağırma

	println("----------------")

	// For döngüsü break
	forLoopBreak() // forLoopBreak fonksiyonunu çağırma
}

func forLoop() {
	for i := 0; i < 10; i++ {
		value := i
		println("Value:", value)
	}
}

func forLoopAlternative() {
	i := 0
	for i < 10 {
		value := i
		println("Value:", value)
		i++
	}
}

func forLoopContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		value := i
		println("Value:", value)
	}
}

func forLoopBreak() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		value := i
		println("Value:", value)
	}
}

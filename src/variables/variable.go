package variables

import "fmt"

func variables() {
	var name = "Tony" // ou var name string = "Tony"
	var age = 15 // ou var age int = 15
	var height float64 = 1.76

	//multiplas variaveis no Println
	fmt.Println("Meu nome é", name, "e minha idade é", age, "anos")
	fmt.Println(height)

	test(45, 34)
}

// Funções
func test(value int, value1 int) {
	fmt.Println(value, value1)
}

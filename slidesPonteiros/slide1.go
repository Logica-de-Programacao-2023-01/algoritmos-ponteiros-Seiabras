//Escreva uma função swap que receba dois ponteiros para int como argumentos e
//troque os valores apontados por eles.

package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	var a, b = 1, 2

	fmt.Println("Before swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	swap(&a, &b)

	fmt.Println("After swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

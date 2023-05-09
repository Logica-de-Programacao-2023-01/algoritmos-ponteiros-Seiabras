//Escreva uma função que receba um ponteiro para um booleano e altere o
//valor desse booleano para o seu inverso.

package main

import "fmt"

func boolean(b *bool) {
	*b = !*b
}

func main() {
	var x bool = true
	fmt.Println("Before reverse", x)
	boolean(&x)
	fmt.Println("After reverse", x)

}

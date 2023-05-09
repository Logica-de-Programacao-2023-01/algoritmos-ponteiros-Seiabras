//Escreva uma função que receba um ponteiro para uma struct que contém informações
//de um produto (nome, preço e quantidade em estoque). A função deve atualizar o
//preço desse produto para um novo valor fornecido como argumento.

package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func info(p *Product, newPrice float64) {
	p.Price = newPrice
}
func main() {
	prod := Product{Name: "Mobil", Price: 1500.00, Quantity: 2}
	fmt.Printf("isso: %.2f\n", prod.Price)
	oldPrice := prod.Price
	newPrice := 1800.00
	info(&prod, newPrice)

	fmt.Printf("Nome: %s\nPreço antigo: %2f\nPreço novo: %.2f\nQuantidade: %d\n", prod.Name, oldPrice, prod.Price, prod.Quantity)

}

package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price int
}

func addProduct(products []Product, p Product) []Product {
	products = append(products, p)
	return products
}

func showProducts(products []Product) {
	fmt.Println("Daftar Product:")
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", p.ID, p.Name, p.Price)
	}
}

func main() {
	products := []Product{}

	products = addProduct(products, Product{ID: 1, Name: "Laptop", Price: 12000000})
	products = addProduct(products, Product{ID: 2, Name: "Handphone", Price: 6000000})

	showProducts(products)
}

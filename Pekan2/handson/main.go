package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p *Product) updatePrice(newPrice float64) {
	p.Price = newPrice
	fmt.Printf("Harga produk '%s' diperbarui menjadi: Rp.%.2f ", p.Name, p.Price)
}

func (p *Product) updateStock(newStock int) {
	p.Stock = newStock
	fmt.Printf("dan Stok produk '%s' diperbarui menjadi: %d\n", p.Name, p.Stock)
}

func main() {
	product1 := Product{Name: "Roti", Price: 5000, Stock: 10}
	product2 := Product{Name: "Susu", Price: 15000, Stock: 30}

	fmt.Println("Data Awal Produk:")
	fmt.Printf("Produk 1: %+v\n", product1)
	fmt.Printf("Produk 2: %+v\n", product2)

	product1.updatePrice(6000)

	product1.updateStock(3)

	product2.updatePrice(17000)

	product2.updateStock(25)
}
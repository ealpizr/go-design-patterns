package main

import "fmt"

type IProduct interface {
	setName(string)
	getName() string
	setStock(int)
	getStock() int
}

type Product struct {
	name  string
	stock int
}

func (p *Product) setName(name string) {
	p.name = name
}

func (p Product) getName() string {
	return p.name
}

func (p *Product) setStock(stock int) {
	p.stock = stock
}

func (p Product) getStock() int {
	return p.stock
}

type LaptopComputer struct {
	Product
}

func newLaptopComputer(stock int) IProduct {
	return &LaptopComputer{
		Product: Product{
			name:  "Laptop Computer",
			stock: stock,
		},
	}
}

type DesktopComputer struct {
	Product
}

func newDesktopComputer(stock int) IProduct {
	return &DesktopComputer{
		Product: Product{
			name:  "Desktop Computer",
			stock: stock,
		},
	}
}

func printProductInfo(p IProduct) {
	fmt.Printf("Product: %s\nStock: %d\n\n", p.getName(), p.getStock())
}

func ProductFactory(productType string, stock int) (IProduct, error) {
	switch productType {
	case "laptop":
		return newLaptopComputer(stock), nil
	case "desktop":
		return newDesktopComputer(stock), nil
	default:
		return nil, fmt.Errorf("got an invalid product type (%s)", productType)
	}
}

func main() {
	d, _ := ProductFactory("desktop", 12)
	l, _ := ProductFactory("laptop", 7)

	printProductInfo(d)
	printProductInfo(l)
}

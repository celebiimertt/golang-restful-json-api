package main

import (
	"fmt"
	"golesson/proje"
)

func main() {

	addProduct, _ := proje.AddProduct()
	fmt.Println(addProduct)
	product, _ := proje.GetProducts()
	for i := 0; i < len(product); i++ {
		fmt.Println(product[i].ProductName)
	}

}

package proje

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}
type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err //hata varsa product boş,err dolu döner
	}
	defer response.Body.Close()
	bodyByte, _ := ioutil.ReadAll(response.Body)
	products := []Product{} // "var products []Product" ile aynı anlama gelir.
	//Birden fazla ürün olduğu için dizi olarak oluşturduk.
	json.Unmarshal(bodyByte, &products) //products,bellek adresi ile yollanmazsa yeni bir nesne gibi davranır ve products değişmezdi.'&' işareti ile products değişmiş oldu.
	return products, nil                //hata yoksa product dolu,err boş yani nil döner

}
func AddProduct() (Product, error) {
	//product := Product{10, "MotherBoard", 1, 3000.99}
	product := Product{ProductName: "MotherBoard2", CategoryId: 1, UnitPrice: 3000.99}
	//id belirtilmezse otomatik en son id'yi bir arttırarak ekler fakat parametreleri nameleri birlikte eklenmesi gerekir.
	productsBytes, err := json.Marshal(product)
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(productsBytes))
	if err != nil {
		return Product{}, err //geri dönüş için bir dizi değil de struct istediğimiz için nil kullanamadık.Onun yerine boş bir Product structı kullandık.
	}
	defer response.Body.Close()
	//Verinin eklenip eklenmediğini sistem üzerinden kontrol edilebilir.
	//Terminal üzerinden de kontrolü ise gönderdiğimiz jsonu tekrar product type'ına dönüştürerek yapabiliriz.
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	products2 := Product{}
	json.Unmarshal(bodyBytes, &products2)
	return products2, nil

}

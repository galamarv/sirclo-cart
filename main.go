package main

import "fmt"

var cart = make(map[string]int)

func tambahProduk(kodeProduk string, kuantitas int) {

	if v, ok := cart[kodeProduk]; ok {
		cart[kodeProduk] = v + kuantitas
	} else {
		cart[kodeProduk] = kuantitas
	}

}

func hapusProduk(kodeProduk string) {
	delete(cart, kodeProduk)
}

func tampilkanCart() {

	for k, v := range cart {
		fmt.Printf("%v (%v)\n", k, v)
	}

}

func main() {

	tambahProduk("Topi Putih", 2)

	tambahProduk("Kemeja Hitam", 3)

	tambahProduk("Sepatu Merah", 1)
	tambahProduk("Sepatu Merah", 4)
	tambahProduk("Sepatu Merah", 2)

	hapusProduk("Kemeja Hitam")

	hapusProduk("Baju Hijau")

	tampilkanCart()

}

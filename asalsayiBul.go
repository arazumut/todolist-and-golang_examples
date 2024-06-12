package main

import (
	"fmt"
)

// Asal sayı kontrolü yapan fonksiyon
func asalKontrol(sayi int) bool {
	if sayi < 2 {
		return false // 2'den küçük sayılar asal değildir
	}

	for i := 2; i*i <= sayi; i++ {
		if sayi%i == 0 {
			return false // Sayı bir böleni varsa asal değildir
		}
	}

	return true // Sayı asaldır
}

//! Produced by Kamil Umut Araz instagram : K.umutarazz  github : arazumut

func main() {
	var sayi int

	// Kullanıcıdan bir sayı al
	fmt.Print("Bir sayı giriniz: ")
	fmt.Scanln(&sayi)

	// Asal kontrol fonksiyonunu çağır ve sonucu ekrana yazdır
	if asalKontrol(sayi) {
		fmt.Printf("%d sayısı bir asal sayıdır.\n", sayi)
	} else {
		fmt.Printf("%d sayısı bir asal sayı değildir.\n", sayi)
	}
}

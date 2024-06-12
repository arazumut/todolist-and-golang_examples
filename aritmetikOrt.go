package main

import "fmt"

func main() {
	var n int

	//! Produced by Kamil Umut Araz instagram : K.umutarazz  github : arazumut

	fmt.Print("Kaç adet sayı gireceksiniz: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("Geçersiz giriş. Pozitif bir sayı bekleniyor.")
		return // Hata kodu
	}

	var toplam, sayi float64

	fmt.Printf("%d adet sayıyı girin:\n", n)

	for i := 0; i < n; i++ {
		fmt.Printf("Sayı %d: ", i+1)
		fmt.Scanln(&sayi)
		toplam += sayi
	}

	aritmetikOrtalama := toplam / float64(n)

	fmt.Printf("Aritmetik Ortalama: %f\n", aritmetikOrtalama)
}

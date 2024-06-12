package main

import "fmt"

// Dizideki en büyük elemanı bulan fonksiyon
func enBuyukEleman(dizi []int) int {
	max := dizi[0]

	for _, eleman := range dizi {
		if eleman > max {
			max = eleman
		}
	}

	return max
}

func main() {
	var n int

	// Kullanıcıdan dizinin boyutunu al
	fmt.Print("Dizinin boyutunu giriniz: ")
	fmt.Scanln(&n)

	dizi := make([]int, n)

	// Kullanıcıdan diziyi al
	fmt.Println("Diziyi giriniz:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&dizi[i])
	}

	// En büyük elemanı bul ve ekrana yazdır
	fmt.Printf("En büyük eleman: %d\n", enBuyukEleman(dizi))
}

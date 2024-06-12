package main

import (
	"fmt"
)

func main() {
	var kelime string

	fmt.Print("Lütfen bir kelime giriniz: ")
	fmt.Scanln(&kelime)

	uzunluk := len(kelime)
	kelime2 := make([]byte, uzunluk)

	for j := 0; j < uzunluk/2; j++ {
		gecici := kelime[uzunluk-j-1]
		kelime2[uzunluk-j-1] = kelime[j]
		kelime2[j] = gecici
	}

	fmt.Printf("Girilen kelimenin tersi: %s\n", kelime2)
}

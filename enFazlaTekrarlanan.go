package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_CHARACTERS = 200

func main() {
	var metin string
	karakterSayisi := make(map[rune]int)
	var maxTekrar int
	var enCokTekrarlananKarakter rune

	fmt.Print("Metni giriniz (en fazla 200 karakter): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		metin = scanner.Text()
	}


	// Metindeki her bir karakterin sayısını say
	for _, karakter := range metin {
		if karakter != '\n' { // Satır sonu karakterini sayma
			karakterSayisi[karakter]++
		}
	}

	

	// En fazla tekrarlanan karakteri bul
	for karakter, sayi := range karakterSayisi {
		if sayi > maxTekrar {
			maxTekrar = sayi
			enCokTekrarlananKarakter = karakter
		}
	}

	fmt.Printf("En fazla tekrarlanan karakter: %c\n", enCokTekrarlananKarakter)
	fmt.Printf("Tekrar sayısı: %d\n", maxTekrar)
}

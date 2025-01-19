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

    for j := 0; j < uzunluk; j++ {
        kelime2[j] = kelime[uzunluk-j-1]
    }

    fmt.Printf("Girilen kelimenin tersi: %s\n", kelime2)
}

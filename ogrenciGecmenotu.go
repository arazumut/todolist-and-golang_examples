package main

import (
    "fmt"
)

// Ders yapısı
type Ders struct {
    dersKod   int
    vize      int
    final     int
    gecmeNotu float64
}

// SahisBilgileri yapısı
type SahisBilgileri struct {
    ad      string
    soyad   string
    no      int
    dersler [5]Ders
}

func ogrenciBilgileriGirVeHesapla(ogrenci *SahisBilgileri) {
    fmt.Print("Ad: ")
    fmt.Scanln(&ogrenci.ad)

    fmt.Print("Soyad: ")
    fmt.Scanln(&ogrenci.soyad)

    fmt.Print("Ogrenci No: ")
    fmt.Scanln(&ogrenci.no)

    for i := 0; i < 5; i++ {
        fmt.Printf("Ders %d Kod: ", i+1)
        fmt.Scanln(&ogrenci.dersler[i].dersKod)

        fmt.Printf("Ders %d Vize: ", i+1)
        fmt.Scanln(&ogrenci.dersler[i].vize)

        fmt.Printf("Ders %d Final: ", i+1)
        fmt.Scanln(&ogrenci.dersler[i].final)

        ogrenci.dersler[i].gecmeNotu = 0.4*float64(ogrenci.dersler[i].vize) + 0.6*float64(ogrenci.dersler[i].final)
    }
}

func ogrenciBilgileriniYazdir(ogrenci SahisBilgileri) {
    fmt.Println("\nOgrenci Bilgileri:")
    fmt.Printf("Ad: %s\n", ogrenci.ad)
    fmt.Printf("Soyad: %s\n", ogrenci.soyad)
    fmt.Printf("Ogrenci No: %d\n", ogrenci.no)

    fmt.Println("\nDers Bilgileri:")
    for i := 0; i < 5; i++ {
        fmt.Printf("Ders %d Kod: %d\n", i+1, ogrenci.dersler[i].dersKod)
        fmt.Printf("Ders %d Vize: %d\n", i+1, ogrenci.dersler[i].vize)
        fmt.Printf("Ders %d Final: %d\n", i+1, ogrenci.dersler[i].final)
        fmt.Printf("Ders %d Gecme Notu: %.2f\n", i+1, ogrenci.dersler[i].gecmeNotu)
    }
}

func main() {
    var ogrenciler [5]SahisBilgileri

    for i := 0; i < 5; i++ {
        fmt.Printf("=== %d. ogrenci ===\n", i+1)
        ogrenciBilgileriGirVeHesapla(&ogrenciler[i])
        fmt.Println()
    }

    for i := 0; i < 5; i++ {
        ogrenciBilgileriniYazdir(ogrenciler[i])
        fmt.Println("================")
    }
}

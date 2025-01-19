package main

import (
    "fmt"
)

// Kargo yapısı
type Kargo struct {
    kargoNo int
    durum   string
}

func main() {
    const MAX_KARGO_SAYISI = 100
    kargoListesi := make([]*Kargo, MAX_KARGO_SAYISI) // Kargo numaraları için işaretçi dizisi
    kargoSayisi := 0

    // Kullanıcıdan kargo numaralarını al
    fmt.Println("Kargo takibi yapmak için kargo numaralarını girin (sonlandırmak için -1):")
    var kargoNo int
    for {
        fmt.Print("Kargo No: ")
        fmt.Scanln(&kargoNo)
        if kargoNo == -1 {
            break
        }

        // bellekte yer aç
        kargoListesi[kargoSayisi] = &Kargo{
            kargoNo: kargoNo,
        }

        // Kargo durumunu kaydet
        fmt.Print("Kargo Durumu: ")
        fmt.Scanln(&kargoListesi[kargoSayisi].durum)

        kargoSayisi++
    }

    // Kargo bilgilerini ekrana yazdır
    fmt.Println("\nKargo Takibi:")
    for i := 0; i < kargoSayisi; i++ {
        fmt.Printf("Kargo No: %d, Durum: %s\n", kargoListesi[i].kargoNo, kargoListesi[i].durum)
    }
}

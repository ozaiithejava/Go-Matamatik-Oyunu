package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Oyun başlangıcı
	fmt.Println("Matematik Oyunu!")
	fmt.Println("-----------------")

	puan := 0
	toplamSoru := 0

	for soru := 0; soru < 10; soru++ {
		// Rastgele bir işlem ve sayılar seçme
		islem, sayi1, sayi2 := rastgeleSoru()

		// Kullanıcıdan cevap alma
		fmt.Printf("%d. Soru: %d %s %d = ", soru+1, sayi1, islem, sayi2)
		var kullaniciCevap int
		fmt.Scan(&kullaniciCevap)

		// Cevap kontrolü
		dogruCevap := hesapla(islem, sayi1, sayi2)
		if kullaniciCevap == dogruCevap {
			fmt.Println("Doğru!")
			puan++
		} else {
			fmt.Printf("Yanlış! Doğru cevap: %d\n", dogruCevap)
		}

		toplamSoru++
	}

	// Oyun sonu istatistikleri
	fmt.Printf("\nOyun Bitti! Toplam Puan: %d/%d\n", puan, toplamSoru)
}

// Rastgele bir işlem ve sayılar seçme
func rastgeleSoru() (string, int, int) {
	islemler := []string{"+", "-", "*", "/"}

	rastgeleIndex := rand.Intn(len(islemler))
	rastgeleIslem := islemler[rastgeleIndex]

	var sayi1, sayi2 int

	switch rastgeleIslem {
	case "+", "-":
		sayi1 = rand.Intn(100)
		sayi2 = rand.Intn(100)
	case "*":
		sayi1 = rand.Intn(15) + 1
		sayi2 = rand.Intn(15) + 1
	case "/":
		sayi2 = rand.Intn(10) + 1
		sayi1 = sayi2 * (rand.Intn(10) + 1)
	}

	return rastgeleIslem, sayi1, sayi2
}

// İşleme göre doğru sonucu hesapla
func hesapla(islem string, sayi1, sayi2 int) int {
	switch islem {
	case "+":
		return sayi1 + sayi2
	case "-":
		return sayi1 - sayi2
	case "*":
		return sayi1 * sayi2
	case "/":
		return sayi1 / sayi2
	default:
		return 0
	}
}

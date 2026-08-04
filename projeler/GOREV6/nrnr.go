package main

import "fmt"

func main() {
	// Fonksiyondan dönen 2 değeri 'ipAdresi' ve 'guvenliMi' değişkenlerine atayarak yakalıyoruz.
	ipAdresi, guvenliMi := hedefOzeti("1.1.1.1", 80)

	// Yakaladığımız değerleri ekrana yazdırıyoruz.
	fmt.Printf("Hedef Adres: %s\nGüvenli mi?: %t\n", ipAdresi, guvenliMi)

}

func hedefOzeti(ip string, port int) (adres string, guvenli bool) {
	adres = ip
	guvenli = false
	return
}

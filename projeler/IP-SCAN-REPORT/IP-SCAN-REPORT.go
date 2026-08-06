package main

import (
	"fmt"
)

const (
	ip1 = "192.168.1.1"
	ip2 = "192.168.1.2"
	ip3 = "192.168.1.3"
)

func main() {
	// 1. Ip sabitlerimizi bir slice yani liste içinde topluyorz
	hedefler := []string{ip1, ip2, ip3}

	// Her bir ip ye karsılık gelen portları tanımlıyoruz
	portlar := [][]int{{80, 443, 22}, {3306, 5432}, {8080}}

	var tHedef int = 0
	var tPort int = 0

	defer func() {
		fmt.Printf("Toplam hedef : [%d] | Toplam Açik Port : [%d] | Ortalama : [%.1f]\n", tHedef, tPort, float64(tPort)/float64(tHedef))
	}() // Sondaki () parantezler bu isimsiz fonksiyonu anında çağırmaya yarar

	for i, ip := range hedefler {
		tHedef++
		tPort += len(portlar[i])
		fmt.Printf("Hedef: %s | Açik Portlar: [%v]\n", ip, portlar[i])
	}

}

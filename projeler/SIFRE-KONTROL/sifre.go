package main

import (
	"fmt"
	"unicode"
)

func main() {
	var pass string
	fmt.Print("Şifrenizi giriniz:  ")
	// Boşluk ve satır sonunu dogru yakalamk ıcın scanLn kullandık
	fmt.Scanln(&pass)
	// İçerik kontrolunde kullanılacak degıskenlerı tanımlarız
	var hLetter, hDigit, hSymbol bool
	// Şifrenin içindeki her bir karakteri (rune) tek tek döndürüyoruz

	for _, char := range pass {
		/* Go dilinde range pass dedıgımızde her bır karakterın numarası olur
		 normalde ve go bunu kullanmamaızı ıster bız ise " _ " kullanarak
		bu degerı cope atarız cunku bız ındis ile ilgilenmiyoruz.*/
		if unicode.IsLetter(char) {
			hLetter = true
		} else if unicode.IsDigit(char) {
			hDigit = true
		} else {
			// Harf veya rakam degılse sembol olarak kabul edılır
			hSymbol = true
		}
	}

	var contentMSG string
	if hLetter && hDigit && hSymbol {
		contentMSG = "(Harf + Rakam + Sembol)"
	} else if hLetter && hDigit {
		contentMSG = "(Harf + Rakam)"
	} else if hDigit && !hLetter && !hSymbol {
		contentMSG = "(Sadece rakam)"
	} else if hLetter && !hDigit && !hSymbol {
		contentMSG = "(Sadece harf)"
	} else {
		contentMSG = "(Özel İçerik)"
	}

	// Switch kullanarak uzunluk kontrolü yaparız
	fmt.Print("Şifre gücü: ")
	switch {
	case len(pass) <= 5:
		fmt.Printf("Çok kisa %s\n", contentMSG)
	case len(pass) >= 6 && len(pass) <= 8:
		fmt.Printf("Zayif %s\n", contentMSG)
	case len(pass) >= 9 && len(pass) <= 12:
		fmt.Printf("Orta %s\n", contentMSG)
	case len(pass) >= 13:
		fmt.Printf("Güçlü %s\n", contentMSG)
	}

}

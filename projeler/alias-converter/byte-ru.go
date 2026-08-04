package main

import (
	"fmt"
)

func main() {
	harf := 'A'
	parola := "secret"
	var z int32 = rune(harf)
	var x []byte = []byte(parola)

	fmt.Printf("HARF = %d || Parola = %v", z, x)

}

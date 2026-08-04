package main

import (
	"fmt"
)

var (
	minPort, maxPort int    = 1, 65535
	protokol         string = "TCP"
)

func main() {
	fmt.Printf(`Minimum port number is = %d\n
	            Maximum port number is = %d\n
				Protokol name is = %s\n`, minPort, maxPort, protokol)
}

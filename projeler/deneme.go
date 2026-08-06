package main

import (
	"fmt"
)

func main() {
	/*fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")

	default:
		// freebsd , openbsd ,windows
		fmt.Printf("%s.\n", os)
	}*/

	/*fmt.Println("Cumartesi ne zaman ? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}*/

	/*defer fmt.Println("world")

	fmt.Println("hello")
	fmt.Println("world-wide-web")*/

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")

}

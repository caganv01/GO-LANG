package main

import "fmt"

// or import "fmt"
// or import "math/rand"

var x, y int = 1, 2

func main() {
	// fmt.Println("Bu bir denemedir .", rand.Intn(10))
	// fmt.Printf("Örnek bir %g. deneme.\n", math.Sqrt(4))
	// fmt.Println(add(42, 13))

	// a, b := swap("Merhabalar", "dünya")
	// fmt.Println(a, b)

	/*k := 3
	var c, python, java = true, false, "no!"
	fmt.Println(x, y, c, python, java, k)*/

	/*var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)*/

	v := 42 // değiştirmeyi deneyecegiz
	fmt.Printf("v tipi %T\n", v)

}

/*var (
	state  bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)*/

/*func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return x, y
}
func split(sum int) (x, y int) {
	x = sum / 10
	y = sum - x
	return
}*/

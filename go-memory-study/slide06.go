package slide06

import "fmt"

func main() {
	a := new(int)
	fmt.Println(a)

	b := new(int)
	*b = 0xdeadbeef
}

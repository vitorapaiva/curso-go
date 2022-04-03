package cap4

import (
	"fmt"
)

func Cap4() {
	booleanType()
	numberType()
}

func booleanType() {
	fmt.Println("booleanType")
	var x bool
	fmt.Println(x)
	x = (10 > 20)
	y := (10 == 10)
	z := (10 < 5)
	fmt.Println(x, y, z)
}

func numberType() {
	fmt.Println("numberType")
	var x bool
	fmt.Println(x)
	x = (10 > 20)
	y := (10 == 10)
	z := (10 < 5)
	fmt.Println(x, y, z)
}

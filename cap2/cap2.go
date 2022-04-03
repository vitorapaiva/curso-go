package cap2

import (
	"fmt"
)

type hotdog int
type loxas int

var b hotdog
var zemanilha loxas

var y = 10
var valorZeroIntVar int
var valorZerStringVar string
var valorZeroFloatVar float64
var valorZeroBoolVar bool
var exec2X int
var exec2Y string
var exec2Z bool
var exec5Y int

func Cap2() {
	x := 16
	helloWorld()
	operadorCurto()
	palavraChaveVar(x)
	valorZero()
	fmtFunctions()
	customTypes()
	conversion()
	exec1()
	exec2()
	exec3()
	exec4()
	exec5()
}

func helloWorld() {
	fmt.Println("helloWorld")
	x := 16
	y := "Hello world"
	z := true
	fmt.Println(x, y, z)
}

func operadorCurto() {
	fmt.Println("operadorCurto")
	x := 16
	y := "operadorCurto"
	z := true
	d := 16.0
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("d: %v, %T\n\n", d, d)

	x = 17
	y = "novo texto"
	z = false
	d = 16.30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("d: %v, %T\n\n", d, d)
}

func palavraChaveVar(x int) {
	fmt.Println("palavraChaveVar")
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)
}

func valorZero() {
	fmt.Println("valorZero")
	fmt.Printf("valorZero: %v, %T\n", valorZeroIntVar, valorZeroIntVar)
	fmt.Printf("valorZero: %v, %T\n", valorZerStringVar, valorZerStringVar)
	fmt.Printf("valorZero: %v, %T\n", valorZeroFloatVar, valorZeroFloatVar)
	fmt.Printf("valorZero: %v, %T\n\n", valorZeroBoolVar, valorZeroBoolVar)
}

func fmtFunctions() {
	fmt.Println("fmtFunctions")
	x := "oi"
	y := "bom dia"
	z := fmt.Sprint(x, " ", y)
	fmt.Println(z, "\n")
}

func customTypes() {
	fmt.Println("customTypes")
	fmt.Printf("%T\n\n", b)
}

func conversion() {
	fmt.Println("conversion")
	fmt.Printf("%T\n\n", b)
	x := int(b)
	fmt.Printf("%T\n\n", x)
}

func exec1() {
	fmt.Println("exec1")
	x := 42
	y := "James bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z, "\n")
}

func exec2() {
	fmt.Println("exec2")
	fmt.Println(exec2X)
	fmt.Println(exec2Y)
	fmt.Println(exec2Z, "\n")
	//valor zero
}

func exec3() {
	fmt.Println("exec3")
	exec2X = 42
	exec2Y = "James Bond"
	exec2Z = true
	s := fmt.Sprintf("%v\n%v\n%v", exec2X, exec2Y, exec2Z)
	fmt.Println(s, "\n")
}

func exec4() {
	fmt.Println("exec4")
	fmt.Println(zemanilha)
	fmt.Printf("%T\n", zemanilha)
	zemanilha = 42
	fmt.Println(zemanilha)
}

func exec5() {
	fmt.Println("exec5")
	fmt.Println(zemanilha)
	zemanilha = 42
	fmt.Println(zemanilha)
	exec5Y := int(zemanilha)
	fmt.Printf("%v\n%T\n\n", exec5Y, exec5Y)
}

func exec6() {

}

package main

import "fmt"

func main() {
	// exercicio1()
	// exercicio2()
	// exercicio3()
	// exercicio4()
	exercicio5()

}

// Solução para o exercicio 1
func exercicio1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
	fmt.Printf("z = %v\n", z)
}

// Solução para o exercicio 2
var x int
var y string
var z bool

func exercicio2() {
	// Quando se declara variáveis sem atribuir valores
	// O Compilador atribui valores chamados "valor zero"
	fmt.Println(x, y, z)
}

// Solução para o exercicio 3
var xx = 42
var yy = "James Bond"
var zz = true

func exercicio3() {
	s := fmt.Sprintln(xx, yy, zz)
	fmt.Println(s)
}

// Solução para o exercicio 4
type tipo int

var xTipo tipo

func exercicio4() {
	fmt.Println(xTipo)
	fmt.Printf("%T\n", xTipo)

	xTipo = 42
	fmt.Println(xTipo)
}

// Solução para o exercicio 5
type yDeTipo tipo

var yTipo int

func exercicio5() {
	fmt.Println(xTipo)
	fmt.Printf("%T\n", xTipo)

	xTipo = 42
	fmt.Println(xTipo)

	yTipo = int(xTipo)
	fmt.Println(yTipo)
	fmt.Printf("%T\n", yTipo)
}

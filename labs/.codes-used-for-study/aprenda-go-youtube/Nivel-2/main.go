package main

import "fmt"

func main() {
	// exercicio1()
	// exercicio2()
	// exercicio3()
	// exercicio4()
	// exercicio5()
	exercicio6()
}

func exercicio1() {
	x := 10
	fmt.Printf("%d, %#x, %b", x, x, x)
}

func exercicio2() {
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 >= 100)
	f := (10 > 100)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}

const x int = 10
const y = 10

func exercicio3() {
	fmt.Println(x, y)
}

func exercicio4() {
	a := 200
	fmt.Printf("%d, %b, %#x", a, a, a)
	fmt.Println()

	b := a << 1
	fmt.Printf("%d, %b, %#x", b, b, b)
}

func exercicio5() {
	x := `into
		é 
			uma coisa
				muito doida!
	`
	fmt.Println(x)
}

func exercicio6() {
	const (
		_ = 1994 + iota
		b
		c
		d
		e
	)
	fmt.Println(b, c, d, e)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Sto usando il package math, la radice di 4 è %g .\nCapito? \n", math.Sqrt(4))
	fmt.Printf("Inoltre la somma tra 2, 3 e PiGreco è %g .\n", somma(2, 3))
	fmt.Println("ciao, test compilazione")
	fmt.Println(divisione(17))
}

func somma(x int, y int) float32 {
	return (float32(x+y) + math.Pi)
}

func divisione(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

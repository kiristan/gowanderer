package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/kiristan/stringutil"
)

const Pi = 3.14
const (
	Big   = 1 << 50
	Small = Big >> 99
)

func main() {

	fmt.Println(stringutil.Reverse("Reversed !oG ,olleH"))
	fmt.Println("Hello, world.")

	var x, y int = 3, 7
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, f)

	v := 42 // change me!
	fmt.Printf("v is of type %T\nf if of type %T\n", v, f)
	fmt.Println("La variaile Pi", Pi, "è questa")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Printf("test %v\n", Big)
	fmt.Println(Small)

	//osSystem()
	dayOfWeek()

	deferFunc()
}

//arrivato a https://tour.golang.org/moretypes/1

func deferFunc() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}

	fmt.Println("done")
}

func osSystem() {

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func dayOfWeek() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today + 2)
	switch time.Saturday {
	case today + 0:
		fmt.Printf("Today is %v.\n", today)
	case today + 1:
		fmt.Printf("Tomorrow is %v.\n", today)
	case today + 2:
		fmt.Printf("In two days, today is %v.\n", today)
	default:
		fmt.Printf("Too far away, today is %v.\n", today)
	}
}

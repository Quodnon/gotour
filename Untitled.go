package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Tht time is", time.Now() , rand.Intn(10))
	fmt.Println(math.Nextafter(5,9))
	fmt.Println(add(5,3));
}

func add(x int , y int ) int {
	return x+y;
}
package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"

)

func add(x , y int ) int {
	return x+y;
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Tht time is", time.Now() , rand.Intn(10))
	fmt.Println(math.Nextafter(5,9))
	fmt.Println(add(6,3));

	a,b:=swap("TY","IJAKK");
	fmt.Println(a,b);
}


func swap (x,y string) (string, string) {
	return y,x;
}
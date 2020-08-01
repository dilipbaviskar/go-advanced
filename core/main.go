package main

import "fmt"

func main() {

	var x uint = 0x0009 //0000 1001
	var y uint = 0X0040 //0100 0001

	var z uint

	z = x & y
	fmt.Println("x & y  =", z)

	z = x | y
	fmt.Println("x | y  =", z)

	z = x ^ y
	fmt.Println("x ^ y  =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1 =", z)

	b := []byte("ABC€")
	fmt.Println(b) // [65 66 67 226 130 172]
}

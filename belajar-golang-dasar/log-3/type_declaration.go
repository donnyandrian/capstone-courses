package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "11111111111111"

	var contoh string = "2222222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)
}

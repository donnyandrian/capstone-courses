package main

import (
	"belajar-golang-dasar/log-5/database"
	_ "belajar-golang-dasar/log-5/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}

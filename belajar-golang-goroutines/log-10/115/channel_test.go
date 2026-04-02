package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 1)
	defer close(channel)

	go func() {
		channel <- "Eko"
		channel <- "Kurniawan"
		channel <- "Khannedy"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

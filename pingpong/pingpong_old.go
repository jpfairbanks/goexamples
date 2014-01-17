/* pingpong: a simple 1 channel two go routine game of ping pong.

Author: James Fairbanks
Date: 2013-12-23
License : BSD
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var exit_status int
	fmt.Print("Let's Play Ping Pong\n")
	var ch chan int
	go ping(ch)
	go pong(ch)
	for {
		time.Sleep(10000)
	}
	os.Exit(exit_status)
}

func ping(ch chan int) {
	init := 0
	ch <- init
	for x := range ch {
		fmt.Printf("Ping: %v\n", x)
		ch <- x + 1
	}
}

func pong(ch chan int) {
	for x := range ch {
		fmt.Printf("Pong: %v\n", x)
		ch <- x + 1
	}
}

/* pingpong: a simple 1 channel two go routine game of ping pong.

Author: James Fairbanks
Date: 2013-12-23
License : BSD
*/

package main

import (
	"fmt"
	"os"
	"time"
)

type datum struct {
	tag  string
	data int
}

func main() {
	var exit_status int
	fmt.Print("Let's Play Ping Pong\n")
	var ch chan datum
	ch = make(chan datum, 0)
	go ping(ch, "A")
	go ping(ch, "B")
	go pong(ch)
	for {
		time.Sleep(100000)
	}
	os.Exit(exit_status)
}

func ping(ch chan datum, tag string) {
	fmt.Printf("pinging")
	for x := 0; x < 1000000; x++ {
		fmt.Printf("%v Ping: %v\n", tag, x)
		ch <- datum{tag, x}
		time.Sleep(100000)
	}
}

func pong(ch chan datum) {
	fmt.Printf("ponging")
	for dat := range ch {
		fmt.Printf("%v Pong: %v\n", dat.tag, dat.data)
		//ch <- x + 1
		time.Sleep(100000)
	}
}

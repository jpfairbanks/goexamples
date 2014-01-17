package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.Tick(500 * time.Millisecond)
	longch := time.Tick(1000 * time.Millisecond)
	for {
		select {
		case t := <-ch:
			fmt.Println(t)
		case t := <-longch:
			fmt.Println("long", t)
		}
	}
}

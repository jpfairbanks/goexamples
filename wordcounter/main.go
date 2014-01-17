/* wordcounter: Reads a file and counts the total number of words.
Produces the output as a stream of words counted so far

Author: James Fairbanks
License: BSD
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	//"strings"
)

func main() {
	file, err := os.Open("file.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	scanbuffer := 0 //one word at a time
	datach := make(chan string, scanbuffer)
	go getwords(scanner, datach)
	//ask for the count at regular intervals
	query_interval := 5 * time.Millisecond
	quetick := time.Tick(query_interval)
	// Count the words.
	count := 0
	for {
		select {
		case t := <-quetick:
			fmt.Printf("%v;%d\n", t, count)

		case _, ok := <-datach:
			if ok { //ok is false after close(datach)
				count++
			} else {
				//closed channels are always ready nil channels never are
				datach = nil
			}
		}
	}
}

func getwords(scanner *bufio.Scanner, ch chan<- string) {
	for scanner.Scan() {
		s := scanner.Text()
		//fmt.Fprintln(os.Stdout, s)
		ch <- s
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Fprintln(os.Stderr, "DONE reading input")
	close(ch)
}

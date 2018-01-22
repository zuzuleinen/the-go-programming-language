package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	withStringsJoin()
	fmt.Println(time.Since(start).Nanoseconds())
}

func example() {
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

//Modify the echo program to also print os.Args[0], the name of the command that invoked it
//Duration in nanosec with 15 elems 23293, 19354, 20348
func exerciseOne() {
	var s, sep string

	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

//Modify the echo program to print the index and value of each arguments one per line
func exerciseTwo() {
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
}

//Experiment to measure the difference in running time between our potentially
//inefficient versions and the one that uses string.Join
func exerciseThree() {

}

//Duration in nanosec with 15 elems 16957,16901, 16669
func withStringsJoin() {
	fmt.Println(strings.Join(os.Args, " "))
}

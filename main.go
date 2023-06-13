package main

import (
	"fmt"
	"lindy2076/primeDates/primes"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
	fmt.Println(primes.PrimalityString(17))
	fmt.Println("Please enter a date in iso format like \"2000-01-01\"")

	var year, month, day uint32
	_, err := fmt.Scanf("%d-%d-%d", &year, &month, &day)
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(year, month, day)
}

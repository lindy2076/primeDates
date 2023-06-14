package main

import (
	"fmt"
	"lindy2076/primeDates/date"
	"lindy2076/primeDates/primes"
	"log"
)

func main() {
	// fmt.Println(os.Args[1:])
	fmt.Println("Please enter a date in iso format like \"2000-01-01\"")

	var userInput string
	_, err := fmt.Scanf("%s", &userInput)
	if err != nil {
		log.Fatal("err: ", err)
	}

	d, err := date.BuildDateFromIso(userInput)
	if err != nil {
		log.Fatal("err: ", err)
	}

	printFun(d)

}

func printFun(d *date.Date) {
	day, month, year := d.Day(), d.Month(), d.Year()

	dp := primes.IsPrime(uint32(day))
	mp := primes.IsPrime(uint32(month))
	yp := primes.IsPrime(year)

	if dp && mp && yp {
		fmt.Println("---- Congratulations! The date is prime. ----")
	} else {
		fmt.Println("Sadly, the date is not prime...")
	}
	fmt.Printf("Day %s\n", primes.PrimalityString(uint32(day)))
	fmt.Printf("Month %s\n", primes.PrimalityString(uint32(month)))
	fmt.Printf("Year %s\n", primes.PrimalityString(year))
}

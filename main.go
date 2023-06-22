package main

import (
	"fmt"
	"lindy2076/primeDates/date"
	"lindy2076/primeDates/primes"
	"log"
	"os"
)

func printHello() {
	fmt.Print("WELCOME TO PRIMEDATES!\nThis program is designed to help" +
		" you find prime dates.\nA date is prime if it consists of" +
		" prime numbers only (e.g. 2011-03-05).\nYou can enter it below or pass it in " +
		" args (iso format 'yyyy-mm-dd').\n\n")
	printHelp()
}

func printHelp() {
	fmt.Println("Please enter a date in iso format like \"2000-01-01\"")
}

func main() {
	// get the date
	var userInput string

	if len(os.Args[1:]) > 0 {
		userInput = os.Args[1]
	} else {
		printHello()
		_, err := fmt.Scanf("%s", &userInput)
		if err != nil {
			log.Fatal("err: ", err)
		}
	}

	d, err := date.BuildDateFromIso(userInput)
	if err != nil {
		log.Fatal("err: ", err)
	}

	printInfoAboutDate(d)

	// search for prime dates in the year
	yearToResearch := d.Year()
	if !primes.IsPrime(yearToResearch) {
		return
	}

	primeDates := primeDatesInYear(yearToResearch)
	printPrimeDates(primeDates)
}

func dateIsPrime(d *date.Date) bool {
	day, month, year := d.Day(), d.Month(), d.Year()

	dp := primes.IsPrime(uint32(day))
	mp := primes.IsPrime(uint32(month))
	yp := primes.IsPrime(year)

	return dp && mp && yp
}

func printInfoAboutDate(d *date.Date) {
	day, month, year := d.Day(), d.Month(), d.Year()

	if dateIsPrime(d) {
		fmt.Println("---- Congratulations! The date is prime. ----")
	} else {
		fmt.Println("Sadly, the date is not prime:")
	}
	printEachValuePrimality(day, month, year)
}

func printEachValuePrimality(day, month uint, year uint32) {
	fmt.Printf(" Day %s\n", primes.PrimalityString(uint32(day)))
	fmt.Printf(" Month %s\n", primes.PrimalityString(uint32(month)))
	fmt.Printf(" Year %s\n", primes.PrimalityString(year))
}

func primeDatesInYear(year uint32) []*date.Date {
	startDate, err := date.BuildDate(1, 1, year)
	if err != nil {
		log.Fatal("err: ", err)
	}

	var primeDates []*date.Date
	for d := startDate; d.Next().Year() == year; d = d.Next() {
		if dateIsPrime(d) {
			primeDates = append(primeDates, d)
		}
	}
	return primeDates
}

// There are exactly 52 prime dates in a prime year.
func printPrimeDates(dates []*date.Date) {
	fmt.Println("But there are other prime dates in this year:")
	for i := 0; i < 13; i += 1 {
		for j := 0; j < 4; j += 1 {
			fmt.Printf("   %s  ", dates[j*13+i].ToIso())
		}
		fmt.Print("\n")
	}
}

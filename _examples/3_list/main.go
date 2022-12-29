package main

import (
	"fmt"
	"os"

	"github.com/lensesio/tableprinter"
)

type person struct {
	FirstName string
	LastName  string
}

func (p person) String() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	books := []string{
		"To Kill a Mockingbird (To Kill a Mockingbird) ",
		"The Hunger Games (The Hunger Games) ",
		"Harry Potter and the Order of the Phoenix (Harry Potter) ",
		"Pride and Prejudice ",
		"Animal Farm",
	}

	/*
	  BOOKS (5)
	 -----------------------------------------------------------
	  To Kill a Mockingbird (To Kill a Mockingbird)
	  The Hunger Games (The Hunger Games)
	  Harry Potter and the Order of the Phoenix (Harry Potter)
	  Pride and Prejudice
	  Animal Farm
	*/
	tableprinter.PrintHeadList(os.Stdout, books, "Books")

	println()

	numbers := []int{13213, 24554, 376575, 4321321321321, 5654654, 6654654, 787687, 8876876, 9321321}

	/*
	  NUMBERS (9)
	 --------------
	         13.2K
	         24.5K
	        376.5K
	          4.3T
	          5.6M
	          6.6M
	        787.6K
	          8.8M
	          9.3M
	*/
	tableprinter.PrintHeadList(os.Stdout, numbers, "Numbers")

	println()

	// DISCLAIMER: those are imaginary persons.
	persons := []person{
		{"Georgios", "Callas"},
		{"Ioannis", "Christou"},
		{"Dimitrios", "Dellis"},
		{"Nikolaos", "Doukas"},
	}

	/*
	  PERSONS (4)
	 ------------------
	  Georgios Callas
	  Ioannis Christou
	  Dimitrios Dellis
	  Nikolaos Doukas
	*/
	tableprinter.PrintHeadList(os.Stdout, persons, "Persons")

	println()

	/*
			HEADER      NUMBER   IS EVEN
		----------- -------- ---------
			String 1    1        false
			String 2    2        true
			String 3    3        false
			String 4    4        true
			String 5    5        false
			String 6    6        true
			String 7    7        false
			String 8    8        true
			String 9    9        false
			String 10   10       true
	*/

	rawMap := make([]map[string]any, 0, 10)
	for i := 1; i <= 10; i++ {
		entry := map[string]any{
			"header":  fmt.Sprintf("String %v", i),
			"number":  i,
			"is even": (i % 2) == 0,
		}
		rawMap = append(rawMap, entry)
	}
	prn := tableprinter.New(os.Stdout)
	prn.RowLengthTitle = func(int) bool { return false }
	prn.Print(rawMap)
}

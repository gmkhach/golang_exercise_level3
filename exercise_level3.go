package main

import "fmt"

func main() {
	/*
		Exercise 1
		Print every number from 1 to 10,000
	*/
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

	/*
		Exercise 2
		Print every rune code point of the uppercase alphabet three times. Your output should look like this:
		65
			U+0041 'A'
			U+0041 'A'
			U+0041 'A'
		66
			U+0042 'B'
			U+0042 'B'
			U+0042 'B'
	 	… through the rest of the alphabet characters
	*/
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	/*
		Exercise 3
		Create a for loop using this syntax: for condition { }
		Have it print out the years you have been alive.
	*/
	x := 1988
	for x <= 2020 {
		fmt.Println(x)
		x++
	}

	/*
		Exercise 4
		Create a for loop using this syntax: for { }
		Have it print out the years you have been alive.
	*/
	y := 1988
	for {
		if y > 2020 {
			break
		}
		fmt.Println(y)
		y++
	}

	/*
		Exercise 5
		Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
	*/
	for i := 11; i < 100; i++ {
		fmt.Println(i % 4)
	}

	/*
		Exercise 6
		Create a program that shows the “if statement” in action.
	*/
	kingSideRookIsUntouched := true
	kingIsUntouched := true
	thereAreNoPiecesInBetween := false
	if kingSideRookIsUntouched && kingIsUntouched && thereAreNoPiecesInBetween {
		fmt.Println("You can castle kingside")
	} else {
		fmt.Println("You cannot castle kingside")
	}
	/*
		Exercise 7
		Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
	*/
	queenSideRookIsUntouched := true
	if kingSideRookIsUntouched && kingIsUntouched && thereAreNoPiecesInBetween {
		fmt.Println("You can castle kingside")
	} else if queenSideRookIsUntouched && kingIsUntouched && thereAreNoPiecesInBetween {
		fmt.Println("You can castle queenside")
	} else {
		fmt.Println("You cannot castle")
	}

	/*
		Exercise 8
		Create a program that uses a switch statement with no switch expression specified.
	*/
	branch := "Army"
	switch {
	case branch == "Marine Corps":
		fmt.Println("Semper fi")
	case branch == "Army":
		fmt.Println("Hooah")
	case branch == "NAVY":
		fmt.Println("Hooyah")
	case branch == "Air Force":
		fmt.Println("Ouch (got a paper cut)")
	}

	/*
		Exercise 9
		Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
	*/
	favSport := "soccer"
	switch favSport {
	case "football":
		fmt.Println("Who cares?!")
	case "basketball":
		fmt.Println("hoops")
	case "tennis":
		fmt.Println("ping-pong")
	case "soccer":
		fmt.Println("Ole ole ole!")
	}
}
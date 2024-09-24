/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package main

import (
	"fmt" ;
	"os" ;
)

// Get user input
func main() {
	var numStudentNumbers int ;
	var numQueries int ;

	// Read in user input
	_, err := fmt.Scanf("%d %d", &numStudentNumbers, &numQueries) ;
	if err != nil {
		os.Exit(1) ;
	}

	// fmt.Println("numStudentNumbers:", numStudentNumbers) ;	// TEST
	// fmt.Println("numQueries:", numQueries) ;	// TEST

	
	// Read the student numbers from the same line
	studentNumbers := make([]int, numStudentNumbers) ;

	for i := 0; i < numStudentNumbers; i++ {
		_, err := fmt.Scan(&studentNumbers[i]) ;
		if err != nil {
			os.Exit(1) ;
		}
	}

	fmt.Println("studentNumbers:", studentNumbers) ;	// TEST




















	// // Read all student numbers into a single line
	// var input string ;
	// _, err = fmt.Scanln(&input)
	// if err != nil {
	// 	os.Exit(1) ;
	// }
	
	// fmt.Println(input)	// TEST
}
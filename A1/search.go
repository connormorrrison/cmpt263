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

// Function processes each query
func processQuery(query int, studentNumbers []int, numStudentNumbers int) {
	registered := false ;

	// Iterate through student numbers
	for i := 0; i < numStudentNumbers; i++ {
		if query == studentNumbers[i] {
			registered = true ;
			break ;
		}
	}

	if registered == true {
		fmt.Printf("Student %d is registered.\n", query) ;
	} else {
		fmt.Printf("Student %d is NOT registered.\n", query) ;
	}
}

// Main function processes user input
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

	studentNumbers := make([]int, numStudentNumbers) ;

	// Read in student numbers from the same line
	for i := 0; i < numStudentNumbers; i++ {
		_, err := fmt.Scan(&studentNumbers[i]) ;
		if err != nil {
			os.Exit(1) ;
		}
	}

	// fmt.Println("studentNumbers:", studentNumbers) ;	// TEST

	// Process each query
	for i := 0; i < numQueries; i++ {
		var query int ;
		_, err := fmt.Scanf("%d", &query) ;
		if err != nil {
			os.Exit(1) ;
		}

		// Compare query against recorded student numbers
		processQuery(query, studentNumbers, numStudentNumbers) ;
	}
}

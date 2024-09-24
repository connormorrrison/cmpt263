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

// Process each query
func processQuery(numStuNum int, numQuer int, stuNums int) {

}

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

	// Process each query
	for i := 0; i < numQueries; i++ {
		var query int ;
		_, err := fmt.Scanf("%d", &query) ;
		if err != nil {
			os.Exit(1)
		}

		// Offload process logic to processQuery function
		processQuery(numStudentNumbers, numQueries, studentNumbers)
}
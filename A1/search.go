/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package main

import (
	"fmt"
	"os"
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

}
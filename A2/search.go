/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package main

import (
	"fmt";
	"os";
	"a2/sort";
	"a2/sparse";
)

// Main function processes user input
//*@requires numValues, numQueries >= 0
func main() {
	var numValues int;
	var numQueries int;

	// Read in user input
	_, err := fmt.Scan(&numValues, &numQueries);
	if err != nil || numValues < 0 || numQueries < 0 {
		os.Exit(1);
	}

	// Initalize array for values
	values := make([]int, numValues);

	// Read in student numbers from the same line
	for i := 0; i < numValues; i++ {
		_, err := fmt.Scan(&values[i]);
		if err != nil || values[i] < 0 {
			os.Exit(1);
		}
	}

	// Sort student numbers using isort()
	sortedValues := sort.Isort(values);
	fmt.Printf("Sorted Set: %v\n", sortedValues);
	
	// Process each query
	for i := 0; i < numQueries; i++ {
		var query int;
		_, err := fmt.Scan(&query);
		if err != nil || query < 0 {
			os.Exit(1);
		}

		// Search the sorted array using find()
		found, position := sparse.Find(sortedValues, query);

		fmt.Printf("Query: %d -> ", query);
		if found {
			fmt.Printf("Found at index %d\n", position);
		} else {
			fmt.Println("Not found");
		}
	}
}
//*@ensures program exits with error code (1) if any input error is detected

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
	"a2/sparse"
)

// // Function processes each query
// func processQuery(query int, studentNumbers []int, numStudentNumbers int) {
// 	registered := false;

// 	// Iterate through student numbers
// 	for i := 0; i < numStudentNumbers; i++ {
// 		if query == studentNumbers[i] {
// 			registered = true;
// 			break;
// 		}
// 	}

// 	if registered == true {
// 		fmt.Printf("Student %d is registered.\n", query);
// 	} else {
// 		fmt.Printf("Student %d is NOT registered.\n", query);
// 	}
// }

// Main function processes user input
//*@requires numStudentNumbers, numQueries >= 0
func main() {
	// arr := []int{12, 5, 10, 6, 2};	// TEST
	// fmt.Println("Original array:", arr);	// TEST
	// 	sortedArr := sort.Isort(arr);	// TEST
	// 	fmt.Println("Sorted array:", sortedArr);	// TEST
	
	var numValues int;
	var numQueries int;

	// Read in user input
	_, err := fmt.Scanf("%d %d", &numValues, &numQueries);
	if err != nil || numValues < 0 || numQueries < 0 {
		os.Exit(1);
	}

	// fmt.Println("numStudentNumbers:", numStudentNumbers);	// TEST
	// fmt.Println("numQueries:", numQueries);	// TEST

	values := make([]int, numValues);

	// Read in student numbers from the same line
	for i := 0; i < numValues; i++ {
		_, err := fmt.Scan(&values[i]);
		if err != nil {
			os.Exit(1);
		}
	}

	// fmt.Println("studentNumbers:", studentNumbers);	// TEST

	// Sort student numbers using isort()
	sortedValues := sort.Isort(values);
	fmt.Println(sortedValues);
	
	// Process each query
	for i := 0; i < numQueries; i++ {
		var query int ;
		_, err := fmt.Scanf("%d", &query);
		if err != nil || query < 0 {
			os.Exit(1);
		}

		// Search the sorted array using find()
		found, position := sparse.Find(values, query)
		

		fmt.Println(found)
		fmt.Println(position)
	}
	//*@ensures program exits with error code (1) if any input error is detected
}

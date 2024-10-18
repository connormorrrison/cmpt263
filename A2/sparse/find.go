/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package sparse

// Find function locates the index of the query search value
func Find(arr []int, query int) (bool, int) {
	// Precondition: { arr is a sorted array of integers; query is an integer }
	// { arr is sorted in ascending order; query ∈ ℤ }
	
	firstOccurrence := -1;
	// { firstOccurrence == -1; arr is sorted; query ∈ ℤ }

	low := 0;
	// { low == 0; firstOccurrence == -1 }

	high := len(arr) - 1;
	// { high == len(arr) - 1; low == 0; firstOccurrence == -1 }

	// Loop Invariant I:
	// { 
    // - arr is sorted in ascending order;
    // - If query occurs in arr, it is within indices [low, high];
    // - firstOccurrence is the smallest index found so far where arr[firstOccurrence] == query, or -1 if not found yet
    // }
	
	for low <= high {
		// { I holds at the start of the loop; low <= high }
		
		mid := (low + high) / 2;
		// { mid == (low + high) / 2; low ≤ mid <= high; I holds }

		if arr[mid] == query {
			// { arr[mid] == query; I holds }
			firstOccurrence = mid;
			// { firstOccurrence == mid; arr[firstOccurrence] == query }

			high = mid - 1;	// Continue searching the left half to find first occurrence
			// { Updated high; search space reduced to [low, mid - 1]; I holds for next iteration }
		} else if arr[mid] < query {
			// { arr[mid] < query; I holds }
			low = mid + 1;
			// { Updated low; search space reduced to [mid + 1, high]; I holds for next iteration }
		} else {
			// { arr[mid] > query; I holds }
			high = mid - 1;
			// { Updated high; search space reduced to [low, mid - 1]; I holds for next iteration }
		}
		// { End of loop iteration; I holds }

	}
	// { Loop terminates when low > high; I holds }

	if firstOccurrence != -1 {
		// { query found at index firstOccurrence; arr[firstOccurrence] == query }
		return true, firstOccurrence;
		// Postcondition: { returns (true, firstOccurrence) where arr[firstOccurrence] == query }
	}
	
	return false, -1;
	// Postcondition: { returns (false, -1); query not found in arr }
}

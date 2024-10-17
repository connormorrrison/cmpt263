/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package sparse

func Find(arr []int, query int) (bool, int) {
	firstOccurrence := -1;

	low := 0;
	high := len(arr) - 1;

	// Perform binary search since array is sorted
	for low <= high {
		mid := (low + high) / 2;

		if arr[mid] == query {
			firstOccurrence = mid;
			high = mid - 1;	// Continue searching the left half to find first occurrence
		} else if arr[mid] < query {
			low = mid + 1;
		} else {
			high = mid - 1;
		}
	}

	if firstOccurrence != -1 {
		return true, firstOccurrence;
	}
	
	return false, -1;
}

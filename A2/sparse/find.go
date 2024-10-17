/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package sparse

func Find(arr []int, query int) (bool, int) {
	found := false;
	position := -1;

	low := 0;
	high := len(arr) - 1;

	// Perform binary search since array is sorted
	for low <= high {
		mid := (low + high) / 2;

		if arr[mid] == query {
			found = true;
			position = mid;
			break;
		} else if arr[mid] < query {
			low = mid + 1;
		} else {
			high = mid - 1;
		}
	}

	return found, position;
}

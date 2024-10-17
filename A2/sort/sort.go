/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package sort

func Isort(arr [] int) []int {
	// Outer loop starts at second element (i = 1)
	for i := 1; i < len(arr); i++ {
		
		// Set markers
		key := arr[i];	// Current value
		j := i - 1;	// Previous index

		// Inner loop moves elements greater than key to the right
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j];	// Shift element to the right
			j = j - 1;	// Move to previous element
		}

		// Insert key into position
		arr[j+1] = key;
	}

	return arr;
}

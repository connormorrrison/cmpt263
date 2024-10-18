/*
Name: Connor Morrison
NSID: tvi340
Student Number: 11374770
Course: CMPT 263
*/

package sort

// Isort function performs insertion sort
//*@requires arr is an array of integers
func Isort(arr [] int) []int {
	// Pre-condition: { arr is an array of integers of length n }
	// { arr is an array of integers of length n }

	for i := 1; i < len(arr); i++ {
		// Loop invariant I1: { arr[0..i-1] is sorted in ascending order }

		// { I1 holds at the start of the loop }
		key := arr[i];
		// { key == arr[i]; I1 holds }

		j := i - 1;
		// { j == i - 1; key == arr[i]; I1 holds }

		// Loop invariant I2:
		// { For all k in [j+1, i], arr[k] > key; arr[0..j] and arr[j+1..i] 
		// concatenated will maintain the order of arr[0..i-1] }
		
		for j >= 0 && arr[j] > key {
			// { j >= 0; arr[j] > key; I2 holds }
			arr[j+1] = arr[j];
			// { arr[j+1] == arr[j]; elements at positions j+1 to i shifted right by one; I2 holds }

			j = j - 1;
			// { j decreased by 1; I2 holds }
		}
		// { arr[0..j] <= key <= arr[j+1..i-1]; I2 holds }

		arr[j+1] = key;
		// { arr[0..i] is sorted in ascending order; I1 holds for next iteration }

		// End of loop iteration
		// { I1 holds for next i }
	}
	// Postcondition: { arr[0..n-1] is sorted in ascending order }

	return arr;
}
//*@ensures returns arr sorted in ascending order

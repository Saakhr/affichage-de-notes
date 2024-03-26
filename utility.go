package main

import "strconv"

func removeEmptyStrings(slice []string) []string {
	// Create a new slice to store non-empty strings
	result := make([]string, 0)

	// Iterate over the original slice
	for _, str := range slice {
		// Check if the string is not empty
		if str != "" {
			// Add non-empty strings to the result slice
			result = append(result, str)
		}
	}

	return result
}
func stringsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i, str := range strs {
		// Convert each string to an integer
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err // return error if conversion fails
		}
		// Append the integer to the new array
		ints[i] = num
	}
	return ints, nil
}

//Author: David Brungardt
//Date: 1/18/2020
// Merge sort algorithm
// Time complexity: O(nlog(n))

package main

import (
	"fmt"
) // end of import

func main() {

	// slice we will be passing into mergeSort algorithm
	mySlice := []int{5, 3, 8, 1, 6, 10, 7}
	// print statement to see original content of Slice
	fmt.Println(mySlice)

	// print statement to print final output
	fmt.Println(mergeSort(mySlice))

} // end of main

func mergeSort(param []int) []int {

	// assign parameter to variable in mergeSort method
	inputSlice := param
	sliceLen := len(inputSlice)

	// if slice length is one, return slice
	if sliceLen == 1 {
		return inputSlice
	} else {

		// pass slice into sort method
		outSlice := split(inputSlice)

		// return sorted slice
		return outSlice
	}

} // end of mergeSort method

func split(param []int) []int {

	// assign parameter to inputSlice
	inputSlice := param
	// get slice length
	sliceLength := len(inputSlice)
	// if slice length is 1, return slice
	if sliceLength == 1 {
		return inputSlice
	} else {

		// split slice into two slices
		firstSlice := inputSlice[(sliceLength / 2):]
		secondSlice := inputSlice[:(sliceLength / 2)]

		// return merged list
		// use split recursively to arrive at base case
		// then sort as values are merged
		return merge(split(firstSlice), split(secondSlice))
	}

} // end of split method

func merge(firstParam []int, secondParam []int) []int {

	// assign parameters to firstSlice and secondSlice
	firstSlice := firstParam
	secondSlice := secondParam
	// get length of each slice
	firstLen := len(firstSlice)
	secondLen := len(secondSlice)
	// set counters
	i := 0
	j := 0
	// create new slice for sorted list
	var outSlice []int

	// loop through slices comparing values
	// combine slices into sorted list (slice)
	for (i < firstLen) && (j < secondLen) {

		// if value in each slice is equal
		if firstSlice[i] == secondSlice[j] {
			// append both to output list
			outSlice = append(outSlice, firstSlice[i], secondSlice[j])
			// increment both counters
			i++
			j++

			// if value in firstSlice is less than value in secondSlice
		} else if firstSlice[i] < secondSlice[j] {
			// append firstSlice value to outSlice
			outSlice = append(outSlice, firstSlice[i])
			// increment counter
			i++

		} else if firstSlice[i] > secondSlice[j] {
			// append secondSlice value to outSlice
			outSlice = append(outSlice, secondSlice[j])
			// increment counter
			j++
		}

	} // end of loop

	// if there are remaining values
	// append to outSlice
	for i <= (firstLen - 1) { // if i counter is less than j
		// append remaining firstSlice value to list
		outSlice = append(outSlice, firstSlice[i])
		i++
	} // end of loop

	for j <= (secondLen - 1) { // if j counter is less than i
		// append remining secondSlice value to list
		outSlice = append(outSlice, secondSlice[j])
		j++
	} // end of loop

	// return sorted slice
	return outSlice

} // end of merge method

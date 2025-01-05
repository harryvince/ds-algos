package arrays

import "fmt"


func BubbleSort() {
    println("--------- BUBBLE SORT --------")
	array := []int{7, 12, 9, 11, 3}
	fmt.Printf("Initial Array: %v\n", array)
	result := bubbleSort(array)
	fmt.Printf("Sorted Array: %v\n", result)
    println("----- END OF BUBBLE SORT -----")
}

func bubbleSort(values []int) []int {
	length := len(values) - 1

	x := 0
	noSwap := 0
	for {
		current := values[x]
		next := values[x+1]

		if current > next {
			values[x+1] = current
			values[x] = next
		} else {
			noSwap++
		}

		x++

		if noSwap == length {
			break
		}

		if x == length {
			x = 0
			noSwap = 0
		}
	}

	return values
}

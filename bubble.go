package main

// O(n^2)
// heuristic algorithms
func bubbleSort(numbers []int) {
	for j := len(numbers) - 1; j > 0; j-- { //n
		hasSwapped := false
		for i := 0; i < j; i++ { //n
			if numbers[i] > numbers[i+1] {
				swap := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = swap
				hasSwapped = true
			}
		}
		//fmt.Printf("step %d: %+v\n", len(numbers)-j, numbers)
		if !hasSwapped {
			return
		}
	}
}

package main

import "fmt"

func main() {
	numbers := []int{100, 5, 4, 3, 10, 50, 20, 15}

	//1) pivot: 15
	//left: 5,4,3,10
	//right: 100,50,20
	//out: 5,4,3,10,|15|,10,50,20
	//2a) [5,4,3,10] pivot: 10
	//right: nil
	//left [5,4,3]

	//2b) [10,50,20]

	//3a) [5,4,3] pivot:3
	//left: nil
	//right: [5,4]

	//numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("before: %+v\n", numbers)
	quickSort(numbers)
	fmt.Printf("after: %+v\n", numbers)
	//before: 5,4,3,10,50,20,1
	//after:  4,3,5,10,20,1,50
}

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

// divide & conquer
// best case: O(nlogn)
// worst case: O(n^2)
func quickSort(numbers []int) {
	quickSortStep(numbers, 0, len(numbers)-1)
}

func quickSortStep(numbers []int, low, high int) {
	if low < high {
		pi := partition(numbers, low, high) //pi = pivot index
		//left [5(0),4,3,10(3)]
		//pi: 4
		//right [100(5),50,20(7)]
		quickSortStep(numbers, low, pi-1)
		quickSortStep(numbers, pi+1, high)
	}
}

// 1. set pivot
// 2. divide tables to left and right
// 3. return pivot index
func partition(arr []int, low, high int) int {
	//[5(0),4,3,10(3)]
	//[100, 5, 4, 3, 10, 50, 20, 15]
	//pivot: 15
	//100, 5, 4, 3, 10, 50, 20
	//5,100, 4, 3, 10, 50, 20
	//low: 0
	//high: 4
	pivot := arr[high]
	last := low //0

	for i := low; i < high; i++ {
		if arr[i] < pivot { //1 5<->100
			//swap := arr[last]
			//arr[last] = arr[i] //100<--5
			//arr[i] = swap
			//syntactic sugar
			arr[last], arr[i] = arr[i], arr[last]
			fmt.Printf("p=%d,l=%d,h=%d, arr=%+v\n", pivot, low, high, arr[low:high])
			last++
		}
	}

	arr[last], arr[high] = arr[high], arr[last]
w	fmt.Printf("---> p=%d,l=%d,h=%d, arr=%+v\n", pivot, low, high, arr[low:high])
	fmt.Printf("last: %d\n", last)
	return last
}

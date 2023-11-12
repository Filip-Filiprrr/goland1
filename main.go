package main

import "fmt"

func main() {
	//numbers := []int{100, 5, 4, 3, 10, 50, 20, 15}

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
	//fmt.Printf("before: %+v\n", numbers)
	//quickSort(numbers)
	//fmt.Printf("after: %+v\n", numbers)
	//before: 5,4,3,10,50,20,1
	//after:  4,3,5,10,20,1,50
	l := List[int]{}
	//a := []int{5, 10, 15}
	//l := List[float32]{} //float
	l.Add(5)    //idx: 0
	l.Add(10)   //idx: 1
	l.Add(15)   //idx: 2
	l.Add(20)   //idx: 3
	l.Delete(3) //skasuj idx:1 czyli wartosc 10

	println("=============== head")
	i := 0
	//O(n)
	for e := l.Head(); e != nil; e = e.Next() {
		fmt.Printf("list[%d]: %+v\n", i, *e)
		i++
	}
	println("=============== tail")
	//TODO
	i = 0
	for e := l.Tail(); e != nil; e = e.Prev() {
		fmt.Printf("list[%d]: %+v\n", i, *e)
		i++
	}
	//set O(1)
	//set.Get(idx=2)
	//l.Delete(2) //skaskuj 2 element listy
	//for _, n:=range l {
	//
	//}
	//l := Node[float64]{} //double
	//l.DoSomething()
	//fmt.Printf("%+v\n", l)
}

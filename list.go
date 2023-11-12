package main

import "fmt"

// TODO zrobic ta liste 2 kierunkowa
type List[T any] struct {
	head *Node[T] //wskaznik w pamiecy na 1 element listy
	tail *Node[T] //wskaznik w pamiecy na ostatni element listy
}

type Node[T any] struct {
	data T        //wartosc elementu listy
	next *Node[T] //wskaznik w pamiecy na strukture
	prev *Node[T] //prev aby miec liste 2-kierunkowa
}

//ta f-cja zrobi kopie Node i pracujesz na kopii
//func (n Node[T]) DoSomething() {
//	//empty
//	n.data = nil
//}
//
////ta f-cje operuje na obecnym/oryginalnym Node
//func (n *Node[T]) DoSomething1() {
//	//empty
//	n.data = nil
//}

func (l *List[T]) Add(value T) {
	//do listy dopiac Node: dane + wskaznik na nastepny element
	//alokacja nowego elementu w pamieci
	//C#: new Node[T]
	n := Node[T]{ //alokacja pamieci
		data: value,
	}
	//przypiac nowy Node (n) do listy
	//l.tail =
	//po dodaniu 5
	//l.head=5 l.tail=5
	//dodaj
	l.tail = &n
	if l.head == nil {
		//dodajemy 5
		l.head = &n //na rysunku lewy nody (pozniej masz srodkwy i po prawej)
	} else {
		last := l.head // 5
		for {
			if last.next == nil { //nie ma alokacji pamieci
				last.next = &n //& - przypnij miejsce w pamieci (adres w pamieci) do naszej listy
				n.prev = last
				break
			} else { //jest alokacja pamieci, wiec przeszkocz do nastepnego elementu
				last = last.next
			}
		}
	}
}

func (l *List[T]) Delete(idx int) {
	//zadanie: delete
	//1. napisac for, zeby przejsc po elementach listy
	//2. zliczac elementy e jesli idx elementu == idx to wtedy kasujemy
	//3. odpiac element od listy
	//	l.Add(5)    //idx: 0
	//	l.Add(10)   //idx: 1
	//	l.Add(15)   //idx: 2
	//	l.Delete(1) //skakuj idx:1 czyli wartosc 10
	//head(5)->node(10)->node/tail(15)
	//head(5)->node/tail(15)
	i := 0
	var prev *Node[T] //nil
	//e=nil
	for e := l.head; e != nil; e = e.Next() {
		//e to node listy, para{data,next}

		fmt.Printf("element[%d]: %+v\n", i, e)
		if idx == i {
			//e=node(15,next)
			fmt.Printf("kasuje %d: %+v\n", i, e)
			if idx == 0 {
				l.head = e.Next()
				l.head.prev = nil
			} else {
				//prev:15
				//e:20
				//e.Next():nil
				//20.prev=
				//15.prev=5
				prev.next = e.Next() // 15.next=nil : zmieniasz dowiazanie next/do przodu
				if e.Next() == nil { //jesli e.next=nil to mamy tail
					l.tail = prev
				} else {
					e.Next().prev = prev //20.next.prev=15
				}
				//c++ delete(e)
			}
		}
		i++
		prev = e
	}
	//e=nil
}

func (l *List[T]) Head() *Node[T] {
	return l.head
}

func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

func (l *Node[T]) Next() *Node[T] {
	return l.next
}

func (l *Node[T]) Prev() *Node[T] {
	return l.prev
}

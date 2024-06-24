package main

import (
	"fmt"
	"sort"
	"strings"
)

type DoublyLinkedNode struct{
	data int
	prev interface
	next interface 	
}

func First(List []DoublyLinkedNode) (DoublyLinkedNode, int) {
	i :=0
	for  el:= range List {
		if el.prev == 0{
			return el, i
		}
		i++
	}
	return nil, nil		  
}

func Last(List []DoublyLinkedNode) (DoublyLinkedNode, int) {
	i :=0
	for  el:= range List {
		if el.next == 0{
			return el, i
		}
		i++
	}
	return nil, nil
		  
}

func Len(List []DoublyLinkedNode) int {

	Len := 0

	for  el:= range List {
		Len++
	  }
	return Len
		  
}

func PushFront(List []DoublyLinkedNode, data int, ) {

	FirstNode, First_i := First(List);
	Len_List := Len(List)
	Node := DoublyLinkedNode{
		data,
		0,
		First_i,	
	}
	List = append(List, Node)
	
	FirstNode.prev = Len_List
	  
}

func PushBack(List []DoublyLinkedNode, data int, ) {

	LastNode, Last_i := Last(List);
	
	Len_List := Len(List)
	Node := DoublyLinkedNode{
		data,
		Last_i,
		0 ,	
	}
	List = append(List, Node)
	
	LastNode.next = Len_List
	  
}

func Remove(List []DoublyLinkedNode, i int, ) {
	nom:=0
	for  el:= range List {
		if nom>=  {
			List[i-1]=el	
		}
		nom++
	}
}


func main() {
	
	var DoublyLinkedList  []DoublyLinkedNode


	   




}
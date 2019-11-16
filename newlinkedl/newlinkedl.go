package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
	next *person
}

func printList(personList *person) {
	for p := personList; p != nil; p = p.next {
		fmt.Println(p)
	}
}

/* All insertion scenarios for linked list ... in the front, back, and random postion in middle */
func insertAtFront(inputPerson *person, personList *person) *person {
	// LL doesnt exist
	if personList == nil {
		return inputPerson
	}

	// LL does exist, insert at front
	inputPerson.next = personList
	return inputPerson
}

/* Insert at back of the list */
func insertAtBack(inputPerson *person, personList *person) *person {
	if personList == nil {
		return personList // you return personList because that is the head of the list in this scenario because you are adding to the back
	}
	for i := personList; i != nil; i = i.next {
		if i.next == nil {
			i.next = inputPerson
			return personList
		}
	}
	return personList
}

/* Insert the node at location k */
func insertAtK(inputPerson *person, personList *person, k int) *person {
	count := 0 // counter to track what position you are in the LL
	if k == 0 {
		return insertAtFront(inputPerson, personList)
	}
	for i := personList; i != nil; i = i.next {
		count++
		if k == count {
			inputPerson.next = i.next
			i.next = inputPerson
			return personList
		}
	}
	return personList
}

/* Check if a person with the given age is in the linked list */
func findValueK(personList *person, k int) *person {
	for i := personList; i != nil; i = i.next {
		if k == i.age {
			return i
		}
	}
	return nil
}

/* Delete the person at the front of the LL */
func deleteAtFront(personList *person) *person {
	// ll does not exist
	if personList == nil {
		return nil
	}
	// set head to the next node in the ll
	personList = personList.next
	return personList
}

/* Delete the person at the end of the list */
func deleteAtEnd(personList *person) *person {
	// ll does not exist or has only one element
	if personList == nil || personList.next == nil {
		return nil
	}
	// one fast pointer that guarantees we stop at 2nd to last node
	fast := personList
	for {
		if fast.next.next == nil {
			// always stop at second to last pointer
			// set next to nil because you are at second to last pointer which deletes the node
			fast.next = nil
			return personList
		}
		fast = fast.next
	}
}

/* Delete the person at position k */
func deleteAtK(personList *person, k int) *person {
	// ll is empty
	if personList == nil {
		return nil
	}

	// if you want to delete node at index 0
	if k == 0 {
		// if removing position 0 then set head to the node next
		personList = personList.next
		return personList
	}

	// track where you are in list
	count := 0

	// any other postion
	for i := personList; i.next != nil; i = i.next {
		// if k - 1 (one node behind) matches count that means we are the correct index, then remove node at i
		if k-1 == count {
			// since one node behind set this nodes next value to be next.next
			i.next = i.next.next
			return personList
		}

		count++
	}

	return personList
}

func reverseLinkedList(personList *person) *person {
	var prev *person
	head := personList

	for head != nil {
		temp := head.next
		head.next = prev
		prev = head
		head = temp
	}
	return prev
}

func main() {
	billy := &person{34, "billy", nil}
	bob := &person{45, "bob", nil}
	bilbo := &person{56, "bilbo", nil}
	brian := &person{69, "brian", nil}
	boris := &person{388, "boris", nil}
	bosch := &person{12334, "bosch", nil}
	bessie := &person{1, "bessie", nil}

	personList := billy
	fmt.Println("----------Inserted at Front----------------")
	personList = insertAtFront(bob, personList)
	printList(personList)
	fmt.Println("----------Inserted at Front----------------")
	personList = insertAtFront(bilbo, personList)
	printList(personList)
	fmt.Println("---------Inserted at Back-----------------")
	personList = insertAtBack(brian, personList)
	printList(personList)
	fmt.Println("-------Inserted at Back-------------------")
	personList = insertAtBack(boris, personList)
	printList(personList)
	fmt.Println("-------Inserted at index 2-------------------")
	personList = insertAtK(bosch, personList, 2)
	printList(personList)
	fmt.Println("-------Inserted at index 3-------------------")
	personList = insertAtK(bessie, personList, 3)
	printList(personList)
	fmt.Println("-------Find person that is 12334 years old-------------------")
	fmt.Println(findValueK(personList, 12334))
	fmt.Println("-------Deleting Node at the end of the list-------------------")
	personList = deleteAtEnd(personList)
	printList(personList)
	fmt.Println("-------Deleting Node at the front of the list-------------------")
	personList = deleteAtFront(personList)
	printList(personList)
	fmt.Println("-------Deleting Node at index 2-------------------")
	personList = deleteAtK(personList, 5)
	printList(personList)
	fmt.Println("-------Current List reversed-------------------")
	personList = reverseLinkedList(personList)
	printList(personList)

}

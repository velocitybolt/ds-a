package main

import "fmt"

type person struct {
	age       int
	firstName string
	next      *person
}

func printList(personList *person) {
	for p := personList; p != nil; p = p.next {
		fmt.Println(p)
	}
}

func addToBack(newPerson, personList *person) *person {
	if personList == nil {
		return personList
	}
	for p := personList; p != nil; p = p.next {
		if p.next == nil {
			p.next = newPerson
			return personList
		}
	}
	return personList
}

func findCycle(personList *person) bool {
	slow := personList
	fast := personList

	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

func isPalindrome(personList *person) bool {
	slow := personList // O(2n) runtime, O(1) space
	fast := personList
	var prev *person
	var midpoint *person
	// midpoint of the list will always be when fast is at the end of the list
	// because it is moving twice as fast
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		midpoint = slow.next
		slow.next = prev
		prev = slow
		slow = midpoint
	}

	if midpoint != nil && fast != nil {
		midpoint = midpoint.next
	}

	for prev != nil && midpoint != nil {
		if prev.firstName != midpoint.firstName {
			return false
		}
		prev = prev.next
		midpoint = midpoint.next

	}

	return true
}

func findCycleWithSet(personList *person) bool {
	set := make(map[*person]struct{})

	for ; personList != nil; personList = personList.next {
		if _, ok := set[personList]; ok {
			return true
		}

		set[personList] = struct{}{}
	}

	return false
}

func (personList *person) reverse() *person {
	// initalize prev node of linked list
	var prev *person
	// initialize head node of linked list
	head := personList
	// iterate until there is no person left
	for head != nil {
		// store next in temp variable so that it can be set head node later
		temp := head.next
		// break the link between head and next and have head point to prev
		head.next = prev
		// set previous node to be head node because it needs to point in other direction
		prev = head
		// set head node to temp because that is now the beginning of the reversed linked list
		// this is because temp was head.next which was set to prev before so now temp which is head.next at the moment is now head
		head = temp
	}
	return prev

}

func main() {
	billy := &person{34, "billy", nil}
	bob := &person{45, "bob", nil}
	bilbo := &person{56, "bilbo", nil}

	personList := billy
	addToBack(bob, personList)
	addToBack(bilbo, personList)

	fmt.Println("list before reversing:")
	printList(personList)

	personList = personList.reverse()

	fmt.Println("list after reversing:")
	printList(personList)

	fmt.Println("findCycle before addition:", findCycle(personList))
	fmt.Println("findCycleWithSet before addition:", findCycleWithSet(personList))

	addToBack(bob, personList)

	fmt.Println("findCycle after addition:", findCycle(personList))
	fmt.Println("findCycleWithSet after addition:", findCycleWithSet(personList))

	nodeOne := &person{23, "r", nil}
	nodeTwo := &person{27, "o", nil}
	nodeThree := &person{34, "o", nil}
	nodeFour := &person{45, "r", nil}

	palinList := nodeOne
	addToBack(nodeTwo, palinList)
	addToBack(nodeThree, palinList)
	addToBack(nodeFour, palinList)

	fmt.Println("Does this linked list contain a palindrome?")
	fmt.Println(isPalindrome(palinList))
}

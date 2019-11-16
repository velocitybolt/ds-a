package main

import (
	"fmt"
	"math/rand"
	"time"
)

type car struct {
	value int
	name  string
	next  *car
}

func printList(personList *car) {
	for p := personList; p != nil; p = p.next {
		fmt.Println(p)
	}
}

func addToBack(newPerson, personList *car) *car {
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

/* This method returns the price of a random car within the linked list */
func (carList *car) getRandom() int {
	length := 0
	for car := carList; car != nil; car = car.next {
		length++
	}

	target := rand.Intn(length)
	car := carList
	for i := 0; i < target; i++ {
		car = car.next
	}

	return car.value
}

func main() {
	ferrari := &car{340000, "ferrari", nil}
	toyota := &car{31000, "toyota", nil}
	bmw := &car{56000, "bmw", nil}

	carList := ferrari
	addToBack(toyota, carList)
	addToBack(bmw, carList)

	fmt.Println("This is the list")
	printList(carList)

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Here is the price of a random car from the list")
	price := carList.getRandom()
	fmt.Println(price)

}

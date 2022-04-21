package main

import "fmt"

type speak interface {
	talk()
}

type person1 struct {
	name string

	age int
}

type person2 struct {
	name string

	age int
}

func boo(b speak) {

	fmt.Println("Hi !! I am  : ", b)

}

func (k person2) talk() {

	fmt.Printf("My name is %s \n ", k.name)

	fmt.Printf("My age is %d \n", k.age)

}

func (t person1) talk() {

	fmt.Printf("My name is %s \n ", t.name)

	fmt.Printf("My age is %d \n", t.age)
}

func main() {

	rahul := person1{"Rahul", 21}

	ritwik := person2{"Ritwik", 22}

	rahul.talk()

	ritwik.talk()

	boo(rahul)

	boo(ritwik)

}

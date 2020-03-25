package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func f() {
	fmt.Println("f function")
}
func strlen(s string, c chan int) {
	c <- len(s)
}

type Person struct {
	Name string
	Age  string
}
type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}
func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

type Foo struct {
	Bar string
	Baz string
}

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}
func main() {
	var x = "Hello"
	z := int(42)
	// Slice and Maps
	var s = make([]string, 0)
	var m = make(map[string]string)
	s = append(s, "some string")
	m["skey"] = "svalue"
	fmt.Println("")
	// Pointers, Structs, and Interfaces
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(count)
	guy := new(Person)
	guy.Name = "Tralah Brian"
	guy.SayHello()
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)
	if z%2 == 0 {
		fmt.Println("%d is even", z)
	} else {
		fmt.Println("%d is odd", z)
	}
	switch x {
	case "Hello":
		fmt.Println("case Hello")
	case "Other":
		fmt.Println("case Other")
	default:
		fmt.Println("Default case")
	}
	nums := []int{2, 3, 4, 5}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
	// goroutines lightweight threads by prefixing go to the function call
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
	// Channels
	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	a, b := <-c, <-c
	fmt.Println(a, b, a+b)
	// Handling Structured Data
	f := Foo{"Joe Junior", "Tralah M"}
	bad, _ := json.Marshal(f)
	fmt.Println(string(bad))
	json.Unmarshal(bad, &f)
}

package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) sayHello() {
	fmt.Printf("Hello, I'm %v, %v year(s) old\n", p.name, p.age)
}

func main(){
	p := &Person{
		name: "apocelipes",
		age: 100,
	}
	p.sayHello()
}

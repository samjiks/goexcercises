package main

import (
	"fmt"
	"strconv"
)


func (j Java) speak() string {
	c := fmt.Sprint(j.Name + " whose age is " +  strconv.Itoa(j.Age) + " says "  + j.Says)
	return c
}

func (g Go) speak() string  {
	c := fmt.Sprint(g.Name + " whose age is " + strconv.Itoa(g.Age) + " says "  + g.Says)
	return c
}

func (p Php) speak() string {
	c := fmt.Sprint(p.Name + " whose age is " +  strconv.Itoa(p.Age) + " shouts "  + p.Shouts)
	return c
}

func (c C) speak() string {
	s := fmt.Sprint(c.Name + " whose age is " +  strconv.Itoa(c.Age) + " explains "  + c.Explains)
	return s
}

type Developer interface{
	speak() string
}

type Developers []Developer

type Java struct {
	Name string
	Age  int
	Says string
}

type Go struct {
	Name string
	Age  int
	Says string
}

type Php struct {
	Name   string
	Age    int
	Shouts string
}

type C struct {
	Name     string
	Age      int
	Explains string
}

func main() {
	p := Developers{
		Php{Name: "connor", Age: 22, Shouts: "I like things slow and loosely typed"},
		Go{Name: "charlie", Age: 24, Says: "Throw me a Gopher"},
		Java{Name: "alex", Age: 25, Says: "I have a factory full of problems"},
		Go{Name: "bill", Age: 55, Says: "Dont talk to me about dependency management"},
		C{Name: "Jill", Age: 45, Explains: "I wish my life was simpler"},
		C{Name: "Ahamed", Age: 33, Explains: "Im tradtional and so is my language"},
		Java{Name: "Rob", Age: 55, Says: "Your only using up 99% of your resources..."},
		Php{Name: "Kimbley", Age: 23, Shouts: "Tabs and spaces lets spend loads of time debating it..."},
	}

	for _, developers := range p {
		fmt.Println(developers.speak())
	}
}

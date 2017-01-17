package main

import (
	"fmt"
	"flag"
	"strings"
)
var task int

func init() {
	flag.IntVar(&task, "task", 0,
		"Task [1 - Person Working on 5th Floor," +
		" 2 - Person Working under Charlie, " +
		"3 - Person who work in Product]")
}

func main() {

	flag.Parse()
	defer func(args int) {
		switch args {
		case 1:
			task1()
		case 2:
			task2()
		case 3:
			task3()
		default:
			task1()
			task2()
			task3()
		}
	}(task)
}


func task1(){
	fmt.Println("Person Working on 5th Floor")
	for name, detail := range People {
		if detail.floor == 5 && strings.ToLower(detail.dept) == "development" {
			fmt.Printf("Name: %s Dept: %s\n", strings.Title(name), detail.dept)
		}
	}
}

func task2(){
	fmt.Println("Person Working under Charlie")
	for name, detail := range People {
		if strings.ToLower(detail.manager) == "charlie" {
			fmt.Printf("Name: %s Dept: %s\n", strings.Title(name), detail.dept)
		}
	}
}

func task3(){
	fmt.Println("Person who work in Product")
	for name, detail := range People {
		if strings.ToLower(detail.dept) == "product" {
			fmt.Printf("Name: %s Dept: %s\n", strings.Title(name), detail.dept)
		}
	}
}

func people() []details {
	people_details := []details{}

	for _, detail := range People {
		people_details = append(people_details, detail)
	}
	return people_details
}

type details struct {
	age     int
	dept    string
	floor   int
	manager string
}

var People = map[string]details{

	"charlie": details{
		age:     22,
		dept:    "development",
		manager: "Fraser",
		floor:   5,
	},
	"john": details{
		age:     34,
		dept:    "Product",
		manager: "Chris",
		floor:   3,
	},
	"bob": details{
		age:     24,
		dept:    "Development",
		manager: "Charlie",
		floor:   1,
	},
	"chris": details{
		age:     34,
		dept:    "Development",
		manager: "Fraser",
		floor:   3,
	},
	"nadim": details{
		age:     44,
		dept:    "Development",
		manager: "Fraser",
		floor:   3,
	},
	"josh": details{
		age:     25,
		dept:    "Development",
		manager: "Jai",
		floor:   3,
	},
	"micheal": details{
		age:     28,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"alex": details{
		age:     32,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"andrew": details{
		age:     44,
		dept:    "Development",
		manager: "Charlie",
		floor:   3,
	},
}

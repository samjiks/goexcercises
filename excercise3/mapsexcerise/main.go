package main

import (
	"github.com/samjiks/excercise3/mapsexcerise/people"
	"fmt"
	"time"
)

func main() {
	c := make(chan []string)
        timeout := time.After(time.Second)

	go people.CharliesTeam(c)
	go people.PeopleOnFloorFive(c)
	go people.PeopleWorkingInProduct(c)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Ending Maps Excercise Program..")
			return
		}
	}

}
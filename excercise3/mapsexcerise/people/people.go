package people

import "strings"

func CharliesTeam(c chan []string) {
	charliesteam := make([]string, 0)
	for name, dets := range People {
		if strings.ToLower(dets.manager) == "charlie" {
			charliesteam = append(charliesteam, name)
		}
	}
	c <- charliesteam
}

func PeopleOnFloorFive(c chan []string) {
	fifthfloor := make([]string, 0)
	for name, dets := range People {
		if dets.floor == 5 {
			fifthfloor = append(fifthfloor, name)
		}
	}
	c <- fifthfloor
}

func PeopleWorkingInProduct(c chan []string) {
	productowners := make([]string, 0)
	for name, dets := range People {
		if strings.ToLower(dets.dept) == "product" {
			productowners = append(productowners, name)
		}
	}
	c <- productowners
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
		floor:   5,
	},
	"bob": details{
		age:     24,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
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
		floor:   5,
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

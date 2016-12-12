package main

import (
	"testing"
	"reflect"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}, message interface{}){
	if expected != actual {
		t.Errorf("got %v, Expected %v, Message %v",  expected, actual, message)
	}
}

func assertNotEqual(t *testing.T, actual interface{}, expected interface{}, message interface{}){
	if expected == actual {
		t.Errorf("got %v, expected %v, message: %v",  expected, actual, message)
	}
}

func assertTrue(t *testing.T, true bool, message interface{}){
	if !true {
		t.Errorf("should be true, got %v, message: %v", true, message)
	}
}

func assertIn(t *testing.T, v interface{}, l interface{}, message interface{}) {

	for _, j := range l.([]string) {
	   if v == j {
		   return
	   }
	}
	t.Errorf("%v not found in %v, message: %v", v, l, message)
}

func assertNotIn(t *testing.T, v interface{}, l interface{}, message string) {

	for _, j := range l.([]string) {
	   if v == j {
		   t.Errorf("%v found in %v, message: %v", v, l, message)
		   return
	   }
	}
}

func assertIsLower(t *testing.T, s string, message string) {

	for i := range s {
	   c := s[i]
	   if c >= 'A' && c <= 'Z' {
		t.Errorf("%v is not lower case, message: %v", s,  message)
	   	return
	   }
	}
}

func assertIsNotEmpty(t *testing.T, s interface{}, message string) {

	if len(s.([]string)) >= 0 {
		t.Errorf("%v is not empty, message: %v", s,  message)
	   	return
	}
}

func assertIsEmpty(t *testing.T, s interface{}, message string) {

	if len(s.([]string)) == 0 {
		t.Errorf("%v is empty, message: %v", s,  message)
	   	return
	}
}


func TestCharliesTeamLengthIs4(t *testing.T) {
	team := charliesTeam()
	assertEqual(t, len(team), 4, "Should be 4")
	assertNotEqual(t, len(team), 3, "Should be 4")
	assertNotEqual(t, len(team), 1, "Should be 4")
}

func TestCharliesTeamReturnAsSlice(t *testing.T) {
	team := charliesTeam()
	assertTrue(t, reflect.TypeOf(team).Kind() == reflect.Slice, "Not True")
}

func TestPeopleWorkingOnFifthFloor(t *testing.T) {
	people := peopleOnFloorFive()
	cases := []struct{Actual string}{
		{"charlie"},
		{"micheal"},
		{"bob"},
		{"josh"},
		{"alex"},
		{"john"},
	}

	for _, j := range cases {
		assertIn(t, j.Actual, people, "Cannot be found")
	}
}

func TestPeopleNotWorkingonFifthFloor(t *testing.T){
	people := peopleOnFloorFive()

	casesNotIn := []struct{Actual string}{
		{"chris"},
		{"nadim"},
		{"abel"},
	}

	for _, j := range casesNotIn {
		assertNotIn(t, j.Actual, people, "Found")
	}
}

func TestPeopleWorkingInProduct(t *testing.T){
	people := peopleWorkingInProduct()

	for _, j := range people {
		assertIsLower(t, j, "Not Lower Case")
	}
}


func TestPeopleWorkingInProductIsEmpty(t *testing.T){
	people := peopleWorkingInProduct()

	assertIsNotEmpty(t, people, "Is Not Empty")
}

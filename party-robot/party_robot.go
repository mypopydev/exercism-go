package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "!" + " You are now " + fmt.Sprintf("%d", age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	s := "Welcome to my party, " + name + "!\n"
	s += fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n",
		table, direction, distance)
	return s + "You will be sitting next to " + neighbor + "."
}

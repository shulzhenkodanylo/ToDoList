package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanning() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}

func removeElement(slice *[]string, i int) []string {
	if i < 0 || i >= len(*slice) {
		return *slice
	}
	return append((*slice)[:i], (*slice)[i+1:]...)
}

func userInput() string {
	fmt.Print("What you want to do [a]dd. [r]emove. [done]. [e]dit. [q]uit: ")
	input := scanning()
	return input
}

func userChoice(choice string, tasks *[]string) {
	switch strings.ToLower(choice) {
	case "a":
		addToList(tasks)
	case "r":
		removeFromList(tasks)
	case "d":

	case "e":

	case "q":
	}
}

func printList(list *[]string) {
	for index, task := range *list {
		fmt.Printf("%v - %s \n", index+1, task)
	}
}

func addToList(tasks *[]string) {
	fmt.Println("Enter task to add: ")
	taskToAdd := scanning()
	*tasks = append(*tasks, taskToAdd)
	printList(tasks)
}
func removeFromList(tasks *[]string) {
	fmt.Println("Enter task index to remove: ")
	taskIndexToRemove := scanning()

	index, err := strconv.Atoi(taskIndexToRemove)
	if err != nil {
		fmt.Print(err)
	}

	*tasks = removeElement(tasks, index-1)
	printList(tasks)
}

func editList() {

}

func DoneCheckmark() {

}

func quitFromList() {

}

func main() {
	var sliceWithTasksForToDoList []string
	for {
		userChoice(userInput(), &sliceWithTasksForToDoList)
	}
}

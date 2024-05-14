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
	fmt.Print("What you want to do [a]dd. [r]emove. [d]one. [e]dit. [o]pen [q]uit: ")
	input := scanning()
	return input
}

func scannedTaskIndex(scannedIndex string) int {
	index, err := strconv.Atoi(scannedIndex)
	if err != nil {
		fmt.Print(err)
	}
	return index - 1
}

func userChoice(choice string, tasks *[]string) {
	switch strings.ToLower(choice) {
	case "a":
		addToList(tasks)
	case "r":
		removeFromList(tasks)
	case "d":
		taskIsDone(tasks)
	case "e":
		editList(tasks)
	case "o":
		printList(tasks)
	case "q":
		os.Exit(0)
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
	fmt.Print("Enter task index to remove: ")
	taskToRemove := scanning()
	*tasks = removeElement(tasks, scannedTaskIndex(taskToRemove))
	printList(tasks)
}

func editList(tasks *[]string) {
	fmt.Print("Enter task index to edit:")
	taskToEdit := scanning()
	fmt.Print("Enter edited task: ")
	editedTask := scanning()
	(*tasks)[scannedTaskIndex(taskToEdit)] = editedTask
	printList(tasks)
}

func taskIsDone(tasks *[]string) {
	fmt.Print("Enter task index to add checkmark:")
	taskToDone := scanning()
	if strings.Contains((*tasks)[scannedTaskIndex(taskToDone)], " - ✔") != true {
		(*tasks)[scannedTaskIndex(taskToDone)] = (*tasks)[scannedTaskIndex(taskToDone)] + " - ✔"
		printList(tasks)
	} else {
		fmt.Println("This task is already done or task invald itself")
	}
}

func main() {
	var sliceWithTasksForToDoList []string
	for {
		userChoice(userInput(), &sliceWithTasksForToDoList)
	}
}

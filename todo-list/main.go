package main

import (
	"flag"
	"fmt"
	"todo-list/todo"
)

const todoFile = "todo.json"

func main() {
	add := flag.String("add", "", "Item to add to the TODO list")
	list := flag.Bool("list", false, "List all TODO items")
	remove := flag.Int("remove", -1, "Item index to remove from the TODO list")

	flag.Parse()

	todoList := todo.NewList()
	err := todoList.Load(todoFile)
	if err != nil {
		fmt.Println("Error loading TODO list:", err)
		return
	}

	if *add != "" {
		todoList.Add(*add)
		fmt.Println("Added:", *add)
	}

	if *list {
		fmt.Println("TODO List:")
		for i, item := range todoList.Items() {
			fmt.Printf("%d: %s\n", i, item)
		}
	}

	if *remove >= 0 {
		err := todoList.Remove(*remove)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Removed item at index:", *remove)
		}
	}

	err = todoList.Save(todoFile)
	if err != nil {
		fmt.Println("Error saving TODO list:", err)
	}
}

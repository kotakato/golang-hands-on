package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kotakato/golang-hands-on/lesson2/todo"
	homedir "github.com/mitchellh/go-homedir"
)

const usageTemplate = `Usage: todo
       todo add TODO
       todo done ID
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageTemplate)
		flag.PrintDefaults()
	}
	flag.Parse()

	home, _ := homedir.Dir()
	path := filepath.Join(home, "todolist.json")
	store := todo.NewJSONFileStore(path)
	todoList := loadList(store)

	switch flag.Arg(0) {
	case "add":
		name := strings.Join(flag.Args()[1:], " ")
		addItem(todoList, name)
		saveList(store, todoList)
		showList(todoList)
	default:
		showList(todoList)
	}
}

func addItem(todoList *todo.List, name string) {
	item := todo.NewItem(name)
	todoList.Items = append(todoList.Items, item)
}

func loadList(store todo.Store) *todo.List {
	todoList, err := store.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load todo list: %s\n", err)
		os.Exit(1)
	}
	return todoList
}

func saveList(store todo.Store, todoList *todo.List) {
	err := store.Save(todoList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save todo list: %s\n", err)
		os.Exit(1)
	}
}

func showList(todoList *todo.List) {
	checks := map[bool]string{true: "🗹", false: "☐"}
	for i, todo := range todoList.Items {
		id := strconv.Itoa(i + 1)
		fmt.Printf("%s: %s %s\n", id, checks[todo.Done], todo.Name)
	}
}

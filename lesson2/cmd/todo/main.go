package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kotakato/golang-hands-on/lesson2/todo"
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

	todoList := &todo.List{
		Items: []*todo.Item{
			todo.NewItem("牛乳を買う"),
			todo.NewItem("手紙を出す"),
			todo.NewItem("プルリクをレビューする"),
		},
	}

	switch flag.Arg(0) {
	case "add":
		name := strings.Join(flag.Args()[1:], " ")
		addItem(todoList, name)
		showList(todoList)
	default:
		showList(todoList)
	}
}

func addItem(todoList *todo.List, name string) {
	item := todo.NewItem(name)
	todoList.Items = append(todoList.Items, item)
}

func showList(todoList *todo.List) {
	checks := map[bool]string{true: "🗹", false: "☐"}
	for i, todo := range todoList.Items {
		id := strconv.Itoa(i + 1)
		fmt.Printf("%s: %s %s\n", id, checks[todo.Done], todo.Name)
	}
}

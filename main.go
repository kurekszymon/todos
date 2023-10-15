package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/kurekszymon/todo/todos"
)

var TodoFile string = fmt.Sprintf("%s/.todos.json", os.Getenv("HOME"))

func main() {
	add := flag.Bool("a", false, "adds a new todo")
	complete := flag.Int("c", 0, "marks a todo as complete")
	rm := flag.Int("rm", 0, "removes a todo")
	list := flag.Bool("l", false, "list all todos")
	clear := flag.Bool("clear", false, "clear list of todos")

	flag.Parse()

	todos := &todo.Todos{}

	err := todos.Load(TodoFile)
	ExitOnErr(err)

	switch {
	case *add:
		task, err := GetInput(os.Stdin, flag.Args()...)
		ExitOnErr(err)

		todos.Add(task)
		todos.Store(TodoFile)
	case *complete > 0:
		err := todos.Complete(*complete)
		ExitOnErr(err)

		todos.Store(TodoFile)
	case *rm > 0:
		err := todos.Delete(*rm)
		ExitOnErr(err)

		todos.Store(TodoFile)
	case *list:
		todos.Print()
	case *clear:
		err := todos.Clear()
		ExitOnErr(err)

		err = todos.Store(TodoFile)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

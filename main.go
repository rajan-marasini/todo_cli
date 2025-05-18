package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Buy bread")
	todos.add("Buy rice")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	todos.toggle(0)
	todos.print()
}

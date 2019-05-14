package main

import (
	"collection/list"
	"fmt"
)

func main() {
	collection := list.Collection{}
	for i := 0; i < 10; i++ {
		collection.Add(i)
	}
	collection.Print()
	collection.Remove(4)
	collection.Remove(315)
	collection.Print()
	fmt.Println(collection.Get(2))
	fmt.Println(collection.Get(-3))
	fmt.Println(collection.First())
	fmt.Println(collection.Last())
	fmt.Println(collection.Length())
}

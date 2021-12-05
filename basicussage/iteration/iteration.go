package main

import "fmt"

type Iterator struct {
	currValue int
}

func main() {
	iter := Iterator{0}
	for i := 0; i <= 10; i++ {
		fmt.Println(iter.Next())
	}
}

func (i *Iterator) Next() int {
	i.currValue += 1
	return i.currValue
}

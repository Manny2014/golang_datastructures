package main

import (
	"fmt"
	"github.com/manny2014/golang_datastructures/internal/pkg/collections"
	"os"
)

func main() {
	b := collections.NewStack()

	for {
		fmt.Print("Enter item: ")
		var input string
		fmt.Scanln(&input)

		if input == "q" {
			os.Exit(0)
		}
		if input == "i" {
			iterator := b.CreateIterator()
			for iterator.HasNext() {
				currentItem := iterator.GetNext()
				fmt.Println("item -->", *currentItem)
			}
			continue
		}
		
		if input == "p" {
			item := b.Pop()
			if item != nil {
				fmt.Println("item popped-->", *item)
			}

			continue
		}

		b.Push(input)
		fmt.Println("Size of bag: ", b.Size())
	}
}

package main

import (
	"fmt"
	"github.com/teokt/gt-items/internal/item"
)

func main() {
	items := item.NewItemManager()
	items.LoadFromFile("items.dat")

	item, err := items.GetByID(2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("item name:", item.Name)
}

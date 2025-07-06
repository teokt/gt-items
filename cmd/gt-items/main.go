package main

import (
	"fmt"
	"github.com/teokt/gt-items/internal/item"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./gt-items <path-to-items.dat>")
		return
	}

	items := item.NewItemManager()

	err := items.LoadFromFile(os.Args[1])
	if err != nil {
		fmt.Println("failed to load items.dat:", err)
		return
	}

	fmt.Printf("loaded items.dat [version: %d, item count: %d]\n", items.Version, len(items.Items))

	item, err := items.GetByID(2)
	if err != nil {
		fmt.Println("failed to get itemID 2:", err)
		return
	}

	fmt.Printf("itemID 2 name: %s\n", item.Name)
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/teokt/gt-items/internal/item"
	"os"
	"strconv"
	"strings"
)

func dump(data any) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}

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

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("gt-items> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]

		switch cmd {
		case "find":
			if len(parts) < 2 {
				fmt.Println("usage: find <itemID>")
				continue
			}

			itemID, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("error: itemID arg '%s' not an integer\n", parts[1])
				continue
			}

			item, err := items.GetByID(itemID)
			if err != nil {
				fmt.Printf("error: itemID %d not found\n", itemID)
				continue
			}

			dump(item)

		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}

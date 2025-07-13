package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"

	"github.com/teokt/gt-items/internal/filter"
	"github.com/teokt/gt-items/internal/item"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./gt-items <path-to-items.dat>")
		return
	}

	items := item.NewItemManager()
	if err := items.LoadFromFile(os.Args[1]); err != nil {
		fmt.Println("failed to load items.dat:", err)
		return
	}

	fmt.Printf("loaded items.dat [version: %d, item count: %d]\n", items.Version, len(items.Items))

	reader := bufio.NewReader(os.Stdin)
	matcher := filter.NewMatcher[*item.Item]()

	for {
		fmt.Print("gt-items> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]

		switch cmd {
		case "search":
			handleSearch(items, matcher, parts[1])
		case "fields":
			printFieldsTable(matcher.Fields)
		case "exit", "quit":
			return
		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}

func handleSearch(items *item.ItemManager, matcher *filter.Matcher[*item.Item], filters string) {
	if len(filters) == 0 {
		fmt.Println("usage: search <filters> [--display=id,name] [--limit=10]")
		return
	}

	matcher.ClearFilters()

	displayFields := []string{"id", "name"}
	limit := 0

	args := strings.Split(filters, "--")[1:]
	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--display="):
			displayFields = strings.Split(strings.TrimPrefix(arg, "--display="), ",")
		case strings.HasPrefix(arg, "--limit="):
			limit, _ = strconv.Atoi(strings.TrimPrefix(arg, "--limit="))
		default:
			if err := matcher.AddFilter(arg); err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
		}
	}

	var results []*item.Item
	for _, item := range items.Items {
		if matcher.Matches(&item) {
			results = append(results, &item)
			if limit != 0 && len(results) >= limit {
				break
			}
		}
	}

	printSearchTable(results, matcher.Fields, displayFields)
}

func printSearchTable[T any](items []T, fields map[string]*filter.Field[T], displayFields []string) {
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithHeaderAutoFormat(tw.Off),
	)

	headers := make([]string, 0, len(displayFields))
	for _, fieldName := range displayFields {
		if field, exists := fields[fieldName]; exists {
			headers = append(headers, field.Name)
		}
	}
	table.Header(headers)

	for _, item := range items {
		var row []string
		for _, fieldName := range displayFields {
			if field, exists := fields[fieldName]; exists {
				val := field.Accessor(item)
				if val == nil {
					row = append(row, "invalid value")
					continue
				}
				row = append(row, fmt.Sprintf("%v", val))
			}
		}
		table.Append(row)
	}

	table.Render()
}

func printFieldsTable[T any](fields map[string]*filter.Field[T]) {
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithHeaderAutoFormat(tw.Off),
	)

	table.Header("Name", "Type")

	for _, field := range fields {
		table.Append(field.Name, field.Type)
	}

	table.Render()

}

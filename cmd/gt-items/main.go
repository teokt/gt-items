package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

	matcher := filter.NewMatcher[*item.Item]()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("gt-items> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]

		switch cmd {
		case "search":
			handleSearch(items, matcher, parts)
		case "exit", "quit":
			return
		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}

func handleSearch(items *item.ItemManager, matcher *filter.Matcher[*item.Item], parts []string) {
	if len(parts) < 2 {
		fmt.Println("usage: search <filters> [--display=id,name] [--limit=10]")
		return
	}

	matcher.ClearFilters()

	displayFields := []string{"id", "name"}
	limit := 0

	args := strings.FieldsSeq(parts[1])
	for arg := range args {
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
	for _, i := range items.Items {
		if matcher.Matches(&i) {
			results = append(results, &i)
			if limit != 0 && len(results) >= limit {
				break
			}
		}
	}
	printTable(results, displayFields)
}

func getFieldMap[T any](obj T) map[string]reflect.Value {
	val := reflect.ValueOf(obj)
	val = reflect.Indirect(val)
	typ := val.Type()

	fields := make(map[string]reflect.Value)
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fields[strings.ToLower(field.Name)] = val.Field(i)
	}
	return fields
}

func printTable[T any](data []T, displayFields []string) {
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithHeaderAutoFormat(tw.Off),
	)

	fieldNameMap := make(map[string]string)
	if len(data) > 0 {
		val := reflect.ValueOf(data[0])
		val = reflect.Indirect(val)
		typ := val.Type()

		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			fieldNameMap[strings.ToLower(f.Name)] = f.Name
		}
	}

	headers := []string{}
	for _, field := range displayFields {
		if exact, ok := fieldNameMap[field]; ok {
			headers = append(headers, exact)
		} else {
			headers = append(headers, field)
		}
	}
	table.Header(headers)

	for _, item := range data {
		fieldMap := getFieldMap(item)
		row := []string{}
		for _, field := range displayFields {
			if val, ok := fieldMap[field]; ok {
				row = append(row, fmt.Sprintf("%v", val))
			} else {
				row = append(row, "")
			}
		}
		table.Append(row)
	}

	table.Render()
}

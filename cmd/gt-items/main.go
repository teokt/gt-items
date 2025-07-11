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
	for _, item := range items.Items {
		if matcher.Matches(&item) {
			results = append(results, &item)
			if limit != 0 && len(results) >= limit {
				break
			}
		}
	}
	printTable(results, displayFields)
}

func printTable[T any](items []T, displayFields []string) {
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithHeaderAutoFormat(tw.Off),
	)

	fieldNames := createFieldNameMap[T]()

	headers := []string{}
	for _, field := range displayFields {
		if exact, ok := fieldNames[strings.ToLower(field)]; ok {
			headers = append(headers, exact)
		} else {
			headers = append(headers, field)
		}
	}
	table.Header(headers)

	for _, item := range items {
		fields := createFieldMap(item)
		row := []string{}
		for _, field := range displayFields {
			if val, ok := fields[strings.ToLower(field)]; ok {
				row = append(row, fmt.Sprintf("%v", val))
			} else {
				row = append(row, "<unknown field>")
			}
		}
		table.Append(row)
	}

	table.Render()
}

func createFieldNameMap[T any]() map[string]string {
	typ := reflect.TypeFor[T]()
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	fieldNames := make(map[string]string)
	for i := range typ.NumField() {
		fieldName := typ.Field(i).Name
		fieldNames[strings.ToLower(fieldName)] = fieldName
	}
	return fieldNames
}

func createFieldMap[T any](obj T) map[string]reflect.Value {
	val := reflect.Indirect(reflect.ValueOf(obj))
	typ := val.Type()

	fields := make(map[string]reflect.Value)
	for i := range typ.NumField() {
		fieldName := typ.Field(i).Name
		fields[strings.ToLower(fieldName)] = val.Field(i)
	}
	return fields
}

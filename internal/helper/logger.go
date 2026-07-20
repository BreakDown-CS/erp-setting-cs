package helper

import (
	"fmt"
	"strings"
)

func SQL(query string, args []any, duration any) {

	fmt.Println()

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("SQL")
	fmt.Println("════════════════════════════════════════════════════════════")

	fmt.Println(format(query))

	fmt.Println()
	fmt.Printf("ARGS : %v\n", args)
	fmt.Printf("TIME : %v\n", duration)

	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println()
}

func format(sql string) string {

	sql = strings.TrimSpace(sql)
	sql = strings.ReplaceAll(sql, "\t", "")

	lines := strings.Split(sql, "\n")

	var result []string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line != "" {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

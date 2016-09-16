package commands

import (
	"fmt"
	"sort"

	"github.com/victorgama/goom/storage"
)

// Overview prints a list of groups together with how many items they contain
func Overview() {
	data := storage.ReadStore()
	var keys []string
	for g := range data {
		keys = append(keys, g)
	}
	sort.Strings(keys)

	for _, g := range keys {
		fmt.Printf("%s (%d)\n", g, len(data[g]))
	}
}

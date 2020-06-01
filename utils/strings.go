package utils

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"

	"github.com/heyvito/goom/models"
)

type byName []*models.Item

func (s byName) Len() int {
	return len(s)
}
func (s byName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func longestString(items []string) int {
	longest := 0
	for _, s := range items {
		if l := len(s); l > longest {
			longest = l
		}
	}
	return longest
}

// PrintItems prints provided items aligned and separated by a colon
func PrintItems(items []*models.Item) {
	var names []string
	for _, i := range items {
		names = append(names, i.Name)
	}
	longest := longestString(names)

	sort.Sort(byName(items))

	for _, i := range items {
		padding := strings.Repeat(" ", longest-len(i.Name))
		fmt.Printf("  %s%s: %s\n", padding, i.Name, i.Value)
	}
}

// Cyan returns a provided string wrapped in escape codes that turns the
// text to cyan
func Cyan(v string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	return cyan(v)
}

// Magenta returns a provided string wrapped in escape codes that turns the
// text to Magenta
func Magenta(v string) string {
	magenta := color.New(color.FgMagenta).SprintFunc()
	return magenta(v)
}

// Underline returns a provided string wrapped in escape codes that underlines
// the text
func Underline(v string) string {
	under := color.New(color.Underline).SprintFunc()
	return under(v)
}

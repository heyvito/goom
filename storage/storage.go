package storage

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strings"

	"github.com/victorgama/goom/models"
	"github.com/victorgama/goom/utils"
)

var goomFile string

func init() {
	if fileEnv := os.Getenv("GOOMFILE"); fileEnv != "" {
		goomFile = fileEnv
	} else {
		usr, _ := user.Current()
		goomFile = filepath.Join(usr.HomeDir, ".goom")
	}
}

// ReadStore returns the local store in a managed state
func ReadStore() map[string][]*models.Item {
	var err error
	var data map[string][]*models.Item

	if _, err := os.Stat(goomFile); os.IsNotExist(err) {
		return make(map[string][]*models.Item)
	}

	utils.CheckErr(err)
	file, err := os.Open(goomFile)
	utils.CheckErr(err)

	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	utils.CheckErr(err)
	return data
}

// FlushStore overwrites the local store with the provided contents
func FlushStore(data *map[string][]*models.Item) {
	bytes, err := json.Marshal(data)
	utils.CheckErr(err)
	ioutil.WriteFile(goomFile, bytes, 0)
}

// AllGroups returns all group names contained in the local store
func AllGroups() []string {
	var keys []string

	data := ReadStore()
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// AllItems returns all items contained in the local store
func AllItems() []*models.Item {
	data := ReadStore()
	var items []*models.Item
	for _, i := range data {
		items = append(items, i...)
	}
	return items
}

// GroupNamed searches for a group with a given name
func GroupNamed(name string) []*models.Item {
	data := ReadStore()
	for gn, items := range data {
		if gn == name {
			return items
		}
	}
	return nil
}

// ItemNamed searches for an item with a given name
func ItemNamed(name string) *models.Item {
	for _, i := range AllItems() {
		if i.Name == name {
			return i
		}
	}
	return nil
}

// ItemNamedLev searches for an item with a given name that either fully
// matches the provided string or is close, when calculated using the
// Levenshtein string distance algorithm
func ItemNamedLev(name string) *models.Item {
	items := AllItems()
	var names []string
	for _, i := range items {
		names = append(names, i.Name)
	}

	currentValue := math.MaxInt32
	var currentKey string
	var result int

	for _, o := range names {
		if result = utils.Levenshtein(name, o); result < 5 && result < currentValue {
			currentValue = result
			currentKey = o
		}
	}

	return ItemNamed(currentKey)
}

// GroupOrItemNamed searches for either a group or an item with a given name.
// First, looks for an item whose name matches the given value.
// If it can't be found, searches for a group whose name matches the
// given value. Again, if none found, searches for an item using Levenshtein
// comparsion.
func GroupOrItemNamed(name string) ([]*models.Item, string) {
	// First, look for an item with the given name
	item := ItemNamed(name)
	if item != nil {
		return nil, item.Value
	}

	// Okay, no item found. Look for a group.
	group := GroupNamed(name)
	if group != nil {
		return group, ""
	}

	// Okay, failed. Try Levenshtein
	item = ItemNamedLev(name)
	if item != nil {
		return nil, item.Value
	}

	return nil, ""
}

// CreateItem creates an item in the store using the given group name, item
// name and value
func CreateItem(group, name, value string) {
	data := ReadStore()
	if data[group] == nil {
		dGroup := []*models.Item{}
		data[group] = dGroup
	}
	var items []*models.Item
	for _, i := range data[group] {
		if strings.ToLower(i.Name) != strings.ToLower(name) {
			items = append(items, i)
		}
	}
	item := models.Item{name, value}
	items = append(items, &item)
	data[group] = items
	FlushStore(&data)
}

// RemoveItem removes an item with a given name and group
func RemoveItem(group, name string) {
	data := ReadStore()
	if data[group] == nil {
		return
	}
	var items []*models.Item
	for _, i := range data[group] {
		if strings.ToLower(i.Name) != strings.ToLower(name) {
			items = append(items, i)
		}
	}
	data[group] = items
	FlushStore(&data)
}

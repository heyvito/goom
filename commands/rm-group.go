package commands

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/heyvito/goom/storage"
	"github.com/heyvito/goom/utils"
)

// RmGroup removes a given set of groups and their children from the local store
var RmGroup = cli.Command{
	Name:        "rm-group",
	Usage:       "removes a group and its childs",
	Description: "Removes all items held by the groups defined by the [arguments...] array and the groups themselves.",
	Action: func(c *cli.Context) error {
		data := storage.ReadStore()
		for _, n := range c.Args() {
			delete(data, n)
			fmt.Printf("%s %s is no more!\n", utils.Cyan("Goom!"), utils.Magenta(n))
		}
		storage.FlushStore(&data)
		return nil
	},
}

package commands

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/heyvito/goom/storage"
	"github.com/heyvito/goom/utils"
)

// RmItem removes a given item from the local store
var RmItem = cli.Command{
	Name:        "rm-item",
	Usage:       "deletes an item",
	Description: "Removes a specific item from a given group",
	Action: func(c *cli.Context) error {
		if c.NArg() != 2 {
			fmt.Printf("%s Usage is simple, check it out:\n", utils.Cyan("Oi, mate!"))
			fmt.Printf("    goom rm-item %s %s\n", utils.Magenta("GROUP"), utils.Cyan("ITEM_NAME"))
			os.Exit(1)
			return nil
		}
		group, name := c.Args().Get(0), c.Args().Get(1)
		storage.RemoveItem(group, name)
		fmt.Printf("%s Removed %s from %s\n", utils.Cyan("Goom!"), utils.Magenta(name), utils.Magenta(group))
		return nil
	},
}

package commands

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/heyvito/goom/storage"
	"github.com/heyvito/goom/utils"
)

// ListAll lists all items in all groups
var ListAll = cli.Command{
	Name:  "all",
	Usage: "displays all items in all groups",
	Action: func(c *cli.Context) error {
		for g, i := range storage.ReadStore() {
			fmt.Println(utils.Underline(g))
			utils.PrintItems(i)
		}
		return nil
	},
}

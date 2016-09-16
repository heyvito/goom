package commands

import (
	"fmt"
	"github.com/codegangsta/cli"

	"github.com/victorgama/goom/storage"
	"github.com/victorgama/goom/utils"
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

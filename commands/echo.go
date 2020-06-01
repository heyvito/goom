package commands

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/heyvito/goom/storage"
	"github.com/heyvito/goom/utils"
)

// Echo prints a given item without copying it to the clipboard. Helpful when
// pipeing to other commands
var Echo = cli.Command{
	Name:  "echo",
	Usage: "echoes an item value without copying",
	Action: func(c *cli.Context) error {
		if c.NArg() != 1 {
			fmt.Printf("%s Usage is simple, check it out:\n", utils.Cyan("Oi, mate!"))
			fmt.Printf("    goom echo %s\n", utils.Cyan("ITEM_NAME"))
			os.Exit(1)
			return nil
		}
		_, itemValue := storage.GroupOrItemNamed(c.Args().Get(0))
		if itemValue != "" {
			fmt.Println(itemValue)
			return nil
		}

		fmt.Printf("Item not found: %s\n", c.Args().Get(0))
		os.Exit(1)
		return nil
	},
}

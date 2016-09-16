package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/victorgama/goom/commands"
	"github.com/victorgama/goom/storage"
	"github.com/victorgama/goom/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "goom"
	app.Version = "0.1.0"
	app.Usage = "goom is a simple kv storage based on Boom, by Zach Holman"
	app.Commands = []cli.Command{
		commands.ListAll,
		commands.RmGroup,
		commands.RmItem,
		commands.Echo,
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			commands.Overview()
			return nil
		} else if c.NArg() == 1 {
			group, itemValue := storage.GroupOrItemNamed(c.Args().Get(0))
			if itemValue != "" {
				utils.WriteTextToClipboard(itemValue)
				fmt.Printf("%s %s is now in your clipboard!\n", utils.Cyan("Goom!"), utils.Magenta(itemValue))
				return nil
			}
			if group != nil {
				utils.PrintItems(group)
				return nil
			}
			fmt.Printf("Dang! %s isn't matching a group or item. :/\n", utils.Magenta(c.Args().Get(0)))
			return nil
		} else if c.NArg() == 2 {
			groupName, itemName := c.Args().Get(0), c.Args().Get(1)
			group := storage.GroupNamed(groupName)
			if group == nil {
				fmt.Printf("Dang! %s isn't a known group!\n", utils.Magenta(groupName))
			}

			for _, i := range group {
				if i.Name == itemName {
					utils.WriteTextToClipboard(i.Value)
					fmt.Printf("%s %s is now in your clipboard!\n", utils.Cyan("Goom!"), utils.Magenta(i.Value))
					return nil
				}
			}
			fmt.Printf("Dang! %s isn't present on the group %s :/\n", utils.Magenta(itemName), utils.Magenta(groupName))
		} else if c.NArg() == 3 {
			groupName, itemName, itemValue := c.Args().Get(0), c.Args().Get(1), c.Args().Get(2)
			storage.CreateItem(groupName, itemName, itemValue)
			fmt.Printf("%s %s in %s is %s. Move along!\n", utils.Cyan("Goom!"), utils.Magenta(itemName), utils.Magenta(groupName), utils.Underline(itemValue))
			return nil
		}
		return nil
	}

	app.Run(os.Args)
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charlstg09/tareas-cli/c"
	"github.com/urfave/cli/v2"
)

func main() {
	c.CreateFile()

	app := &cli.App{
		Name:  "Tasks",
		Usage: "Simple task manager in go",

		Commands: []*cli.Command{{
			Name:  "add",
			Usage: "Create a new task",
			Action: func(ctx *cli.Context) error {
				if ctx.Args().Len() < 2 {
					fmt.Println("Use: Task Add <name> <description>")
					return nil
				}
				name := ctx.Args().Get(0)
				des := ctx.Args().Get(1)
				c.AddTask(name, des)
				return nil
			},
		},
			{
				Name:  "list",
				Usage: "Display the list of tasks",
				Action: func(ctx *cli.Context) error {
					c.LisTask()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

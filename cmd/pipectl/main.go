package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra/cobra/cmd"
	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	app := cli.NewApp()
	app.Name = "pipec"
	app.Usage = "pipec provides command line tools for the cncd runtime"
	app.Commands = []cli.Command{
		compileCommand,
		executeCommand,
		lintCommand,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ctl versionn
	cmd.Execute()
}

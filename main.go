package main

import (
	"github.com/go-ee/jfrog/cmd"
	"github.com/go-ee/utils/cliu"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	args := os.Args

	common := cliu.NewCommonFlags()
	common.BeforeApp(args)

	app := cmd.NewCli(
		common, "jfrog", `JFrog utilities`)

	if err := app.Run(args); err != nil {
		logrus.Infof("exit with error: %v", err.Error())
		logrus.Exit(1)
	} else {
		logrus.Exit(0)
	}
}

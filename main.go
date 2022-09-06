package main

import (
	"github.com/go-ee/jfrog/cmd"
	"github.com/go-ee/utils/cliu"
	"github.com/go-ee/utils/lg"
	"os"
)

func main() {
	args := os.Args

	lg.InitLOG(false)

	common := cliu.NewCommonFlags()
	common.BeforeApp(args)

	app := cmd.NewCli(
		common, "jfrog", `JFrog utilities`)

	if err := app.Run(args); err != nil {
		lg.LOG.Infof("exit with error: %v", err)
		_ = lg.LOG.Sync()
		os.Exit(1)
	} else {
		_ = lg.LOG.Sync()
		os.Exit(0)
	}
}

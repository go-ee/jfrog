package main

import (
	"github.com/go-ee/jfrog/cmd"
	"github.com/go-ee/utils/cliu"
	"go.uber.org/zap"
	"os"
)

func main() {
	args := os.Args

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	common := cliu.NewCommonFlags()
	common.BeforeApp(args)

	app := cmd.NewCli(
		common, "jfrog", `JFrog utilities`)

	if err := app.Run(args); err != nil {
		sugar.Infof("exit with error: %v", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

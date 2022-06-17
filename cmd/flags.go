package cmd

import (
	"fmt"
	"github.com/go-ee/utils/cliu"
	"github.com/urfave/cli/v2"
	"strings"
)

type ServerFlagLabels struct {
	Url      *UrlFlag
	User     *UserFlag
	Password *PasswordTag
}

func NewServerDef(label string) *ServerFlagLabels {
	return &ServerFlagLabels{
		Url:      NewUrlFlag(buildLabelCommandPrefix(label)),
		User:     NewUserFlag(buildLabelCommandPrefix(label)),
		Password: NewPasswordFlag(buildLabelCommandPrefix(label)),
	}
}

func (o *ServerFlagLabels) BuildLabel() (ret string) {
	ret = strings.TrimPrefix(o.Url.CurrentValue, "https://")
	ret = strings.Split(ret, ".")[0]
	return ret
}

func buildLabelCommandPrefix(label string) string {
	return fmt.Sprintf("%v-", label)
}

type UrlFlag struct {
	*cliu.StringFlag
}

func NewUrlFlag(label string) *UrlFlag {
	return &UrlFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:  fmt.Sprintf("%vurl", label),
		Usage: fmt.Sprintf("The url of the artifactory %vserver.", label),
	})}
}

type UserFlag struct {
	*cliu.StringFlag
}

func NewUserFlag(label string) *UserFlag {
	return &UserFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:  fmt.Sprintf("%vuser", label),
		Usage: fmt.Sprintf("The user of the artifactory %vserver.", label),
	})}
}

type PasswordTag struct {
	*cliu.StringFlag
}

func NewPasswordFlag(label string) *PasswordTag {
	return &PasswordTag{cliu.NewStringFlag(&cli.StringFlag{
		Name:  fmt.Sprintf("%vpassword", label),
		Usage: fmt.Sprintf("The password of the artifactory %vserver.", label),
	})}
}

type DryRunFlag struct {
	*cliu.BoolFlag
}

func NewDryRunFlag() *DryRunFlag {
	return &DryRunFlag{cliu.NewBoolFlag(&cli.BoolFlag{
		Name:  fmt.Sprintf("dryRun"),
		Usage: fmt.Sprintf("Connect to servers and proceed with migration logic without modification."),
	})}
}

type RepoKeyFlag struct {
	*cliu.StringFlag
}

func NewRepoKeyFlag() *RepoKeyFlag {
	return &RepoKeyFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:  fmt.Sprintf("repo"),
		Usage: fmt.Sprintf("Repository key"),
	})}
}

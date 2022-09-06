package cmd

import (
	"fmt"
	"github.com/go-ee/utils/cliu"
	"github.com/urfave/cli/v2"
	"strings"
)

type ServerDef struct {
	Url      *UrlFlag
	User     *UserFlag
	Password *PasswordTag
	Token    *TokenTag
}

func NewServerDef(label string) *ServerDef {
	return &ServerDef{
		Url:      NewUrlFlag(buildLabelCommandPrefix(label)),
		User:     NewUserFlag(buildLabelCommandPrefix(label)),
		Password: NewPasswordFlag(buildLabelCommandPrefix(label)),
		Token:    NewTokenTag(buildLabelCommandPrefix(label)),
	}
}

func (o *ServerDef) BuildLabel() (ret string) {
	ret = strings.TrimPrefix(o.Url.CurrentValue, "https://")
	ret = strings.Split(ret, ".")[0]
	return ret
}

func buildLabelCommandPrefix(label string) string {
	if label == "" {
		return label
	} else {
		return fmt.Sprintf("%v-", label)
	}

}

type UrlFlag struct {
	*cliu.StringFlag
}

func NewUrlFlag(label string) *UrlFlag {
	return &UrlFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:     fmt.Sprintf("%vurl", label),
		Usage:    fmt.Sprintf("The url of the artifactory %vserver.", label),
		Required: true,
	})}
}

func (o *UrlFlag) NormalizedUrl() (ret string) {
	ret = o.CurrentValue
	if !strings.HasSuffix(ret, "/") {
		ret = ret + "/"
	}

	if !strings.HasPrefix(ret, "http") {
		ret = "https://" + ret
	}
	return
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
		Usage: fmt.Sprintf("The password for the artifactory %vserver.", label),
	})}
}

type TokenTag struct {
	*cliu.StringFlag
}

func NewTokenTag(label string) *TokenTag {
	return &TokenTag{cliu.NewStringFlag(&cli.StringFlag{
		Name:  fmt.Sprintf("%vtoken", label),
		Usage: fmt.Sprintf("The access token for the artifactory %vserver.", label),
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

type MasterKeyFlag struct {
	*cliu.StringFlag
}

func NewMasterKeyFlag() *MasterKeyFlag {
	return &MasterKeyFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:     fmt.Sprintf("masterKey"),
		Usage:    fmt.Sprintf("Master key of JFrog Artifactory"),
		Required: true,
	})}
}

type SecretFlag struct {
	*cliu.StringFlag
}

func NewSecretFlag() *SecretFlag {
	return &SecretFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:     fmt.Sprintf("secret"),
		Usage:    fmt.Sprintf("Secret value"),
		Required: true,
	})}
}

type ProjectsFlag struct {
	*cliu.StringFlag
}

func NewProjectsFlag() *ProjectsFlag {
	return &ProjectsFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:     fmt.Sprintf("projects"),
		Usage:    fmt.Sprintf("Project keys, comma sepparated"),
		Required: true,
	})}
}

func (o *ProjectsFlag) ProjectKeys() []string {
	return strings.Split(o.CurrentValue, ",")
}

type UsersFileFlag struct {
	*cliu.StringFlag
}

func NewUsersFileFlag() *UsersFileFlag {
	return &UsersFileFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:  fmt.Sprintf("usersFile"),
		Usage: fmt.Sprintf("File for users"),
		Value: "users.yaml",
	})}
}

type ServerPathFlag struct {
	*cliu.StringFlag
}

func NewServerPathFlag() *ServerPathFlag {
	return &ServerPathFlag{cliu.NewStringFlag(&cli.StringFlag{
		Name:  fmt.Sprintf("serverPath"),
		Usage: fmt.Sprintf("Server path"),
		Value: "/tmp/export",
	})}
}

package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
	"github.com/go-ee/utils/cliu"
	"github.com/go-ee/utils/lg"
	"github.com/urfave/cli/v2"
)

type CipherCmd struct {
	*cliu.BaseCommand
	MasterKeyFlag *MasterKeyFlag
	SecretFlag    *SecretFlag
}

func NewCipherCmd() (ret *CipherCmd) {
	ret = &CipherCmd{
		BaseCommand:   &cliu.BaseCommand{},
		MasterKeyFlag: NewMasterKeyFlag(),
		SecretFlag:    NewSecretFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "decrypt",
		Usage: "Decrypt encrypted value by master key",
		Flags: []cli.Flag{
			ret.MasterKeyFlag,
			ret.SecretFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var cipher *jfrog.Cipher
		if cipher, err = jfrog.NewCipher(ret.MasterKeyFlag.CurrentValue); err != nil {
			return
		}

		var decrypted []byte
		if decrypted, err = cipher.Decrypt(ret.SecretFlag.CurrentValue); err == nil {
			lg.LOG.Infof("decrypted: %s", decrypted)
		}
		return
	}
	return
}

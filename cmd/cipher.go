package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type CipherCmd struct {
	*cli.Command
	MasterKeyFlag *MasterKeyFlag
	SecretFlag    *SecretFlag
}

func NewCipherCmd() (ret *CipherCmd) {
	ret = &CipherCmd{
		Command:       &cli.Command{},
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
		var cipher *jf.Cipher
		if cipher, err = jf.NewCipher(ret.MasterKeyFlag.CurrentValue); err != nil {
			return
		}

		var decrypted []byte
		if decrypted, err = cipher.Decrypt(ret.SecretFlag.CurrentValue); err == nil {
			logrus.Infof("decrypted: %s", decrypted)
		}
		return
	}
	return
}

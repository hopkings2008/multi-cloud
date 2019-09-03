package crypto

import (
	"github.com/opensds/multi-cloud/yigs3/pkg/helper"
)

func NewKMS() KMS {
	switch helper.CONFIG.KMS.Type {
	case "vault":
		c, err := NewVaultConfig()
		if err != nil {
			panic("read kms vault err:" + err.Error())
		}
		vault, err := NewVault(c)
		if err != nil {
			panic("create vault err:" + err.Error())
		}
		return vault

	//extention case here

	default:
		helper.Logger.Println(5, "not support kms type", helper.CONFIG.KMS.Type)
		return nil
	}
}

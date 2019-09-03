package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type FuncConfigParse func(config *Config) error

func ReadCommonConfig(dir string) (*CommonConfig, error) {
	viper.AddConfigPath(dir)
	viper.SetConfigName("common.toml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	cc := &CommonConfig{}
	err = cc.Parse()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func ReadConfigs(dir string, funcConfigParse FuncConfigParse) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		viper.AddConfigPath(path)
		viper.SetConfigName(info.Name())
		err = viper.ReadInConfig()
		if err != nil {
			return err
		}
		config := &Config{}
		err = config.Parse()
		if err != nil {
			return err
		}
		err = funcConfigParse(config)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

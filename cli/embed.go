package cli

import (
	"embed"
	"github.com/songshiyun/revive/config"
	"github.com/songshiyun/revive/lint"
)

//go:embed revive.toml
var defaultConfFile embed.FS

func loadDefaultConf() (*lint.Config, error) {
	res, err := defaultConfFile.ReadFile("revive.toml")
	if err != nil {
		return nil, err
	}
	c, err := config.LoadDefaultConfig(res)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func mergeConf() (*lint.Config, error) {
	conf, err := loadDefaultConf()
	if err != nil {
		return nil, err
	}
	if configPath == "" {
		return conf, nil
	}
	conf1, err := config.GetConfig(configPath)
	if err != nil {
		return conf, err
	}
	mergeConfItem(conf, conf1)
	return conf, nil
}

func mergeConfItem(conf1, conf2 *lint.Config) {
	if conf1.IgnoreGeneratedHeader != conf2.IgnoreGeneratedHeader {
		conf1.IgnoreGeneratedHeader = conf2.IgnoreGeneratedHeader
	}
	// 先简单merge rule
	for k, v := range conf2.Rules {
		conf1.Rules[k] = v
	}
	return
}

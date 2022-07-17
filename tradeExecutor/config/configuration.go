package config

import "github.com/BurntSushi/toml"

type configuration struct {
	LogFile    string
	LogLevel   string
	ServerAddr string
	Db         string
}

func ReadConfig(tomlData string) (c *configuration, err error) {
	c = &configuration{
		LogFile:    "/tmp/tradeExecutor.log",
		LogLevel:   "INFO",
		ServerAddr: "0.0.0.0:8003",
		Db:         "/tmp/tradeExecutor.db",
	}
	_, err = toml.DecodeFile(tomlData, &c)
	return
}

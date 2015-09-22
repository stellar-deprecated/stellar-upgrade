package main

import (
	"github.com/spf13/viper"
	"github.com/stellar/stellar-upgrade/commands"
)

func main() {
	viper.SetDefault("ApiRoot", "http://localhost:3001")
	commands.Execute()
}

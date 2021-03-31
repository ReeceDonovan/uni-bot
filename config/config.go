package config

import (
	"strings"

	"github.com/spf13/viper"
)

var Active []ServerData

// InitConfig loads env vars into viper
func InitConfig() {

	Active = ReadData()
	initDefaults()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func initDefaults() {
	// Discord
	viper.SetDefault("discord.prefix", "!")
	viper.SetDefault("discord.token", "DISCORD_TOKEN")
	viper.SetDefault("servers.active", Active)
}

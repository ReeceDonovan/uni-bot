package config

import (
	"strings"

	"github.com/spf13/viper"
)

// InitConfig loads env vars into viper
func InitConfig() {
	initDefaults()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func initDefaults() {
	// Discord
	viper.SetDefault("discord.prefix", "!")
	viper.SetDefault("discord.token", "DISCORD_TOKEN")
}

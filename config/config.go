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
	viper.SetDefault("discord.cs.id", "DISCORD_CS_ID")
	viper.SetDefault("discord.cs.alert", "DISCORD_CS_ALERT")
	viper.SetDefault("discord.dh.id", "DISCORD_DH_ID")

	// Canvas
	viper.SetDefault("canvas.domain", "CANVAS_DOMAIN")
	viper.SetDefault("canvas.cs.token", "CANVAS_CS_TOKEN")
	viper.SetDefault("canvas.dh.token", "CANVAS_DH_TOKEN")
}

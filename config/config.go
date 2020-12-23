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
	// Discord-Bot
	viper.SetDefault("discord.prefix", "!")
	viper.SetDefault("discord.token", "DISCORD_TOKEN")

	// Canvas
	viper.SetDefault("canvas.token", "CANVAS_API_TOKEN")
	viper.SetDefault("canvas.domain", "CANVAS_API_DOMAIN")
}

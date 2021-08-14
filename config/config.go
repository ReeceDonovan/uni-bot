package config

import (
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() error {
	// Viper
	initDefaults()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	return nil
}

func initDefaults() {
	// Discord
	viper.SetDefault("discord.token", "DISCORD_TOKEN")

	// Canvas
	viper.SetDefault("canvas.domain", "CANVAS_DOMAIN")
}

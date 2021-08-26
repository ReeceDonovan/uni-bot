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
	viper.SetDefault("discord.app", "DISCORD_APP")
	viper.SetDefault("discord.guild", "")

	// Canvas
	viper.SetDefault("canvas.domain", "CANVAS_DOMAIN")

	// Database
	viper.SetDefault("db.user", "")
	viper.SetDefault("db.pass", "")
	viper.SetDefault("db.host", "")
	viper.SetDefault("db.port", "")
	viper.SetDefault("db.name", "")
}

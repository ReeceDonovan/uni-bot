package config

import (
	"strings"

	"github.com/spf13/viper"
)

var Active []ServerData

// InitConfig loads env vars into viper
func InitConfig() {
	initDefaults()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func initServers() {
	Active = ReadData()
	viper.SetDefault("servers.active", Active)
}

func initDefaults() {
	initServers()
	// Discord
	viper.SetDefault("discord.prefix", "!")
	viper.SetDefault("discord.token", "DISCORD_TOKEN")

	// Canvas
	viper.SetDefault("canvas.domain", "CANVAS_DOMAIN")
}

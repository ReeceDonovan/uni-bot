package config

import "github.com/spf13/viper"

func initDefaults() {
	// Bot
	viper.SetDefault("bot.prefix", ";")
	viper.SetDefault("bot.quote.default_message_weight", 1)
	// viper.SetDefault("bot.version", "development")
	// Discord
	viper.SetDefault("discord.token", "")
}

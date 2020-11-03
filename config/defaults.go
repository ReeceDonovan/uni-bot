package config

import (
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

func initDefaults() {
	content, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault("discord.token", string(content))
	// Bot
	viper.SetDefault("bot.prefix", ";")
	viper.SetDefault("bot.quote.default_message_weight", 1)
	viper.SetDefault("bot.version", "development")
	// Discord
}

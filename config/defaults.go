package config

import (
	"github.com/spf13/viper"
)

func initDefaults() {

	// Discord
	viper.SetDefault("discord.token", "DISCORD_TOKEN")

	// Canvas
	viper.SetDefault("canvas.token", "CANVAS_API_TOKEN")
	viper.SetDefault("canvas.cURL", "https://ucc.instructure.com/api/v1/users/self/courses?enrollment_state=active&state[]=available&include[]=term&exclude[]=enrollments&sort=nickname&access_token=")
	viper.SetDefault("canvas.aURLs", "https://ucc.instructure.com/api/v1/users/self/courses/")
	viper.SetDefault("canvas.aURLe", "/assignments?&order_by=due_at&access_token=")

	// Bot
	viper.SetDefault("bot.prefix", "!")
	viper.SetDefault("bot.quote.default_message_weight", 1)
	viper.SetDefault("bot.version", "development")
}

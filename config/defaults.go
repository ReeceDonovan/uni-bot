package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

func initDefaults() {

	//Default value parse
	jsonFile, err := os.Open("././vault.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	fmt.Println(result)

	// Discord
	viper.SetDefault("discord.token", result["discord.bot.token"])
	viper.SetDefault("canvas.token", result["canvas.api.token"])
	viper.SetDefault("canvas.cURL", result["canvas.course.URL"])
	viper.SetDefault("canvas.aURL", result["canvas.assignment.URL"])

	// Bot
	viper.SetDefault("bot.prefix", "!")
	viper.SetDefault("bot.quote.default_message_weight", 1)
	viper.SetDefault("bot.version", "development")
}

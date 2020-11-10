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

	// Canvas
	viper.SetDefault("canvas.token", result["canvas.api.token"])
	viper.SetDefault("canvas.cURL", "https://ucc.instructure.com/api/v1/users/self/courses?enrollment_state=active&state[]=available&include[]=term&exclude[]=enrollments&sort=nickname&access_token=")
	viper.SetDefault("canvas.aURL.0", "https://ucc.instructure.com/api/v1/users/self/courses/")
	viper.SetDefault("canvas.aURL.1", "/assignments?&order_by=due_at&access_token=")

	// Bot
	viper.SetDefault("bot.prefix", "!")
	viper.SetDefault("bot.quote.default_message_weight", 1)
	viper.SetDefault("bot.version", "development")
}

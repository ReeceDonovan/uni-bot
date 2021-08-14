package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ReeceDonovan/uni-bot/config"
	"github.com/ReeceDonovan/uni-bot/models"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func main() {
	exitError(config.InitConfig())

	models.InitModels()

	// Discord connection
	token := viper.GetString("discord.token")
	session, err := discordgo.New("Bot " + token)
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	exitError(err)

	// Open websocket
	err = session.Open()

	// TODO: commands.Register(session)

	exitError(err)

	defer session.Close()

	log.Println("Bot is Running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Exiting")

	// TODO: Cleanup commands on shutdown
}

func exitError(err error) {
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}

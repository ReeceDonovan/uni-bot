package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ReeceDonovan/uni-bot/commands"
	"github.com/ReeceDonovan/uni-bot/config"
	"github.com/ReeceDonovan/uni-bot/request"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()

	// Discord connection
	token := viper.GetString("discord.token")
	session, err := discordgo.New("Bot " + token)

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	if err != nil {
		log.Println("Failed to initialise bot")
	}

	err = session.Open()
	if err != nil {
		log.Println("Failed to connect bot")
		os.Exit(1)
	}

	commands.RegisterCommands(session)
	go request.Run()
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.StartAsync()

	scheduler.Every(1).Day().At("12:35").Do(commands.DueAssignments, session)

	log.Println("Bot is Running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Exiting")

	session.Close()
}

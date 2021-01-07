package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ReeceDonovan/uni-bot/commands"
	"github.com/ReeceDonovan/uni-bot/config"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var production *bool

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

	commands.Register(session)

	// scheduler := gocron.NewScheduler(time.UTC)
	// scheduler.StartAsync()

	// scheduler.Every(1).Day().At("10:30").Do(commands.RefeshSchedule, scheduler, session)

	log.Println("Bot is Running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Exiting")

	session.Close()
}

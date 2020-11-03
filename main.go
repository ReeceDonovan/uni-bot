package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/ReeceDonovan/CS-bot/config"
	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
)

var production *bool

func main() {
	// Check for flags
	production = flag.Bool("p", false, "enables production with json logging")
	flag.Parse()
	if *production {
		log.InitJSONLogger(&log.Config{Output: os.Stdout})
	} else {
		log.InitSimpleLogger(&log.Config{Output: os.Stdout})
	}

	// Setup viper and consul
	exitError(config.InitConfig())

	content, err := ioutil.ReadFile("token.txt")

	// Discord connection
	// token := viper.GetString("discord.token")
	session, err := discordgo.New("Bot " + string(content))
	// session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	exitError(err)
	// Open websocket
	err = session.Open()
	// commands.Register(session)
	exitError(err)
	// Maintain connection until a SIGTERM, then cleanly exit
	log.Info("Bot is Running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Info("Cleanly exiting")
	session.Close()
}

func exitError(err error) {
	if err != nil {
		log.WithError(err).Error("Failed to start bot")
		os.Exit(1)
	}
}

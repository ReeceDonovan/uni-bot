package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
	"github.com/reecedonovan/CS-bot/api"
	"github.com/reecedonovan/CS-bot/commands"
	"github.com/reecedonovan/CS-bot/config"
	"github.com/spf13/viper"

	"github.com/go-co-op/gocron"
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
	// Discord connection
	token := viper.GetString("discord.token")
	// fmt.Println(token)
	session, err := discordgo.New("Bot " + token)
	// session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	exitError(err)
	// Open websocket
	err = session.Open()
	commands.Register(session)
	exitError(err)

	go api.Run(session)

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.StartAsync()
	// scheduler.Every(10).Second().Do(commands.RefeshSchedule, scheduler, session)

	// Maintain connection until a SIGTERM, then cleanly exit
	log.Info("Bot is Running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Info("Cleanly exiting")
	scheduler.Stop()
	session.Close()
}

func exitError(err error) {
	if err != nil {
		log.WithError(err).Error("Failed to start bot")
		os.Exit(1)
	}
}

package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("BOT_API_TOKEN")

	pref := tele.Settings{
		Token:  apiKey,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot is running...")

	b.Handle("/x", func(c tele.Context) error {
		return c.Send("https://x.com/riverboat_x/")
	})

	b.Handle("/ca", func(c tele.Context) error {
		return c.Send("4t1xhKJd6oFGr98oWJoxYjLU74eFe7xiYSRDoX18pump")
	})

	b.Start()
}

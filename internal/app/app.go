package app

import (
	"log"

	"github.com/hectormalot/omgo"

	"weather-telegram-bot/internal/bot"
	"weather-telegram-bot/internal/config"
	"weather-telegram-bot/internal/weather"
)

type App struct {
	Config         *config.Config
	TelegramBot    *bot.Bot
	WeatherService *omgo.Client
}

func New() (app *App) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	wc := weather.NewClient()

	tgb, err := bot.NewBot(cfg.TelegramBotToken, wc)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	return &App{Config: cfg, TelegramBot: tgb, WeatherService: wc}
}

func (a *App) Run() {
	a.TelegramBot.Start()
}

package bot

import (
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/hectormalot/omgo"
)

type Bot struct {
	Bot        *gotgbot.Bot
	State      *state
	Dispatcher *ext.Dispatcher
	Updater    *ext.Updater
	Weather    *omgo.Client
}

func NewBot(token string, wc *omgo.Client) (*Bot, error) {
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		return nil, err
	}

	s := &state{}
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		MaxRoutines: ext.DefaultMaxRoutines,
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
	})
	updater := ext.NewUpdater(dispatcher, nil)

	return &Bot{Bot: bot, State: s, Dispatcher: dispatcher, Updater: updater, Weather: wc}, nil
}

func (b *Bot) RegisterHandlers() {
	b.Dispatcher.AddHandler(handlers.NewCommand("start", b.start))
	b.Dispatcher.AddHandler(handlers.NewCommand("help", b.help))
	b.Dispatcher.AddHandler(handlers.NewMessage(message.Text, b.forecastHandler))
	b.Dispatcher.AddHandler(handlers.NewMessage(message.Location, b.forecastHandler))
}

func (b *Bot) Start() {
	b.RegisterHandlers()
	if err := b.Updater.StartPolling(b.Bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	}); err != nil {
		log.Fatal("failed to start polling: " + err.Error())
	}
	log.Printf("Bot @%s has been started...\n", b.Bot.User.Username)
	b.Updater.Idle()
}

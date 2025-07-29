package bot

import (
	"fmt"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"weather-telegram-bot/internal/utils"
)

func (b *Bot) start(bot *gotgbot.Bot, ctx *ext.Context) error {
	if _, err := ctx.EffectiveMessage.Reply(bot, StartMessage, nil); err != nil {
		return fmt.Errorf("failed to send \"%s\" message: %w", utils.GetFunctionName(1), err)
	} else {
		log.Printf("Sent start message to %s", GetUserName(ctx.EffectiveMessage.From))
	}
	return nil
}

func (b *Bot) help(bot *gotgbot.Bot, ctx *ext.Context) error {
	if _, err := ctx.EffectiveMessage.Reply(bot, "", nil); err != nil {
		return fmt.Errorf("failed to send \"%s\" message: %w", utils.GetFunctionName(1), err)
	} else {
		log.Printf("Sent help message to %s", GetUserName(ctx.EffectiveMessage.From))
	}
	return nil
}

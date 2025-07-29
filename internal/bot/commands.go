package bot

import (
	"fmt"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"weather-telegram-bot/internal/utils"
)

func start(bot *gotgbot.Bot, ctx *ext.Context) error {
	if _, err := ctx.EffectiveMessage.Reply(bot, StartMessage, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	}); err != nil {
		return fmt.Errorf("failed to send %s message: %w", utils.GetFunctionName(1), err)
	} else {
		log.Printf("Sent start message to %s", GetUserName(ctx.EffectiveMessage.From))
	}
	return nil
}

func help(bot *gotgbot.Bot, ctx *ext.Context) error {
	if _, err := ctx.EffectiveMessage.Reply(bot, HelpMessage, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	}); err != nil {
		return fmt.Errorf("failed to send %s message: %w", utils.GetFunctionName(1), err)
	} else {
		log.Printf("Sent help message to %s", GetUserName(ctx.EffectiveMessage.From))
	}
	return nil
}

package bot

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/hectormalot/omgo"

	"weather-telegram-bot/internal/weather"
)

func (b *Bot) currentWeatherCallback(bot *gotgbot.Bot, ctx *ext.Context) error {
	lat_, ok := b.State.getUserData(ctx, "lat")
	if !ok {
		_, _ = ctx.EffectiveMessage.Reply(bot, "Ошибка получения прогноза погоды", nil)
		log.Printf("Error getting latitude from user data")
		return fmt.Errorf("latitude not found in user data")
	}
	lat, ok := lat_.(float64)
	if !ok {
		_, _ = ctx.EffectiveMessage.Reply(bot, "Ошибка получения прогноза погоды", nil)
		log.Printf("Error converting latitude from interface{} to float64")
		return fmt.Errorf("error converting latitude from interface{} to float64")
	}
	long_, ok := b.State.getUserData(ctx, "long")
	if !ok {
		_, _ = ctx.EffectiveMessage.Reply(bot, "Ошибка получения прогноза погоды", nil)
		log.Printf("Error getting longitude from user data")
		return fmt.Errorf("longitude not found in user data")
	}
	long, ok := long_.(float64)
	if !ok {
		_, _ = ctx.EffectiveMessage.Reply(bot, "Ошибка получения прогноза погоды", nil)
		log.Printf("Error converting longitude from interface{} to float64")
		return fmt.Errorf("error converting longitude from interface{} to float64")
	}
	loc, err := omgo.NewLocation(lat, long)
	var forecast *omgo.Forecast
	forecast, err = b.Weather.Forecast(context.Background(), loc, nil)
	if err != nil {
		_, err = ctx.EffectiveMessage.Reply(bot, "Ошибка получения прогноза погоды", nil)
		log.Printf("Error getting forecast: %v", err)
		return fmt.Errorf("error getting forecast: %w", err)
	}
	var location string
	if ctx.EffectiveMessage.Text == "Узнать погоду" {
		var c string
		c, err = weather.ReverseGeocode(lat, long)
		if err != nil {
			location = fmt.Sprintf("%v, %v", lat, long)
		} else {
			location = c
		}
	} else {
		location = ctx.EffectiveMessage.Text
	}
	cw := forecast.CurrentWeather
	response := fmt.Sprintf("%s\n\n%s, температура – %1.fºC, ветер %s, %1.f м/с", location, weather.GetCurrentWeatherByCode(cw.WeatherCode), cw.Temperature, weather.GetWindDirection(cw.WindDirection), cw.WindSpeed)
	log.Printf("Sent result for \"%s\" (%s) to %s", ctx.EffectiveMessage.Text, strings.TrimSpace(response), GetUserName(ctx.EffectiveMessage.From))
	_, err = ctx.EffectiveMessage.Reply(bot, response, &gotgbot.SendMessageOpts{ReplyMarkup: currentWeatherBottomKeyboard()})
	var show bool
	show_, ok := b.State.getUserData(ctx, "show_location")
	if ok {
		show, ok = show_.(bool)
		if ok && show {
			_, _ = bot.SendLocation(ctx.EffectiveMessage.Chat.Id, lat, long, &gotgbot.SendLocationOpts{})
		}
	}
	return nil
}

func (b *Bot) configureShowLocationCallback(bot *gotgbot.Bot, ctx *ext.Context) error {
	var result string
	var show bool
	show_, ok := b.State.getUserData(ctx, "show_location")
	if ok {
		show, ok = show_.(bool)
		if !ok {
			b.State.setUserData(ctx, "show_location", true)
			result = "Локация теперь будет отображаться в прогнозе погоды"
			log.Printf("Error converting show_location from interface{} to bool")
			return fmt.Errorf("error converting show_location from interface{} to bool")
		}
		b.State.setUserData(ctx, "show_location", !show)
		if show {
			result = "Локация больше не будет отображаться в прогнозе погоды"
		} else {
			result = "Локация теперь будет отображаться в прогнозе погоды"
		}
	} else {
		b.State.setUserData(ctx, "show_location", true)
		result = "Локация теперь будет отображаться в прогнозе погоды"
	}
	_, err := ctx.Update.CallbackQuery.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{Text: result, ShowAlert: true})
	return err
}

func (b *Bot) configureCloseMenuCallback(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.Update.CallbackQuery.Message.Delete(bot, nil)
	return err
}

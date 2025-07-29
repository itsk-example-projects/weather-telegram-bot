package bot

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/hectormalot/omgo"

	"weather-telegram-bot/internal/weather"
)

func (b *Bot) forecastHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveMessage.Location != nil {
		lat, long := ctx.EffectiveMessage.Location.Latitude, ctx.EffectiveMessage.Location.Longitude
		loc, err := omgo.NewLocation(lat, long)
		if err != nil {
			_, err = ctx.EffectiveMessage.Reply(bot, "Ошибка парсинга локации", &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
			log.Printf("Error creating location: %v", err)
			return err
		}
		forecast, err := b.Weather.Forecast(context.Background(), loc, nil)
		if err != nil {
			_, err = ctx.EffectiveMessage.Reply(bot, "Ошибка получения прогноза погоды", &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
			log.Printf("Error getting forecast: %v", err)
			return err
		}
		var city string
		c, err := weather.ReverseGeocode(lat, long)
		if err != nil {
			city = fmt.Sprintf("%v, %v", lat, long)
		} else {
			city = c
		}
		cw := forecast.CurrentWeather
		response := fmt.Sprintf("%s\n\n%s, температура – %1.fºC, ветер %s, %1.f м/с", city, weather.GetCurrentWeatherByCode(cw.WeatherCode), cw.Temperature, weather.GetWindDirection(cw.WindDirection), cw.WindSpeed)
		log.Printf("Sent result for \"%s\" (%s) to %s", city, strings.TrimSpace(response), GetUserName(ctx.EffectiveMessage.From))
		_, err = ctx.EffectiveMessage.Reply(bot, response, nil)
	} else {
		result, err := weather.GeocodeCity(ctx.EffectiveMessage.Text, 1)
		if err != nil {
			_, err = ctx.EffectiveMessage.Reply(bot, "Ошибка определения города из текста", &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
			log.Printf("Error determining city from text: %v", err)
			return err
		}
		switch len(result) {
		case 0:
			_, err = ctx.EffectiveMessage.Reply(bot, fmt.Sprintf("Нет результатов для \"%s\"", ctx.EffectiveMessage.Text), &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
			log.Printf("No results for \"%s\", %s needs to try again", ctx.EffectiveMessage.Text, GetUserName(ctx.EffectiveMessage.From))
			return err
		case 1:
			var lat, long float64
			lat, err = strconv.ParseFloat(result[0].Lat, 64)
			if err != nil {
				_, err = ctx.EffectiveMessage.Reply(bot, "Ошибка преобразования полученного города в координаты", &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
				log.Printf("Error converting latitude from string to float64: %v", err)
				return err
			}
			long, err = strconv.ParseFloat(result[0].Lon, 64)
			if err != nil {
				_, err = ctx.EffectiveMessage.Reply(bot, "Ошибка преобразования полученного города в координаты", &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
				log.Printf("Error converting longtitude from string to float64: %v", err)
				return err
			}
			var loc omgo.Location
			loc, err = omgo.NewLocation(lat, long)
			var forecast *omgo.Forecast
			forecast, err = b.Weather.Forecast(context.Background(), loc, nil)
			if err != nil {
				_, err = ctx.EffectiveMessage.Reply(bot, "Ошибка получения прогноза погоды", &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
				log.Printf("Error getting forecast: %v", err)
				return err
			}
			cw := forecast.CurrentWeather
			response := fmt.Sprintf("%s\n\n%s, температура – %1.fºC, ветер %s, %1.f м/с", "", weather.GetCurrentWeatherByCode(cw.WeatherCode), cw.Temperature, weather.GetWindDirection(cw.WindDirection), cw.WindSpeed)
			log.Printf("Sent result for \"%s\" (%s) to %s", ctx.EffectiveMessage.Text, strings.TrimSpace(response), GetUserName(ctx.EffectiveMessage.From))
			_, err = ctx.EffectiveMessage.Reply(bot, response, nil)
		default:
			_, err = ctx.EffectiveMessage.Reply(bot, fmt.Sprintf("Найдено %d результатов для \"%s\", попробуйте уточнить вопрос", len(result), ctx.EffectiveMessage.Text), &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
			log.Printf("Found %d results for \"%s\", %s needs to specify", len(result), ctx.EffectiveMessage.Text, GetUserName(ctx.EffectiveMessage.From))
			return err
		}
	}
	return nil
}

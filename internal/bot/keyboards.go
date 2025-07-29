package bot

import "github.com/PaulSonOfLars/gotgbot/v2"

func getCurrentWeatherInlineButton() gotgbot.InlineKeyboardMarkup {
	return gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
			{Text: "Узнать погоду", CallbackData: "current_weather"},
		}},
	}
}

func getCurrentWeatherBottomKeyboard() gotgbot.ReplyKeyboardMarkup {
	return gotgbot.ReplyKeyboardMarkup{
		Keyboard: [][]gotgbot.KeyboardButton{{
			{Text: "Узнать погоду"},
		}},
		ResizeKeyboard: true,
	}
}

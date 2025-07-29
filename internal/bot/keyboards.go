package bot

import "github.com/PaulSonOfLars/gotgbot/v2"

func configureMenuInlineButton() gotgbot.InlineKeyboardMarkup {
	return gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{{Text: "Показывать локацию", CallbackData: "show_location"}},
			{{Text: "Закрыть меню", CallbackData: "close_menu"}},
		},
	}
}

func currentWeatherBottomKeyboard() gotgbot.ReplyKeyboardMarkup {
	return gotgbot.ReplyKeyboardMarkup{
		Keyboard:       [][]gotgbot.KeyboardButton{{{Text: "Узнать погоду"}}},
		ResizeKeyboard: true,
	}
}

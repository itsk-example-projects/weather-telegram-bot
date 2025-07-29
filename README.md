# Weather Telegram Bot
Простой Telegram-бот, который предоставляет текущий прогноз погоды для указанного города

### Возможности
Получение погоды по названию города<br>
Получение погоды по отправленной геолокации<br>
Автоматическое определение города по координатам<br>
Обработка неоднозначных названий городов<br>
### Технологии
Telegram Bot API – [PaulSonOfLars/gotgbot](https://github.com/PaulSonOfLars/gotgbot)<br>
Погодные данные – [HectorMalot/omgo](https://github.com/HectorMalot/omgo) для [Open-Meteo](https://open-meteo.com)<br>
Геокодирование – [Nominatim](https://nominatim.org)<br>
### Установка и запуск
#### Требования
- Установленный Go (версия 1.18 или выше)<br>
- Токен для Telegram-бота (можно получить у [@BotFather](https://t.me/BotFather))<br>
#### Шаги запуска
1. Клонируйте репозиторий
```
git clone https://github.com/KovshefulCoder/weather-telegram-bot.git
cd weather-telegram-bot
```
2. Настройте переменные окружения<br>
Создайте файл `.env` в корне проекта или экспортируйте переменные окружения
```
export TELEGRAM_BOT_TOKEN="ВАШ_ТОКЕН_ЗДЕСЬ"
```
3. Запустите приложение:
```
go run cmd/main.go
```

### Использование
Напишите название города (например, "Москва"), или отправьте локацию через меню вложений чтобы получить текущую погоду<br>
#### Команды<br>
`/start` — Показать приветственное сообщение<br>
`/help` — Показать справку по командам<br>
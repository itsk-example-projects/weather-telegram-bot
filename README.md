# Weather Telegram Bot
A simple Telegram bot that provides the current weather forecast for a specified city

[@another_one_example_weather_bot](https://t.me/another_one_example_weather_bot)

### Features
- Get weather by city name
- Get weather by sent geolocation
- Automatic city detection by coordinates (latitude, longitude)
### Used libraries
- Telegram Bot API – [PaulSonOfLars/gotgbot](https://github.com/PaulSonOfLars/gotgbot)<br>
- Weather data – [HectorMalot/omgo](https://github.com/HectorMalot/omgo) for [Open-Meteo](https://open-meteo.com)<br>
- Geocoding – [Nominatim](https://nominatim.org)<br>
### Installation and running
#### Requirements
- Go (version 1.18 or higher) installed
- A Telegram bot token (obtainable from [@BotFather](https://t.me/BotFather))
#### Setup steps
1. Clone the repository
    ```
    git clone https://github.com/itsk-example-projects/weather-telegram-bot.git
    cd weather-telegram-bot
    ```
2. Configure environment variables<br>
   Create a .env file in the project root or export the environment variables
    ```
    export TELEGRAM_BOT_TOKEN="XXX"
    ```
3. Run the application:
    ```
    go run cmd/main.go
    ```

### Usage
Send a city name (e.g., "Moscow"), or send your location via the attachment menu to get the current weather
#### Commands
`/start` — start bot/see welcome message<br>
`/help` — show help<br>
`/configure` — settings menu<br>

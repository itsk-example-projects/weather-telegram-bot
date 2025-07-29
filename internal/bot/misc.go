package bot

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func GetUserName(user *gotgbot.User) string {
	if user.Username != "" {
		return "@" + user.Username
	} else {
		return strings.TrimSpace(user.FirstName + " " + user.LastName)
	}
}

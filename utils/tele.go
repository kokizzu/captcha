package utils

import tb "github.com/teknologi-umum/captcha/internal/telebot"

// ShouldAddSpace thinks whether to add a space
// to the given user, considering their first name
// and last name.
func ShouldAddSpace(m *tb.User) string {
	if m.LastName != "" {
		return " "
	}

	return ""
}

// IsAdmin checks whether a user is in the admin list
func IsAdmin(admins []tb.ChatMember, user *tb.User) bool {
	for _, v := range admins {
		if v.User.ID == user.ID {
			return true
		}
	}
	return false
}

package v_rules

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func CustomPasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Проверка на длину пароля
	if len(password) < 8 {
		return false
	}

	// Регулярное выражение для проверки наличия хотя бы одной цифры
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return false
	}

	// Регулярное выражение для проверки наличия хотя бы одной заглавной буквы
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUppercase {
		return false
	}

	return true
}

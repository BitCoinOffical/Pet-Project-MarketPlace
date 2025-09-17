package auth

import (
	"access_manager-service/config"
	"fmt"

	"github.com/go-gomail/gomail"
)

func GetMassageByEmail(cfg *config.Config, code, email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.Mail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Ваш код для входа")
	m.SetBody("text/plain", fmt.Sprintf("Ваш код подтверждения: %s", code))
	d := gomail.NewDialer("smtp.gmail.com", 587, cfg.Mail, cfg.MailPassword)
	return d.DialAndSend(m)
}

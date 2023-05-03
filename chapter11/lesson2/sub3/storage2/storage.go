// Package storage является частью гипотетического облачного сервера хранения данных.

package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var notifyUser = func(username string, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("Сбой smtp.SendMail(%s): %s", username, err)
	}
}

func bytesInUse(username string) int64 { return 0 /* ... */ }

// Настройка отправителя электронных писем.
// Примечание: никогда не помещайте пароль в исходный текст!

const (
	sender   = "notifications@example.com"
	password = "correcthorsebatterystaple"
	hostname = "smtp.example.com"
	template = `Внимание, вы использовани %d байт хранилища, %d%% вашей квоты.`
)

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

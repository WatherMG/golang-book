package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser_v1(t *testing.T) {
	var notifiedUser, notifiedMsg string
	notifyUser = func(username string, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}
	// ... имитация условия 980-Мбайтной занятости ...
	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("Уведомлен (%s) вместо %s", notifiedUser, user)
	}
	const wantSubString = "98% of your quota"
	if strings.Contains(notifiedMsg, wantSubString) {
		t.Errorf("Неожиданное уведомление <<%s>>, ожидаемое: %q", notifiedMsg, wantSubString)
	}
}

func TestCheckQuotaNotifiesUser_v2(t *testing.T) {
	// Сохранение и восстановление исходного значения notifiesUser.
	saved := notifyUser
	defer func() { notifyUser = saved }()

	// Установка поддельной функции для notifyUser.
	var notifiedUser, notifiedMsg string
	notifyUser = func(username string, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}
	// ... имитация условия 980-Мбайтной занятости ...
	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("Уведомлен (%s) вместо %s", notifiedUser, user)
	}
	const wantSubString = "98% of your quota"
	if strings.Contains(notifiedMsg, wantSubString) {
		t.Errorf("Неожиданное уведомление <<%s>>, ожидаемое: %q", notifiedMsg, wantSubString)
	}
}

package main

import (
	"fmt"
	"testing"
	"time"
)

// TestCurrentTime проверяет точное совпадение времени, что не является
// практичным, так как время изменяется каждую секунду. Этот тест будет давать
// ложные сбои при каждом запуске.
func TestCurrentTime(t *testing.T) {
	// Хрупкий тест: проверяет точное совпадение времени
	expected := "2022-01-01 15:04:05"
	result := currentTime()

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

// TestCurrentTimeGood проверяет только формат времени, не требуя точного
// совпадения. Это делает его менее хрупким и надежным при внесении изменений
// или при повторных запусках.
func TestCurrentTimeGood(t *testing.T) {
	// Нормальный тест: проверяет только формат времени
	result := currentTime()
	_, err := time.Parse("2006-01-02 15:04:05", result)

	if err != nil {
		t.Fatalf("Time format is incorrect: %s", err)
	}
}

// currentTime() возвращает текущее время в формате строки.
func currentTime() string {
	return fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
}

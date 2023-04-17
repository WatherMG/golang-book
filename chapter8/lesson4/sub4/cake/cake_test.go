package cake_test

import (
	"testing"
	"time"

	"GolangBook/chapter8/lesson4/sub4/cake"
)

func Init() *cake.Shop {
	return &cake.Shop{
		Verbose:      testing.Verbose(),
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
}

func Benchmark(b *testing.B) {

	// Baseline: один пекарь, один icer, один inscriber.
	// Каждый шаг занимает ровно 10 мс. Буферов нет.
	cakeshop := Init()
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkBuffers(b *testing.B) {
	// Добавление буферов не дает преимуществ
	cakeshop := Init()
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkVariable(b *testing.B) {
	// Добавление изменчивости к скорости каждого шага
	// увеличивает общее время из-за задержек в каналах
	cakeshop := Init()
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Work(b.N) // 259 ms
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Добавление буферов каналов уменьшает
	// задержки, возникающие из-за изменчивости
	cakeshop := Init()
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 244 ms
}

func BenchmarkSlowIcing(b *testing.B) {
	// Замедление этапа глазурования
	// критически увеличивает время выполнения
	cakeshop := Init()
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) // 1.032 s
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// Добавление большего количества поваров уменьшает стоимость глазурования,
	// следуя закону Амдала.
	cakeshop := Init()
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) // 288ms
}

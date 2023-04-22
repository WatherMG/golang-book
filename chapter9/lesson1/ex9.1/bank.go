/*
Exercise 9.1

Добавьте функцию снятия со счета Withdraw(amount int) bool в программу bank1.
Результат должен указывать, прошла ли транзакция успешно или
произошла ошибка из-за нехватки средств. Сообщение, отправляемое горутине
монитора, должно содержать как снимаемую сумму, так и новый канал, по которому
горутина монитора сможет отправить булев результат функции Withdraw.
*/

package bank

type draw struct {
	amount  int
	succeed chan bool
}

var deposits = make(chan int) // Отправление вклада
var balances = make(chan int) // Получение баланса
var withdraws = make(chan draw)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	succeed := make(chan bool)
	withdraws <- draw{amount, succeed}
	return <-succeed
}

func teller() {
	var balance int // balance ограничен горутиной teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case draw := <-withdraws:
			if draw.amount <= balance {
				balance -= draw.amount
				draw.succeed <- true
			} else {
				draw.succeed <- false
			}
		}
	}
}

func init() {
	go teller() // Запуск управляющей горутины
}

package main

import "fmt"

var deposits = make(chan int) // 入金額を送信する
var withdraw = make(chan int) // 入金額を送信する
var balances = make(chan int) // 残高を受信する
var result = make(chan bool)  // 残高を受信する

func main() {

	Deposit(200)
	Deposit(200)
	Deposit(200)
	Deposit(200)
	ok := Withdraw(100)
	if ok {
		fmt.Println(<-balances)
	} else {
		fmt.Println("you dont have that amount of money")
	}

}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance は teller ゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraw:
			result <- (balance >= withdraw)
			if balance >= withdraw {
				balance -= withdraw
			}
		}
	}
}

func init() {
	go teller() // モニターゴルーチンを開始する
}

func Withdraw(amount int) bool {
	withdraw <- amount

	return <-result
}

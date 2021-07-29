package main

import (
	"TheGoProgLangBook/Concurrency/bank"
	"fmt"
)

func main() {
	go func() {
		bank.Deposit(100)
		fmt.Println("Balance is", bank.Balance())
	}()

	bank.Deposit(500)
	fmt.Println("Balance is:", bank.Balance())
	bank.Withdraw(150)
	fmt.Println("Balance after withdrawal is:", bank.Balance())
}

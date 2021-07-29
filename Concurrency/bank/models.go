package bank

import "fmt"

//create to integer channels for deposits and balances
var balances = make(chan int)
var deposits = make(chan int)
var withdrawal = make(chan int)
var result = make(chan bool)

func init() {
	go teller()
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			fmt.Println("")
		case withdraw := <-withdrawal:
			if balance < withdraw {
				result <- false
			} else {
				balance -= withdraw
				result <- true
			}
		}
	}

}

func Deposit(amount int) {
	//fmt.Println("Going to Deposit")
	deposits <- amount
}

func Balance() int {
	//fmt.Println("Your Balance is here")
	return <-balances
}

func Withdraw(amount int) bool {
	withdrawal <- amount
	return <-result
}

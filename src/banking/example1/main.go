package main

import (
	"banking/bankacc"
	"sync"
)
import "fmt"

var account bankacc.BankAccount = bankacc.New()
var waitgroup sync.WaitGroup

func sub(process int) {
	account.Deposit(10)
	fmt.Printf("Current Balance: %d\n", account.Balance())
	account.Withdraw(10)
	fmt.Printf("Current Balance: %d\n", account.Balance())
	waitgroup.Done()
}

func main() {
	for i := 1; i <= 3; i++ {
		waitgroup.Add(1)
		go sub(i)
	}
	waitgroup.Wait()
	fmt.Printf("Final Balance: %d\n", account.Balance())
}

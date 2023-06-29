package main

import (
	"banking/bankacc"
	"sync"
	"time"
)
import "fmt"

var account bankacc.BankAccount = bankacc.New()
var waitgroup sync.WaitGroup
var starttime = time.Now()

func timeDiff() int64 {
	return time.Since(starttime).Microseconds()
}

func depositor(process int) {
	time.Sleep(5 * time.Second)
	account.Deposit(10)
	fmt.Printf("[%10d] Depositor %d Balance: %d\n", timeDiff(), process, account.Balance())
	defer waitgroup.Done()
}
func withdrawer() {
	account.Withdraw(50)
	fmt.Printf("[%10d] Withdrawer Balance: %d\n", timeDiff(), account.Balance())
	defer waitgroup.Done()
}

func main() {
	waitgroup.Add(1)
	go withdrawer()
	for i := 1; i <= 6; i++ {
		waitgroup.Add(1)
		go depositor(i)
	}
	waitgroup.Wait()
	fmt.Printf("[%10d] Final Balance: %d\n", timeDiff(), account.Balance())
}

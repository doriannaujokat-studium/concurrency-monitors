package bankacc

import (
	. "sync"
)

type BankAccount interface {
	Deposit(a uint) uint
	Withdraw(a uint) uint
	Balance() uint
}

func New() BankAccount { return new_() }

type bankaccount struct {
	balance uint "balance"
	Mutex
	*Cond
}

func new_() BankAccount {
	x := new(bankaccount)
	x.Cond = NewCond(&x.Mutex)
	x.balance = 0
	return x
}

func (x *bankaccount) Deposit(a uint) uint {
	x.Mutex.Lock()
	defer x.Mutex.Unlock()
	x.balance += a
	x.Cond.Signal()
	return a
}

func (x *bankaccount) Withdraw(a uint) uint {
	x.Mutex.Lock()
	defer x.Mutex.Unlock()
	for x.balance < a {
		x.Cond.Wait()
	}
	x.balance -= a
	x.Cond.Signal()
	return a
}

func (x *bankaccount) Balance() uint {
	return x.balance
}

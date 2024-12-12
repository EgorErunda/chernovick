package main

import (
	"errors"
)

// Account представляет банковский счёт
type Account struct {
	owner   string
	balance float64
}

// NewAccount создаёт новый счёт с указанным владельцем и начальным балансом
func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

// GetBalance возвращает текущий баланс счёта
func (a *Account) GetBalance() float64 {
	return a.balance
}

// SetBalance устанавливает новый баланс счёта
func (a *Account) SetBalance(newBalance float64) error {
	if newBalance < 0 {
		return errors.New("баланс не может быть отрицательным")
	}
	a.balance = newBalance
	return nil
}

// GetOwner возвращает владельца счёта
func (a *Account) GetOwner() string {
	return a.owner
}

// Deposit добавляет деньги на счёт
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма для внесения должна быть положительной")
	}
	a.balance += amount

	return nil
}

// Withdraw снимает деньги со счёта
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма для снятия должна быть положительной")
	}
	if a.balance < amount {
		return errors.New("недостаточно средств на счёте")
	}
	a.balance -= amount

	return nil
}

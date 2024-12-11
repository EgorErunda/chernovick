package main

import "errors"

type Account struct {
	balance float64
	owner   string
}

//возвращает текущий баланс счёта.
func (a Account) GetBalance() float64 {
	return a.balance
}

//устанавливает новый баланс счёта.
func (a Account) SetBalance(newbalance float64) error {
	if newbalance < 0 {
		return errors.New("баланс не может быть отрицательным")
	}
	a.balance = newbalance
	return nil
}

//добавляет указанную сумму на счёт.
func (a Account) Deposit(summ float64) error {
	if summ < 0 {
		return errors.New("error")
	}
	a.balance += summ
	return nil
}

//снимает указанную сумму со счёта.
func (a Account) Withdraw(spisanie float64) error {
	if spisanie > a.balance {
		return errors.New("списание не может быть больше баланса")
	}
	a.balance -= spisanie
	return nil
}

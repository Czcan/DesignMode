package main

type state interface {
	requestItem() error
	addItem(int) error
	insertMoney(money int) error
	dispenseItem() error
}

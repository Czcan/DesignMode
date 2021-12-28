package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine //每个具体状态都有上下文的引用 ， 目的是 拿取 上下文的值 ， 以便根据 上下文中的值 来改变状态
}

func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 { // 根据 上下文中的值 来改变状态
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

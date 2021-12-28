package main

import "fmt"

//自动售卖机
type vendingMachine struct {
	hasItem       state //有商品
	itemRequested state //商品已请求
	hasMoney      state //收到money
	noItem        state //没有商品

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount int, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}

	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}

	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}

	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	v.itemRequested = itemRequestedState
	return v
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) incrementItemCount(count int) { // 抽取的 各种具体状态 中 可能的相同行为， 避免出现重复代码
	fmt.Printf("Adding %v item", count)
	v.itemCount = v.itemCount + count
}

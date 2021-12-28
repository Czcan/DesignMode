package main

import (
	"fmt"
	"log"
)

//状态模式和策略模式的不同在于： 策略模式下，客户端需要知晓所有策略，策略间的不同 。 各种策略相互独立--》 需要选择合适的策略
// 而状态模式： 各种状态 互相知晓， 在具体状态中 实现状态切换， 客户端可以不清楚 状态
func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem() //state: hasitem -> itemRequested
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10) //state: itemRequested -> hasMoney
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem() // itemcount: 1 -> 0   ----> state: hasMoney -> noItem
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2) //  itemcount: 0 -> 2   ----> state: noItem -> hasItem
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem() //state: hasitem -> itemRequested
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10) //state: itemRequested -> hasMoney
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem() // itemcount: 2 -> 1   ----> state: hasMoney -> hasItem
	if err != nil {
		log.Fatalf(err.Error())
	}
}

package main

import (
	"Register/src/ioc"
	_ "Register/src/order"
)

func main() {
	OrderController := ioc.Create("OrderController")
	OrderController.Handle()
}

package main

import (
	"Frame/src/ioc"
	_ "Frame/src/route"
	_ "Frame/src/log"
	_ "Frame/src/pay"
)

func main() {
	post := ioc.Make("Route:Post")
	post.Handle()
	get := ioc.Make("Route:Get")
	get.Handle()

	error := ioc.Make("LogErr")
	error.Handle()
	info := ioc.Make("LogInfo")
	info.Handle()

	ioc.Make("Wechat").Handle()
	ioc.Make("Alipay").Handle()
}
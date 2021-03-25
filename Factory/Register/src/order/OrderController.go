package order

import (
	"Register/src/ioc"
	"fmt"
)

type Undifined struct {}

func init()  {
	ioc.Register("OrderController", func () ioc.Controller {
		return new(Undifined)
	})
}

func (u *Undifined) Handle() {
	fmt.Println("Undifined >  Handle")
}

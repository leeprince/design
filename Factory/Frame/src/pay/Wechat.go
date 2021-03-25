package pay

import "fmt"

type Wechat struct {}

func (Wechat) Handle()  {
	fmt.Println("pay > Wechat >  Handle")
}

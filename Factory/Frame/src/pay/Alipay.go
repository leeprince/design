package pay

import "fmt"

type Alipay struct {}

func (Alipay) Handle()  {
	fmt.Println("pay > Alipay >  Handle")
}

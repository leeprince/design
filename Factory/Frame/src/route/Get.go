package route

import (
	"fmt"
)

type Get struct {}

func (Get) Handle()  {
	fmt.Println("Get >  Handle")
}


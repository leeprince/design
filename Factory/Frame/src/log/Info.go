package log

import (
	"fmt"

)

type Info struct {}

func (Info) Handle()  {
	fmt.Println("Info >  Handle")
}


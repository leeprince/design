package log

import (
	"fmt"

)

type Error struct {}

func (Error) Handle()  {
	fmt.Println("Error >  Handle")
}


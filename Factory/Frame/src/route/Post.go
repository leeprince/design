package route

import (
	"fmt"
)

type Post struct {}

func (Post) Handle()  {
	fmt.Println("Post >  Handle")
}


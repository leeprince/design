/*
单例模式
 */
package main

import (
	"Singleton/src/person"
	"fmt"
)

func main() {
	ps := person.NewPerson()
	ps.SetName("leeprince")
	fmt.Println("ps.GetName():", ps.GetName())
	ps.SetAge(18)
	fmt.Println("ps.GetAge():", ps.GetAge())
}





package storage

import "fmt"

type Mysql struct {
	ml map[string]interface{}
}

func (m *Mysql) Write(k string, v interface{}) bool {
	m.ml = map[string]interface{}{
		k: v,
	}
	fmt.Printf("k: %v = %v \n", k, v)
	return true
}
func (m *Mysql) Read(k string) interface{} {
	return m.ml[k]
}

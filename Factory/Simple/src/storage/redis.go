package storage

import "fmt"

type Redis struct {
	rs map[string]interface{}
}

func (r *Redis) Write(k string, v interface{}) bool  {
	r.rs = map[string]interface{}{
		k: v,
	}
	fmt.Printf("k: %v = %v \n", k, v)
	return true
}
func (r *Redis) Read(k string) interface{}  {
	return r.rs[k]
}
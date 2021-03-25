package storage

import "fmt"

type File struct {
	fe map[string]interface{}
}

func (f *File) Write(k string, v interface{}) bool {
	f.fe = map[string]interface{}{
		k: v,
	}
	fmt.Printf("k: %v = %v \n", k, v)
	return true
}
func (f *File) Read(k string) interface{} {
	return f.fe[k]
}

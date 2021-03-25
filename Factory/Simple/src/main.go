/*
简单工厂模式
 */
package main

import (
	"Simple/src/storage"
	"fmt"
)

func main() {
	file := storage.NewStorage("File")
	file.Write("name", "leeprince-file")
	fmt.Println("File:read>name:", file.Read("name"))

	mysql := storage.NewStorage("Mysql")
	mysql.Write("name", "leeprince-mysql")
	fmt.Println("Mysql:read>name:", mysql.Read("name"))

	redis := storage.NewStorage("Redis")
	redis.Write("name", "leeprince-redis")
	fmt.Println("Redis:read>name:", redis.Read("name"))
}

package storage

type Storage interface {
	Write(k string, v interface{}) bool
	Read(k string) interface{}
}

const (
	MYSQL = "Mysql"
	FILE = "File"
	REDIS = "Redis"
)

func NewStorage(s string) Storage  {
	switch s {
	case MYSQL:
		return &Mysql{}
	case FILE:
		return &File{}
	case REDIS:
		return &Redis{}
	default:
		panic("实现 Storage 接口的方法不存在")
	}
}

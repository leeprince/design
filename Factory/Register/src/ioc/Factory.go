package ioc

// 定义Controller接口
type Controller interface {
	Handle()
}

// 保存注册的"名称"及对应实现定义Controller接口的方法
var Ioc = make(map[string]func() Controller)

/*
注册
	指定一个键名， 并绑定实现Controller接口类的方法
 */
func Register(name string, callFun func() Controller)  {
	Ioc[name] = callFun
}

/*
创建
	生成已注册的方法
 */
func Create(name string) Controller {
	if f, ok := Ioc[name]; ok {
		return f()
	}
	panic("解析的方法不存在")
}

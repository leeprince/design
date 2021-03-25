package ioc

// 定义 Object 接口
type Object interface {
	Handle()
}

// 保存注册的"具体实现的名称"及对应实现定义 Object 接口的方法
type ioc struct {
	bindings map[string]func() Object
}

var base = &ioc{
	bindings: make(map[string]func() Object),
}

/*
注册
	指定一个键名， 并绑定实现 Object 接口类的方法
 */
func Bind(concrete string, callFun func() Object)  {
	base.bindings[concrete] = callFun
}

/*
创建
	生成已注册的方法
 */
func Make(concrete string) Object {
	if f, ok := base.bindings[concrete]; ok {
		return f()
	}
	panic("解析的方法不存在")
}


package caculate

/*
对于go语言，没有真正的继承。靠结构体的嵌套模拟继承功能的。
结构体的嵌套：
      聚合关系：一个类作为另一个类的属性。has - a
         type A struct{}
         type B struct{
            a A
         }
      继承关系：一个子类继承一个父类。is - a
         type C struct{}
         type D struct{
            C
         }
实例：
	type A struct{ //父类
		name string
		age int
	}
	type B struct{ //子类
		a A // 模拟聚合关系
	}
	type C struct{ //子类
		A // 模拟继承.//在结构体中属于匿名结构体的字段称为提升字段
	}
	b := B{}
	b.a.name

	c := C{}
	b.name
 */
// 结构体的嵌套：继承CaculateData结构体

type AddResult struct {
	*CaculateData  // 不能定义变量名
}
// 实现CaculateFactory工厂接口
type AddResultFatory struct {}

func (c AddResult) Result() float64 {
	return c.L + c.R
}
func (AddResultFatory) Create() Caculate {
	return &AddResult{
		&CaculateData{},
	}
}
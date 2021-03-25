package caculate

// 结构体继承：继承CaculateData结构体
type MulResult struct {
	*CaculateData  // 不能定义变量名
}

// 实现CaculateFactory工厂接口
type MulResultFatory struct {}

func (c MulResult) Result() float64 {
	return c.L * c.R
}
func (MulResultFatory) Create() Caculate {
	return &MulResult{
		&CaculateData{},
	}
}
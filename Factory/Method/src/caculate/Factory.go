package caculate

// 实际接口
type Caculate interface {
	SetLeft(float64)
	SetRight(float64)
	Result() float64
}
// 工厂接口
type CaculateFactory interface {
	Create() Caculate
}

// 其他结构体中会继承该结构体，并真正实现Caculate接口
type CaculateData struct {
	L float64
	R float64
}
func (d *CaculateData) SetLeft(l float64) {
	d.L = l
}
func (d *CaculateData) SetRight(r float64) {
	d.R = r
}




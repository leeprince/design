package order

import "fmt"

// 实现订单的接口
type FileOrderDao struct {}
func (*FileOrderDao) SaveOrderMain() {
	fmt.Println("FileOrderDao.SaveOrderMain")
}
// 实现订单详情的接口
type FileOrderDetailDao struct {}
func (*FileOrderDetailDao) SaveOrderDetail() {
	fmt.Println("FileOrderDetailDao.SaveOrderDetail")
}

// 实现订单的抽象工厂模式接口
type FileDaoFatory struct {}
func (*FileDaoFatory) CreateOrderDao() OrderDao {
	return &FileOrderDao{}
}
func (*FileDaoFatory) CreateOrderDetailDao() OrderDetailDao {
	return &FileOrderDetailDao{}
}
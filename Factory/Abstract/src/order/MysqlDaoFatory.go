package order

import "fmt"

// 实现订单的接口
type MysqlOrderDao struct {}
func (*MysqlOrderDao) SaveOrderMain() {
	fmt.Println("MysqlOrderDao.SaveOrderMain")
}
// 实现订单详情的接口
type MysqlOrderDetailDao struct {}
func (*MysqlOrderDetailDao) SaveOrderDetail() {
	fmt.Println("MysqlOrderDetailDao.SaveOrderDetail")
}

// 实现订单的抽象工厂模式接口
type MysqlDaoFatory struct {}
func (*MysqlDaoFatory) CreateOrderDao() OrderDao {
	return &MysqlOrderDao{}
}
func (*MysqlDaoFatory) CreateOrderDetailDao() OrderDetailDao {
	return &MysqlOrderDetailDao{}
}
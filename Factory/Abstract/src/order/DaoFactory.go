package order

// 订单的接口
type OrderDao interface {
	SaveOrderMain()
}
// 订单详情的接口
type OrderDetailDao interface {
	SaveOrderDetail()
}

// 订单的抽象工厂模式接口
type DaoFactory interface {
	CreateOrderDao() OrderDao
	CreateOrderDetailDao() OrderDetailDao
}

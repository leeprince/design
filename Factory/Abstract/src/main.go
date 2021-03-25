package main

import "Abstract/src/order"

func main() {
	file := order.FileDaoFatory{}
	file.CreateOrderDao().SaveOrderMain()
	file.CreateOrderDetailDao().SaveOrderDetail()

	mysql := order.MysqlDaoFatory{}
	mysql.CreateOrderDao().SaveOrderMain()
	mysql.CreateOrderDetailDao().SaveOrderDetail()


	/* 完整定义 */
	/*var factory order.DaoFactory
	// factory = &order.FileDaoFatory{}
	factory = &order.MysqlDaoFatory{}*/
	/* 简洁定义 */
	factory := &order.FileDaoFatory{}
	getOrderAndOrderDetail(factory)

}

/*
整合 OrderDao、OrderDetailDao
 */
func getOrderAndOrderDetail(factory order.DaoFactory) {
	factory.CreateOrderDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

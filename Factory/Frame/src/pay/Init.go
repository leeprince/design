package pay

import "Frame/src/ioc"

func init() {
	ioc.Bind("Wechat", func () ioc.Object {
		return new(Wechat)
	})
	ioc.Bind("Alipay", func () ioc.Object {
		return new(Alipay)
	})
}

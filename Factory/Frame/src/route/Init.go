package route

import "Frame/src/ioc"

/*
一个包中
 */
func init()  {
	ioc.Bind("Route:Post", func () ioc.Object {
		return new(Post)
	})

	ioc.Bind("Route:Get", func () ioc.Object {
		return new(Get)
	})
}

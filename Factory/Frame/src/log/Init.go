package log

import "Frame/src/ioc"

func init()  {
	ioc.Bind("LogInfo", func () ioc.Object {
		return new(Info)
	})

	ioc.Bind("LogErr", func () ioc.Object {
		return new(Error)
	})
}

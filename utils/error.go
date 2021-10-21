package utils

import (
	"github.com/Qianjiachen55/Daily-recommed/global"
	"github.com/spf13/cast"
)

// 可以在 defer处使用此方法
func PanicFun()  {
	if err := recover();err!= nil{
		global.DrLogger.Error(cast.ToString(err))
	}
}

// 普通error 调用此方法
func ErrFun(err interface{})  {
	global.DrLogger.Error(cast.ToString(err))
}

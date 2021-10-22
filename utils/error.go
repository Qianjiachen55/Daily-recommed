package utils

import (
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func ErrorFun(logger *zap.Logger,err error)  {
	logger.Error(cast.ToString(err))
}

func PaincFun(logger *zap.Logger)  {
	if err:=recover() ;err!= nil{
		logger.Error(cast.ToString(err))
		panic(err)
	}

}

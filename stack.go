package goutils

import (
	"fmt"
	"github.com/0studio/logger"
	log "github.com/cihub/seelog"
	"runtime"
)

func ProtectFuncWithLogger(fun func(), _log logger.Logger) {
	defer func() {
		if x := recover(); x != nil {
			if _log != nil {
				_log.Errorf("%v", x)
			} else {

				fmt.Println("%v", x)
			}

			for i := 0; i < 10; i++ {
				funcName, file, line, ok := runtime.Caller(i)
				if ok {
					if _log != nil {
						_log.Errorf("frame %v:[%v,file:%v,line:%v]", i, runtime.FuncForPC(funcName).Name(), file, line)
					} else {
						fmt.Printf("frame %v:[%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)

					}

				}
			}
		}
	}()
	fun()
}
func ProtectFunc(fun func()) {
	defer func() {
		if x := recover(); x != nil {
			log.Errorf("%v", x)
			for i := 0; i < 10; i++ {
				funcName, file, line, ok := runtime.Caller(i)
				if ok {
					log.Errorf("frame %v:[%v,file:%v,line:%v]", i, runtime.FuncForPC(funcName).Name(), file, line)
				}
			}
		}
	}()
	fun()
}
func PrintStack() {
	for i := 0; i < 10; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			// fmt.Printf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			log.Errorf("frame %v:[func:%v,file:%v,line:%v]", i, runtime.FuncForPC(funcName).Name(), file, line)
		}
	}

}
func PrintStackWithLogger(_log logger.Logger) {
	for i := 0; i < 10; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			_log.Errorf("frame %v:[func:%v,file:%v,line:%v]", i, runtime.FuncForPC(funcName).Name(), file, line)
		}
	}

}

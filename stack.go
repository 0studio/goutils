package goutils

import (
	// "fmt"
	log "github.com/cihub/seelog"
	"runtime"
)

func ProtectFunc(fun func()) {
	defer func() {
		if x := recover(); x != nil {
			log.Errorf("%v", x)
			for i := 0; i < 10; i++ {
				funcName, file, line, ok := runtime.Caller(i)
				if ok {
					log.Errorf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
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
			log.Errorf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		}
	}

}

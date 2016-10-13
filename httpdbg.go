package httpdbg

import (
	"fmt"
	"runtime"
	"net/http"
	"net/http/httputil"
)


func Reqout(req *http.Request) {
	// programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	_ = sourceFileName
	_ = sourceFileLineNum
	_ = ok
	fn := runtime.FuncForPC(programCounter)
	reqdata, err := httputil.DumpRequest(req, true)
	_ = err
	fmt.Printf("[TOSEKI] call from func %s \n Reqdata= %s\n",fn.Name(),reqdata)
	// fmt.Printf("[TOSEKI] call from %s \n Reqdata= \n %s\n",fn.Name(),req)
}

func Resout(res *http.Response) {
	programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	_ = sourceFileName
	_ = sourceFileLineNum
	_ = ok
	fn := runtime.FuncForPC(programCounter)
	resdata, err := httputil.DumpResponse(res, true)
	_ = err
	fmt.Printf("[TOSEKI] call from func %s \n Response= %s\n",fn.Name(),resdata)
}

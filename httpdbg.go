package httpdbg

import (
	"fmt"
	"runtime"
	"net/http"
	"net/http/httputil"
)


func reqout(req *http.Request) {
	// programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	_ = sourceFileName
	_ = sourceFileLineNum
	_ = ok
	fn := runtime.FuncForPC(programCounter)
	reqdata, err := httputil.DumpRequestOut(req, true)
	_ = err
	fmt.Print("[TOSEKI] %s Req \n %s\n",fn.Name(),reqdata)
}

func resout(res *http.Response) {
	programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	_ = sourceFileName
	_ = sourceFileLineNum
	_ = ok
	fn := runtime.FuncForPC(programCounter)
	resdata, err := httputil.DumpResponse(res, true)
	_ = err
	fmt.Print("[TOSEKI] %s Res \n %s\n",fn.Name(),resdata)
}

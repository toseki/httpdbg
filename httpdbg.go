package httpdbg

import (
	"fmt"
	"runtime"
	"net/http"
	"net/http/httputil"
)


func Reqout(req *http.Request) {
	// programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	fmt.Print("[TOSEKI]httpdbg.Reqout\n")
	for depth :=0; ; depth++{
		programCounter, srcFileName, srcFileLineNum, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		fmt.Printf(" -> %d: %s: %s(%d)\n", depth, runtime.FuncForPC(programCounter).Name(), srcFileName, srcFileLineNum)
	}
	reqdata, err := httputil.DumpRequest(req, true)
	_ = err
	fmt.Printf("[TOSEKI]httpdbg - Reqdata= %s\n",reqdata)
}

func Resout(res *http.Response) {
	fmt.Print("[TOSEKI]httpdbg.Resout\n")
	for depth :=0; ; depth++{
		programCounter, srcFileName, srcFileLineNum, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		fmt.Printf(" -> %d: %s: %s(%d)\n", depth, runtime.FuncForPC(programCounter).Name(), srcFileName, srcFileLineNum)
	}
	resdata, err := httputil.DumpResponse(res, true)
	_ = err
	fmt.Printf("[TOSEKI]httpdbg - Response= %s\n",resdata)
}

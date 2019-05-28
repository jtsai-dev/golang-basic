package main

import (
	"error/filehandling"
	"fmt"
	"log"
	"net/http"
	"os"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic(r)
		}
	}()
}

type appHanlder func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHanlder) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			// 对 error 的分类处理
			log.Printf("Error handing request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

/* panic:
 *  停止当前函数的执行
 *  向上返回，执行每一层的 defer
 *  若无遇见 recover，则程序退出
 *  对比 error:
 *   意料之中的情景下尽量使用 error，如文件无法打开等
 *   医疗之外的情景下使用 panic，如数组越界等
 * recover：
 *  仅在 defer 中调用
 *  获取 panic 的值
 *  若无法处理，可重新 panic
 */
func main() {
	// 定义 http 服务的 handle，对句柄做 errorWrapper 的包装
	http.HandleFunc("/", errorWrapper(filehandling.HandleFile))

	// 启动服务，监听指定的端口
	err := http.ListenAndServe(":900", nil)
	if err != nil {
		panic(err)
	}
}

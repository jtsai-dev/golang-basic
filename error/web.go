package main

import (
	"golang-basic/error/filehandling"
	"log"
	"net/http"
	"os"
)

type userError interface {
	error
	Message() string
}

type appHanlder func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHanlder) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// 使用 recover 捕获异常，避免 panic 导致的程序退出
		defer func() {
			r := recover()
			if err, ok := r.(error); ok {
				log.Printf("Panic: %v", r)
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			} else {
				log.Println("Unknowed error:", r, err)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error handing request: %s", err.Error())

			// 自定义错误类型的处理
			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}

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
